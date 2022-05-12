package handlers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/nyakaspeter/raven-torrent/internal/api/v0/responses"
	"github.com/nyakaspeter/raven-torrent/pkg/torrents"
	"github.com/nyakaspeter/raven-torrent/pkg/torrents/output"
)

type ShowMagnetLinksResponse struct {
	Success bool                 `json:"success"`
	Results []output.ShowTorrent `json:"results"`
}

func GetShowTorrentsByImdb() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		log.Printf("Getting tv show magnet link by this imdb id: %v, season: %v, episode: %v\n", vars["imdb"], vars["season"], vars["episode"])

		output := torrents.GetShowTorrents(vars["imdb"], "", vars["season"], vars["episode"], strings.Split(vars["providers"], ","))
		if len(output) > 0 {
			io.WriteString(w, showTorrentsList(output))
			log.Printf("Magnet link found.\n")
		} else {
			http.Error(w, noShowTorrentsFound(), http.StatusNotFound)
			log.Printf("Not found any magnet link.\n")
		}
	}
}

func GetShowTorrentsByQuery() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		log.Printf("Getting tv show magnet link by this query: %v, season: %v, episode: %v\n", vars["query"], vars["season"], vars["episode"])

		output := torrents.GetShowTorrents("", vars["query"], vars["season"], vars["episode"], strings.Split(vars["providers"], ","))
		if len(output) > 0 {
			io.WriteString(w, showTorrentsList(output))
			log.Printf("Magnet link found.\n")
		} else {
			http.Error(w, noShowTorrentsFound(), http.StatusNotFound)
			log.Printf("Not found any magnet link.\n")
		}
	}
}

func GetShowTorrentsByImdbAndQuery() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		log.Printf("Getting tv show magnet link by this imdb id: %v, query: %v, season: %v, episode: %v\n", vars["imdb"], vars["query"], vars["season"], vars["episode"])

		output := torrents.GetShowTorrents(vars["imdb"], vars["query"], vars["season"], vars["episode"], strings.Split(vars["providers"], ","))
		if len(output) > 0 {
			io.WriteString(w, showTorrentsList(output))
			log.Printf("Magnet link found.\n")
		} else {
			http.Error(w, noShowTorrentsFound(), http.StatusNotFound)
			log.Printf("Not found any magnet link.\n")
		}
	}
}

func showTorrentsList(results []output.ShowTorrent) string {
	message := ShowMagnetLinksResponse{
		Success: true,
		Results: results,
	}

	messageString, _ := json.Marshal(message)

	return string(messageString)
}

func noShowTorrentsFound() string {
	message := responses.MessageResponse{
		Success: false,
		Message: "No torrents found.",
	}

	messageString, _ := json.Marshal(message)

	return string(messageString)
}
