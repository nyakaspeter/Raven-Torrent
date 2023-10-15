package v0

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/nyakaspeter/raven-torrent/internal/torrentclient"
	"github.com/nyakaspeter/raven-torrent/internal/torrentclient/types"
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

// @Router /torrents [get]
// @Summary Get list of added torrents
// @Description
// @Tags Torrent client
// @Success 200 {object} TorrentListResultsResponse
// @Failure 404 {object} MessageResponse
func GetActiveTorrents() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("Fetching active torrents.")

		at := torrentclient.GetActiveTorrents()

		if len(at) == 0 {
			http.Error(w, noActiveTorrentsFound(), http.StatusNotFound)
			return
		}

		io.WriteString(w, activeTorrentsList(at))
	}
}

func activeTorrentsList(at []types.TorrentInfo) string {
	var results []TorrentListResponse

	for _, torrent := range at {
		result := TorrentListResponse{
			Name:   torrent.Name,
			Hash:   torrent.Hash,
			Length: torrent.Length,
		}

		results = append(results, result)
	}

	message := TorrentListResultsResponse{
		Success: true,
		Results: results,
	}

	messageString, _ := json.Marshal(message)

	log.Println("Found", len(results), "active torrents.")

	return string(messageString)
}

func torrentNotFound() string {
	message := MessageResponse{
		Success: false,
		Message: "Torrent not found.",
	}

	messageString, _ := json.Marshal(message)

	log.Println("Torrent not found.")

	return string(messageString)
}
