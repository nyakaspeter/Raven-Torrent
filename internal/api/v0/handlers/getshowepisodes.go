package handlers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nyakaspeter/raven-torrent/internal/api/v0/responses"
	"github.com/nyakaspeter/raven-torrent/pkg/mediainfo/tvmaze"
)

func GetShowEpisodesByImdb() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		log.Println("Get TVMaze episodes")

		output := tvmaze.GetEpisodes("", vars["imdb"])
		if output != "" {
			io.WriteString(w, showEpisodeList(output))
		} else {
			http.Error(w, noTvMazeDataFound(), http.StatusNotFound)
		}
	}
}

func GetShowEpisodesByTvdb() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		log.Println("Get TVMaze episodes")

		output := tvmaze.GetEpisodes(vars["tvdb"], "")
		if output != "" {
			io.WriteString(w, showEpisodeList(output))
		} else {
			http.Error(w, noTvMazeDataFound(), http.StatusNotFound)
		}
	}
}

func GetShowEpisodesByImdbAndTvdb() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		log.Println("Get TVMaze episodes")

		output := tvmaze.GetEpisodes(vars["tvdb"], vars["imdb"])
		if output != "" {
			io.WriteString(w, showEpisodeList(output))
		} else {
			http.Error(w, noTvMazeDataFound(), http.StatusNotFound)
		}
	}
}

func showEpisodeList(data string) string {
	return "{\"success\":true,\"results\":" + data + "}"
}

func noTvMazeDataFound() string {
	message := responses.MessageResponse{
		Success: false,
		Message: "No TVMaze data found.",
	}

	messageString, _ := json.Marshal(message)

	return string(messageString)
}
