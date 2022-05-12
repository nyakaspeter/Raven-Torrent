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

type MovieMagnetLinksResponse struct {
	Success bool                  `json:"success"`
	Results []output.MovieTorrent `json:"results"`
}

func GetMovieTorrentsByImdb() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		log.Printf("Getting movie magnet link by this imdb id: %v\n", vars["imdb"])

		output := torrents.GetMovieTorrents(vars["imdb"], "", strings.Split(vars["providers"], ","))
		if len(output) > 0 {
			io.WriteString(w, movieTorrentsList(output))
			log.Printf("Magnet link found.\n")
		} else {
			http.Error(w, noMovieTorrentsFound(), http.StatusNotFound)
			log.Printf("Not found any magnet link.\n")
		}
	}
}

func GetMovieTorrentsByQuery() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		log.Printf("Getting movie magnet link by this query: %v\n", vars["query"])

		output := torrents.GetMovieTorrents("", vars["query"], strings.Split(vars["providers"], ","))
		if len(output) > 0 {
			io.WriteString(w, movieTorrentsList(output))
			log.Printf("Magnet link found.\n")
		} else {
			http.Error(w, noMovieTorrentsFound(), http.StatusNotFound)
			log.Printf("Not found any magnet link.\n")
		}
	}
}

func GetMovieTorrentsByImdbAndQuery() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		log.Printf("Getting movie magnet link by this imdb id: %v, query: %v\n", vars["imdb"], vars["query"])

		output := torrents.GetMovieTorrents(vars["imdb"], vars["query"], strings.Split(vars["providers"], ","))
		if len(output) > 0 {
			io.WriteString(w, movieTorrentsList(output))
			log.Printf("Magnet link found.\n")
		} else {
			http.Error(w, noMovieTorrentsFound(), http.StatusNotFound)
			log.Printf("Not found any magnet link.\n")
		}
	}
}

func movieTorrentsList(results []output.MovieTorrent) string {
	message := MovieMagnetLinksResponse{
		Success: true,
		Results: results,
	}

	messageString, _ := json.Marshal(message)

	return string(messageString)
}

func noMovieTorrentsFound() string {
	message := responses.MessageResponse{
		Success: false,
		Message: "No torrents found.",
	}

	messageString, _ := json.Marshal(message)

	return string(messageString)
}
