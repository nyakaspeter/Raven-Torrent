package torrentclient

import (
	"bytes"
	"crypto/tls"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"net/url"
	"runtime"
	"strings"
	"time"

	alog "github.com/anacrolix/log"
	"github.com/anacrolix/torrent"
	"github.com/anacrolix/torrent/metainfo"
	"github.com/anacrolix/torrent/storage"
	"github.com/nyakaspeter/raven-torrent/internal/torrentclient/memorystorage"
	"github.com/oz/osdb"
	"golang.org/x/time/rate"
)

// Torrent lock structure
type TorrentLeaf struct {
	Torrent     *torrent.Torrent
	Progress    int64          // Downoad stats measurement
	Prevtime    time.Time      // Previous time for progress calculation
	FileClients map[string]int // Count active connections
}

const resolveTimeout = time.Second * 35
const megaByte = 1024 * 1024

var torrentClient *torrent.Client
var receivedTorrent *metainfo.MetaInfo = nil
var maxPieceLength int64 = 16

var ActiveTorrents map[string]*TorrentLeaf

func StartTorrentClient(
	storageType string,
	memorySize int64,
	downloadDir string,
	downloadRate int,
	uploadRate int,
	maxConnections int,
	noDht bool,
	enableLog bool,
) (*torrent.Client, error) {
	ActiveTorrents = make(map[string]*TorrentLeaf)

	cfg := torrent.NewDefaultClientConfig()

	if storageType == "memory" {
		maxPieceLength = int64(math.Floor(float64(memorySize) * 100 / 75 / 8))
		memorystorage.SetMemorySize(memorySize, maxPieceLength)
		cfg.DefaultStorage = memorystorage.NewMemoryStorage()
	} else if storageType == "file" {
		cfg.DefaultStorage = storage.NewFileByInfoHash(downloadDir)
		cfg.DataDir = downloadDir
	}

	cfg.EstablishedConnsPerTorrent = maxConnections
	cfg.NoDHT = noDht
	cfg.DisableIPv6 = true
	cfg.DisableUTP = true

	// Discard or show the logs
	if !enableLog {
		cfg.Logger = alog.Discard
	}
	//cfg.Debug = true

	// up/download speed rate in bytes per second from megabits per second
	downrate := int((downloadRate * 1024) / 8)
	uprate := int((uploadRate * 1024) / 8)

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

func AddTorrent(uri string) *torrent.Torrent {
	var spec *torrent.TorrentSpec
	var t *torrent.Torrent
	var err error = nil

	if strings.HasPrefix(uri, "magnet:") {
		// Add magnet link
		spec, err = torrent.TorrentSpecFromMagnetUri(uri)
		receivedTorrent = nil
	} else if receivedTorrent != nil {
		// Add previously received torrent file
		spec = torrent.TorrentSpecFromMetaInfo(receivedTorrent)
	} else {
		// Download torrent file from
		r, e := fetchTorrent(uri)
		if e == nil {
			receivedTorrent, err = metainfo.Load(r)
			spec = torrent.TorrentSpecFromMetaInfo(receivedTorrent)
		} else {
			urlError, isUrlError := e.(*url.Error)
			if isUrlError && strings.HasPrefix(urlError.URL, "magnet:") {
				// If Jackett redirected to a magnet link
				uri = urlError.URL
				spec, err = torrent.TorrentSpecFromMagnetUri(uri)
				receivedTorrent = nil
			} else {
				err = e
			}
		}
	}

	if err != nil {
		log.Println(err)
		return nil
	}

	infoHash := spec.InfoHash.String()
	if t, ok := ActiveTorrents[infoHash]; ok {
		receivedTorrent = nil
		return t.Torrent
	}

	// // Intended for streaming so only one torrent stream allowed at a time
	// if len(torrents) > 0 || gettingTorrent == true {
	// 	log.Println("Only one torrent stream allowed at a time.")
	// 	return nil
	// }

	if receivedTorrent != nil {
		t, err = torrentClient.AddTorrent(receivedTorrent)
	} else {
		t, err = torrentClient.AddMagnet(uri)
	}

	if err != nil {
		log.Panicln(err)
		return nil
	}

	select {
	case <-t.GotInfo():
		if t.Info().PieceLength <= (maxPieceLength * megaByte) {
			ActiveTorrents[t.InfoHash().String()] = &TorrentLeaf{
				Torrent:     t,
				Progress:    0,
				Prevtime:    time.Now(),
				FileClients: make(map[string]int),
			}
			receivedTorrent = nil
			return t
		} else {
			t.Drop()
			receivedTorrent = nil
			return nil
		}
	case <-time.After(resolveTimeout):
		t.Drop()
		receivedTorrent = nil
		return nil
	}
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

func StopDownloadFile(file *torrent.File) {
	if file != nil {
		file.SetPriority(torrent.PiecePriorityNone)
	}
}

func StopAllFileDownload(files []*torrent.File) {
	for _, f := range files {
		f.SetPriority(torrent.PiecePriorityNone)
	}
}

func IncreaseFileClients(path string, t *TorrentLeaf) int {
	if v, ok := t.FileClients[path]; ok {
		v++
		t.FileClients[path] = v
		return v
	} else {
		t.FileClients[path] = 1
		return 1
	}
}

func DecreaseFileClients(path string, t *TorrentLeaf) int {
	if v, ok := t.FileClients[path]; ok {
		v--
		t.FileClients[path] = v
		return v
	} else {
		t.FileClients[path] = 0
		return 0
	}
}

func GetFileByPath(search string, files []*torrent.File) int {

	for i, f := range files {
		if search == f.DisplayPath() {
			return i
		}
	}

	return -1
}

func fetchTorrent(url string) (*bytes.Reader, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{Transport: tr}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		b, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, errors.New(resp.Status)
		}
		return nil, errors.New(string(b))
	}

	buf := &bytes.Buffer{}

	_, err = io.Copy(buf, resp.Body)
	if err != nil {
		return nil, err
	}

	return bytes.NewReader(buf.Bytes()), nil
}
