package v0

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/nyakaspeter/raven-torrent/internal/torrentclient"
)

// @Router /deleteall [get]
// @Summary Delete all torrents from torrent client
// @Description
// @Tags Torrent client
// @Success 200 {object} MessageResponse
// @Failure 404 {object} MessageResponse
func DeleteAllTorrents() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("Deleting all torrents.")

		if len(torrentclient.ActiveTorrents) == 0 {
			http.Error(w, noActiveTorrentsFound(), http.StatusNotFound)
			return
		}

		for _, torrent := range torrentclient.ActiveTorrents {
			torrentclient.RemoveTorrent(torrent.Torrent.InfoHash().String())
		}

		io.WriteString(w, deletedAllTorrents())
	}
}

func deletedAllTorrents() string {
	message := MessageResponse{
		Success: true,
		Message: "All torrents have been deleted.",
	}

	messageString, _ := json.Marshal(message)

	log.Println("Deleted all torrents.")

	return string(messageString)
}

func noActiveTorrentsFound() string {
	message := MessageResponse{
		Success: false,
		Message: "No active torrents found.",
	}

	messageString, _ := json.Marshal(message)

	log.Println("No active torrents found.")

	return string(messageString)
}
