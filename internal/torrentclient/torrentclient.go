package torrentclient

import (
	"bytes"
	"encoding/base64"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"math"
	"net/http"
	"runtime"
	"strconv"
	"time"

	alog "github.com/anacrolix/log"
	"github.com/anacrolix/torrent"
	"github.com/anacrolix/torrent/metainfo"
	"github.com/anacrolix/torrent/storage"
	"github.com/nyakaspeter/raven-torrent/internal/settings"
	"github.com/nyakaspeter/raven-torrent/internal/torrentclient/memorystorage"
	"github.com/nyakaspeter/raven-torrent/internal/torrentclient/types"
	"github.com/nyakaspeter/raven-torrent/pkg/utils"
	"github.com/oz/osdb"
	"golang.org/x/time/rate"
)

const resolveTimeout = time.Second * 35
const megaByte = 1024 * 1024

var torrentClient *torrent.Client
var receivedTorrent *metainfo.MetaInfo = nil
var maxPieceLength int64 = 16

var ActiveTorrents map[string]*types.TorrentLeaf

func StartTorrentClient() (*torrent.Client, error) {
	ActiveTorrents = make(map[string]*types.TorrentLeaf)

	cfg := torrent.NewDefaultClientConfig()

	if *settings.StorageType == "memory" {
		maxPieceLength = int64(math.Floor(float64(*settings.MemorySize) * 100 / 75 / 8))
		memorystorage.SetMemorySize(*settings.MemorySize, maxPieceLength)
		cfg.DefaultStorage = memorystorage.NewMemoryStorage()
	} else if *settings.StorageType == "file" {
		cfg.DefaultStorage = storage.NewFileByInfoHash(*settings.DownloadDir)
		cfg.DataDir = *settings.DownloadDir
	}

	cfg.EstablishedConnsPerTorrent = *settings.MaxConnections
	cfg.NoDHT = *settings.NoDHT
	cfg.DisableIPv6 = true
	cfg.DisableUTP = true

	// Discard or show the logs
	if !*settings.EnableLog {
		cfg.Logger = alog.Discard
	}
	//cfg.Debug = true

	// up/download speed rate in bytes per second from megabits per second
	downrate := int((*settings.DownloadRate * 1024) / 8)
	uprate := int((*settings.UploadRate * 1024) / 8)

	if downrate != 0 {
		cfg.DownloadRateLimiter = rate.NewLimiter(rate.Limit(downrate), downrate)
	}

	if uprate == 0 {
		cfg.NoUpload = true
	} else {
		cfg.UploadRateLimiter = rate.NewLimiter(rate.Limit(uprate), uprate)
	}

	var err error = nil
	torrentClient, err = torrent.NewClient(cfg)

	return torrentClient, err
}

func StopTorrentClient() {
	if torrentClient == nil {
		return
	}

	torrentClient.Close()

	torrentClient = nil
	ActiveTorrents = nil
	receivedTorrent = nil

	runtime.GC()
}

func AddTorrent(uri string) types.TorrentInfo {
	info := types.TorrentInfo{}
	torrent := addTorrentFromUri(uri)

	if torrent == nil {
		return info
	}

	info.Hash = torrent.InfoHash().String()
	info.Name = torrent.Name()
	info.Length = strconv.FormatInt(torrent.Length(), 10)
	sortFiles(torrent.Files())

	for _, f := range torrent.Files() {
		tf := types.TorrentFile{
			Name:   f.DisplayPath(),
			Url:    "http://" + utils.GetLocalIP() + ":" + strconv.Itoa(*settings.Port) + "/api/v0/get/" + f.Torrent().InfoHash().String() + "/" + base64.StdEncoding.EncodeToString([]byte(f.DisplayPath())),
			Length: strconv.FormatInt(f.FileInfo().Length, 10),
		}

		info.Files = append(info.Files, tf)
	}

	return info
}

