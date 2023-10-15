package v0

import (
	"encoding/base64"
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nyakaspeter/raven-torrent/internal/torrentclient"
	torrentclienttypes "github.com/nyakaspeter/raven-torrent/internal/torrentclient/types"
)

type TorrentFilesResultsResponse struct {
	Success bool                             `json:"success"`
	Hash    string                           `json:"hash"`
	Results []torrentclienttypes.TorrentFile `json:"results"`
}

// @Router /add/{base64uri} [get]
// @Summary Get torrent info and streaming URLs
// @Description
// @Tags Torrent client
// @Param base64uri path string true "Link to torrent file / magnet link (base64 encoded)"
// @Success 200 {object} TorrentFilesResultsResponse
// @Failure 404 {object} MessageResponse
func AddTorrent() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		log.Println("Adding torrent:", vars)

		base64uri := vars["base64uri"]

		uri, err := base64.StdEncoding.DecodeString(base64uri)

		if err != nil {
			http.Error(w, failedToAddTorrent(), http.StatusNotFound)
			return
		}

		t := torrentclient.AddTorrent(string(uri))

		if t.Hash == "" || len(t.Files) == 0 {
			http.Error(w, failedToAddTorrent(), http.StatusNotFound)
			return
		}

		io.WriteString(w, torrentFilesList(t.Hash, t.Files))
	}
}

func torrentFilesList(infohash string, files []torrentclienttypes.TorrentFile) string {
	message := TorrentFilesResultsResponse{
		Success: true,
		Hash:    infohash,
		Results: files,
	}

	messageString, _ := json.Marshal(message)

	log.Println("Added torrent with", len(files), "files.")

	return string(messageString)
}

func failedToAddTorrent() string {
	message := MessageResponse{
		Success: false,
		Message: "Failed to add torrent.",
	}

	messageString, _ := json.Marshal(message)

	log.Println("Failed to add torrent.")

	return string(messageString)
}
