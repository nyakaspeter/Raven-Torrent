package v0

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nyakaspeter/raven-torrent/internal/torrentclient"
)

func DeleteTorrent() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		err := torrentclient.RemoveTorrent(vars["hash"])
		if err != nil {
			log.Println("Torrent not found:", vars["hash"])
			http.Error(w, torrentNotFound(), http.StatusNotFound)
			return
		}

		log.Println("Deleted torrent:", vars["hash"])
		io.WriteString(w, deletedTorrent())
	}
}

func DeleteAllTorrents() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if len(torrentclient.ActiveTorrents) == 0 {
			log.Println("No active torrents found.")
			http.Error(w, noActiveTorrentsFound(), http.StatusNotFound)
			return
		}

		for _, torrent := range torrentclient.ActiveTorrents {
			torrentclient.RemoveTorrent(torrent.Torrent.InfoHash().String())
			log.Println("Deleted torrent:", torrent.Torrent.InfoHash().String())
		}

		io.WriteString(w, deletedAllTorrents())
	}
}

func deletedTorrent() string {
	message := MessageResponse{
		Success: true,
		Message: "Torrent deleted.",
	}

	messageString, _ := json.Marshal(message)

	return string(messageString)
}

func deletedAllTorrents() string {
	message := MessageResponse{
		Success: true,
		Message: "All torrents have been deleted.",
	}

	messageString, _ := json.Marshal(message)

	return string(messageString)
}

func torrentNotFound() string {
	message := MessageResponse{
		Success: false,
		Message: "Torrent not found.",
	}

	messageString, _ := json.Marshal(message)

	return string(messageString)
}

func noActiveTorrentsFound() string {
	message := MessageResponse{
		Success: false,
		Message: "No active torrents found.",
	}

	messageString, _ := json.Marshal(message)

	return string(messageString)
}
