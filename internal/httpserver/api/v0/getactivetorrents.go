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
		at := torrentclient.GetActiveTorrents()

		if len(at) == 0 {
			log.Println("No active torrents found.")
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

	return string(messageString)
}
