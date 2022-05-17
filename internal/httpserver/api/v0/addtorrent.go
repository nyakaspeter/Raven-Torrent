package v0

import (
	"encoding/base64"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/nyakaspeter/raven-torrent/internal/torrentclient"
	torrentclienttypes "github.com/nyakaspeter/raven-torrent/internal/torrentclient/types"
)

type TorrentFilesResultsResponse struct {
	Success bool                             `json:"success"`
	Hash    string                           `json:"hash"`
	Results []torrentclienttypes.TorrentFile `json:"results"`
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

			if t.Hash == "" || len(t.Files) == 0 {
				http.Error(w, failedToAddTorrent(), http.StatusNotFound)
				return
			}

			log.Println("Added torrent:", t.Hash)
			io.WriteString(w, torrentFilesList(t.Hash, t.Files))
		}
	}
}

func torrentFilesList(infohash string, files []torrentclienttypes.TorrentFile) string {
	message := TorrentFilesResultsResponse{
		Success: true,
		Hash:    infohash,
		Results: files,
	}

	messageString, _ := json.Marshal(message)

	return string(messageString)
}

func failedToAddTorrent() string {
	message := MessageResponse{
		Success: false,
		Message: "Failed to add torrent.",
	}

	messageString, _ := json.Marshal(message)

	return string(messageString)
}
