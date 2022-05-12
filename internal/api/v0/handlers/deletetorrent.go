package handlers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nyakaspeter/raven-torrent/internal/api/v0/responses"
	"github.com/nyakaspeter/raven-torrent/internal/torrentclient"
)

func DeleteTorrent() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		if t, ok := torrentclient.ActiveTorrents[vars["hash"]]; ok {

			log.Println("Delete torrent:", vars["hash"])
			torrentclient.StopAllFileDownload(t.Torrent.Files())
			t.Torrent.Drop()
			delete(torrentclient.ActiveTorrents, vars["hash"])

			io.WriteString(w, deletedTorrent())
		} else {
			http.Error(w, torrentNotFound(), http.StatusNotFound)
			log.Println("Torrent not found:", vars["hash"])
		}
	}
}

func DeleteAllTorrents() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if len(torrentclient.ActiveTorrents) > 0 {
			for _, torrent := range torrentclient.ActiveTorrents {
				log.Println("Delete torrent:", torrent.Torrent.InfoHash().String())
				torrentclient.StopAllFileDownload(torrent.Torrent.Files())
				torrent.Torrent.Drop()
				delete(torrentclient.ActiveTorrents, torrent.Torrent.InfoHash().String())
			}
			io.WriteString(w, deletedAllTorrents())
		} else {
			http.Error(w, noActiveTorrentsFound(), http.StatusNotFound)
			log.Println("No active torrents found.")
		}
	}
}

func deletedTorrent() string {
	message := responses.MessageResponse{
		Success: true,
		Message: "Torrent deleted.",
	}

	messageString, _ := json.Marshal(message)

	return string(messageString)
}

func deletedAllTorrents() string {
	message := responses.MessageResponse{
		Success: true,
		Message: "All torrents have been deleted.",
	}

	messageString, _ := json.Marshal(message)

	return string(messageString)
}

func torrentNotFound() string {
	message := responses.MessageResponse{
		Success: false,
		Message: "Torrent not found.",
	}

	messageString, _ := json.Marshal(message)

	return string(messageString)
}

func noActiveTorrentsFound() string {
	message := responses.MessageResponse{
		Success: false,
		Message: "No active torrents found.",
	}

	messageString, _ := json.Marshal(message)

	return string(messageString)
}
