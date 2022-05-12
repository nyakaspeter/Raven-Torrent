package handlers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/nyakaspeter/raven-torrent/internal/api/v0/responses"
	"github.com/nyakaspeter/raven-torrent/internal/torrentclient"
)

func RestartTorrentClient(procRestart chan []int64) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		for _, torrent := range torrentclient.ActiveTorrents {
			log.Println("Delete torrent:", torrent.Torrent.InfoHash().String())
			torrentclient.StopAllFileDownload(torrent.Torrent.Files())
			torrent.Torrent.Drop()
			delete(torrentclient.ActiveTorrents, torrent.Torrent.InfoHash().String())
		}

		_, err := io.WriteString(w, torrentClientRestarting())
		if err == nil {
			go func() {
				time.Sleep(1 * time.Nanosecond)
				dr, _ := strconv.ParseInt(vars["downrate"], 10, 64)
				ur, _ := strconv.ParseInt(vars["uprate"], 10, 64)
				procRestart <- []int64{dr, ur}
			}()
		} else {
			go func() {
				time.Sleep(1 * time.Nanosecond)
				dr, _ := strconv.ParseInt(vars["downrate"], 10, 64)
				ur, _ := strconv.ParseInt(vars["uprate"], 10, 64)
				procRestart <- []int64{dr, ur}
			}()
		}
	}
}

func torrentClientRestarting() string {
	message := responses.MessageResponse{
		Success: true,
		Message: "Restarting torrent client.",
	}

	messageString, _ := json.Marshal(message)

	return string(messageString)
}
