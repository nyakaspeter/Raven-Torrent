package handlers

import (
	"encoding/base64"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"sort"
	"strconv"
	"time"

	"github.com/anacrolix/torrent"
	"github.com/gorilla/mux"
	"github.com/nyakaspeter/raven-torrent/internal/api/v0/responses"
	"github.com/nyakaspeter/raven-torrent/internal/torrentclient"
)

type TorrentFilesResponse struct {
	Name   string `json:"name"`
	Url    string `json:"url"`
	Length string `json:"length"`
}

type TorrentFilesResultsResponse struct {
	Success bool                   `json:"success"`
	Hash    string                 `json:"hash"`
	Results []TorrentFilesResponse `json:"results"`
}

func AddTorrent() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		base64uri := vars["base64uri"]

		uri, err := base64.StdEncoding.DecodeString(base64uri)

		if err != nil {
			http.Error(w, failedToAddTorrent(), http.StatusNotFound)
			return
		}

		for tryCount := 0; tryCount < 4; tryCount++ {
			if tryCount > 0 {
				time.Sleep(10 * time.Second)
			}

			t := torrentclient.AddTorrent(string(uri))

			if t != nil {
				log.Println("Added torrent:", t.InfoHash().String())
				io.WriteString(w, torrentFilesList(r.Host, t.InfoHash().String(), t.Files()))
				return
			} else if len(torrentclient.ActiveTorrents) == 0 {
				http.Error(w, failedToAddTorrent(), http.StatusNotFound)
				return
			}
		}

		if len(torrentclient.ActiveTorrents) > 0 {
			http.Error(w, onlyOneTorrentAllowed(), http.StatusNotFound)
		}
	}
}

func torrentFilesList(address string, infohash string, files []*torrent.File) string {
	sortFiles(files)

	var results []TorrentFilesResponse

	for _, f := range files {
		result := TorrentFilesResponse{
			Name:   f.DisplayPath(),
			Url:    "http://" + address + "/api/get/" + f.Torrent().InfoHash().String() + "/" + base64.StdEncoding.EncodeToString([]byte(f.DisplayPath())),
			Length: strconv.FormatInt(f.FileInfo().Length, 10),
		}

		results = append(results, result)
	}

	message := TorrentFilesResultsResponse{
		Success: true,
		Hash:    infohash,
		Results: results,
	}

	messageString, _ := json.Marshal(message)

	return string(messageString)
}

func onlyOneTorrentAllowed() string {
	message := responses.MessageResponse{
		Success: false,
		Message: "Only one torrent stream allowed at a time.",
	}

	messageString, _ := json.Marshal(message)

	return string(messageString)
}

func failedToAddTorrent() string {
	message := responses.MessageResponse{
		Success: false,
		Message: "Failed to add torrent.",
	}

	messageString, _ := json.Marshal(message)

	return string(messageString)
}

func sortFiles(files []*torrent.File) {
	sort.Slice(files, func(i, j int) bool {
		return files[i].DisplayPath() < files[j].DisplayPath()
	})
}
