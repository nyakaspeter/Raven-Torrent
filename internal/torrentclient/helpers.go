package torrentclient

import (
	"bytes"
	"crypto/tls"
	"errors"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"sort"
	"strings"
	"time"

	"github.com/anacrolix/torrent"
	"github.com/anacrolix/torrent/metainfo"
	"github.com/nyakaspeter/raven-torrent/internal/torrentclient/types"
)

func addTorrentFromUri(uri string) *torrent.Torrent {
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
			ActiveTorrents[t.InfoHash().String()] = &types.TorrentLeaf{
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

func sortFiles(files []*torrent.File) {
	sort.Slice(files, func(i, j int) bool {
		return files[i].DisplayPath() < files[j].DisplayPath()
	})
}