func ServeTorrentFile(w http.ResponseWriter, r *http.Request, file *torrent.File) {
	w.Header().Set("TransferMode.DLNA.ORG", "Streaming")
	w.Header().Set("contentFeatures.dlna.org", "DLNA.ORG_OP=01;DLNA.ORG_CI=0;DLNA.ORG_FLAGS=01700000000000000000000000000000")

	reader := file.NewReader()
	// Never set a smaller buffer than the maximum torrent piece length!
	reader.SetReadahead(maxPieceLength * megaByte)
	reader.SetResponsive()

	path := file.FileInfo().Path
	fname := ""
	if len(path) == 0 {
		fname = file.DisplayPath()
	} else {
		fname = path[len(path)-1]
	}

	http.ServeContent(w, r, fname, time.Unix(0, 0), reader)
}

func GetActiveTorrents() []types.TorrentInfo {
	activeTorrents := []types.TorrentInfo{}

	for _, t := range ActiveTorrents {
		at := types.TorrentInfo{}
		at.Hash = t.Torrent.InfoHash().String()
		at.Name = t.Torrent.Name()
		at.Length = strconv.FormatInt(t.Torrent.Length(), 10)

		for _, f := range t.Torrent.Files() {
			tf := types.TorrentFile{
				Name:   f.DisplayPath(),
				Url:    "http://" + utils.GetLocalIP() + ":" + strconv.Itoa(*settings.Port) + "/api/v0/get/" + f.Torrent().InfoHash().String() + "/" + base64.StdEncoding.EncodeToString([]byte(f.DisplayPath())),
				Length: strconv.FormatInt(f.FileInfo().Length, 10),
			}

			at.Files = append(at.Files, tf)
		}

		activeTorrents = append(activeTorrents, at)
	}

	return activeTorrents
}

func RemoveTorrent(hash string) error {
	if t, ok := ActiveTorrents[hash]; ok {
		for _, f := range t.Torrent.Files() {
			StopFileDownload(f)
		}
		t.Torrent.Drop()
		delete(ActiveTorrents, hash)
		return nil
	}

	return errors.New("torrent not found")
}

func StopFileDownload(file *torrent.File) {
	if file != nil {
		file.SetPriority(torrent.PiecePriorityNone)
	}
}

func IncreaseConnections(path string, t *types.TorrentLeaf) int {
	if v, ok := t.FileClients[path]; ok {
		v++
		t.FileClients[path] = v
		return v
	} else {
		t.FileClients[path] = 1
		return 1
	}
}

func DecreaseConnections(path string, t *types.TorrentLeaf) int {
	if v, ok := t.FileClients[path]; ok {
		v--
		t.FileClients[path] = v
		return v
	} else {
		t.FileClients[path] = 0
		return 0
	}
}

func GetFileIndexByPath(search string, files []*torrent.File) int {

	for i, f := range files {
		if search == f.DisplayPath() {
			return i
		}
	}

	return -1
}

func CalculateOpensubtitlesHash(file *torrent.File) string {
	fileReader := file.NewReader()

	if file.Length() < osdb.ChunkSize {
		return "0"
	}

	// The First and Last 65536 bytes are used to calculate the hash
	buffer := make([]byte, osdb.ChunkSize*2)

	fileReader.Seek(0, 0)
	_, err := fileReader.Read(buffer[:osdb.ChunkSize])
	if err != nil {
		return "0"
	}

	fileReader.Seek(-(osdb.ChunkSize), 2)
	_, err = fileReader.Read(buffer[osdb.ChunkSize:])
	if err != nil && err != io.EOF {
		return "0"
	}

	// Convert to uint64, and sum.
	var hash uint64
	nums := make([]uint64, ((osdb.ChunkSize * 2) / 8))
	bufferReader := bytes.NewReader(buffer)
	err = binary.Read(bufferReader, binary.LittleEndian, &nums)
	if err != nil {
		return "0"
	}
	for _, num := range nums {
		hash += num
	}

	return fmt.Sprintf("%016x", hash+uint64(file.Length()))
}
