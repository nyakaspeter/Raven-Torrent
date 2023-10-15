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

type ShowEpisodesResponse struct {
	Success bool                           `json:"success"`
	Results []mediainfotypes.TvMazeEpisode `json:"results"`
}

// @Router /tvmazeepisodes/imdb/{imdb} [get]
// @Summary Get show episodes by IMDB id
// @Description
// @Tags Media search
// @Param imdb path string true " "
// @Success 200 {object} ShowEpisodesResponse
// @Failure 404 {object} MessageResponse
func GetShowEpisodesByImdb() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		log.Println("Fetching show episodes:", vars)

		showIds := mediainfotypes.ShowIds{}
		showIds.ImdbId = vars["imdb"]

		episodes := mediainfo.GetShowEpisodes(showIds)
		if len(episodes) == 0 {
			http.Error(w, noTvMazeDataFound(), http.StatusNotFound)
			return
		}

		io.WriteString(w, showEpisodeList(episodes))
	}
}

// @Router /tvmazeepisodes/tvdb/{tvdb} [get]
// @Summary Get show episodes by TVDB id
// @Description
// @Tags Media search
// @Param tvdb path string true " "
// @Success 200 {object} ShowEpisodesResponse
// @Failure 404 {object} MessageResponse
func GetShowEpisodesByTvdb() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		log.Println("Fetching show episodes:", vars)

		showIds := mediainfotypes.ShowIds{}
		showIds.TvdbId = vars["tvdb"]

		episodes := mediainfo.GetShowEpisodes(showIds)
		if len(episodes) == 0 {
			http.Error(w, noTvMazeDataFound(), http.StatusNotFound)
			return
		}

		io.WriteString(w, showEpisodeList(episodes))
	}
}

// @Router /tvmazeepisodes/tvdb/{tvdb}/imdb/{imdb} [get]
// @Summary Get show episodes by TVDB id and IMDB id
// @Description
// @Tags Media search
// @Param tvdb path string true " "
// @Param imdb path string true " "
// @Success 200 {object} ShowEpisodesResponse
// @Failure 404 {object} MessageResponse
func GetShowEpisodesByImdbAndTvdb() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		log.Println("Fetching show episodes:", vars)

		showIds := mediainfotypes.ShowIds{}
		showIds.ImdbId = vars["imdb"]
		showIds.TvdbId = vars["tvdb"]

		episodes := mediainfo.GetShowEpisodes(showIds)
		if len(episodes) == 0 {
			http.Error(w, noTvMazeDataFound(), http.StatusNotFound)
			return
		}

		io.WriteString(w, showEpisodeList(episodes))
	}
}

func showEpisodeList(episodes []mediainfotypes.TvMazeEpisode) string {
	response := ShowEpisodesResponse{
		Success: true,
		Results: episodes,
	}

	log.Println("Found", len(episodes), "episodes.")

	json, _ := json.Marshal(response)
	return string(json)
}

func noTvMazeDataFound() string {
	message := MessageResponse{
		Success: false,
		Message: "No TVMaze data found.",
	}

	messageString, _ := json.Marshal(message)

	log.Println("No TVMaze data found.")

	return string(messageString)
}
