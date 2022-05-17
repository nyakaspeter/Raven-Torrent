package v0

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nyakaspeter/raven-torrent/pkg/mediainfo"
	mediainfotypes "github.com/nyakaspeter/raven-torrent/pkg/mediainfo/types"
)

func GetShowEpisodesByImdb() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		log.Println("Get TVMaze episodes")

		showIds := mediainfotypes.ShowIds{}
		showIds.ImdbId = vars["imdb"]

		output := ""

		results := mediainfo.GetShowEpisodes(showIds)
		if len(results) > 0 {
			resultsJson, _ := json.Marshal(results)
			output = string(resultsJson)
		}

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

		showIds := mediainfotypes.ShowIds{}
		showIds.TvdbId = vars["tvdb"]

		output := ""

		results := mediainfo.GetShowEpisodes(showIds)
		if len(results) > 0 {
			resultsJson, _ := json.Marshal(results)
			output = string(resultsJson)
		}

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

		showIds := mediainfotypes.ShowIds{}
		showIds.ImdbId = vars["imdb"]
		showIds.TvdbId = vars["tvdb"]

		output := ""

		results := mediainfo.GetShowEpisodes(showIds)
		if len(results) > 0 {
			resultsJson, _ := json.Marshal(results)
			output = string(resultsJson)
		}

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
	message := MessageResponse{
		Success: false,
		Message: "No TVMaze data found.",
	}

	messageString, _ := json.Marshal(message)

	return string(messageString)
}
