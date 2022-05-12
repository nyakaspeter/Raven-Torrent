package handlers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/nyakaspeter/raven-torrent/internal/torrentclient"
)

type TorrentListResponse struct {
	Name   string `json:"name"`
	Hash   string `json:"hash"`
	Length string `json:"length"`
}

type TorrentListResultsResponse struct {
	Success bool                  `json:"success"`
	Results []TorrentListResponse `json:"results"`
}

func GetActiveTorrents() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if len(torrentclient.ActiveTorrents) > 0 {
			io.WriteString(w, activeTorrentsList())
		} else {
			http.Error(w, noActiveTorrentsFound(), http.StatusNotFound)
			log.Println("No active torrents found.")
		}
	}
}

func activeTorrentsList() string {
	var results []TorrentListResponse

	for _, torrent := range torrentclient.ActiveTorrents {
		result := TorrentListResponse{
			Name:   torrent.Torrent.Name(),
			Hash:   torrent.Torrent.InfoHash().String(),
			Length: strconv.FormatInt(torrent.Torrent.Length(), 10),
		}

		results = append(results, result)
	}

	message := TorrentListResultsResponse{
		Success: true,
		Results: results,
	}

	messageString, _ := json.Marshal(message)

	return string(messageString)
}
