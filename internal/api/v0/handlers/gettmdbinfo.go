package handlers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nyakaspeter/raven-torrent/internal/api/v0/responses"
	"github.com/nyakaspeter/raven-torrent/pkg/mediainfo/tmdb"
)

func TmdbDiscover() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		log.Println("Get TMDB list by genre")

		output := tmdb.Discover(vars["type"], vars["genretype"], vars["sort"], vars["date"], vars["lang"], vars["page"])
		if output != "" {
			io.WriteString(w, discoveredList(output))
		} else {
			http.Error(w, noTmdbDataFound(), http.StatusNotFound)
		}
	}
}

func TmdbSearch() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		log.Println("Get TMDB search")

		output := tmdb.Search(vars["type"], vars["lang"], vars["page"], vars["text"])
		if output != "" {
			io.WriteString(w, discoveredList(output))
		} else {
			http.Error(w, noTmdbDataFound(), http.StatusNotFound)
		}
	}
}

func TmdbInfo() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		log.Println("Get TMDB info")

		output := tmdb.GetInfo(vars["type"], vars["tmdbid"], vars["lang"])
		if output != "" {
			io.WriteString(w, discoveredList(output))
		} else {
			http.Error(w, noTmdbDataFound(), http.StatusNotFound)
		}
	}
}

func discoveredList(data string) string {
	return "{\"success\":true,\"results\":[" + data + "]}"
}

func noTmdbDataFound() string {
	message := responses.MessageResponse{
		Success: false,
		Message: "No TMDB data found.",
	}

	messageString, _ := json.Marshal(message)

	return string(messageString)
}
