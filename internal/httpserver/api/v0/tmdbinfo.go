package v0

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/nyakaspeter/raven-torrent/pkg/mediainfo"
	mediainfotypes "github.com/nyakaspeter/raven-torrent/pkg/mediainfo/types"
)

type TmdbMovieInfoResponse struct {
	Success bool                     `json:"success"`
	Result  mediainfotypes.MovieInfo `json:"result"`
}

type TmdbShowInfoResponse struct {
	Success bool                    `json:"success"`
	Result  mediainfotypes.ShowInfo `json:"result"`
}

// @Router /tmdbinfo/type/movie/tmdbid/{tmdbid}/lang/{lang} [get]
// @Summary Get movie details
// @Description
// @Tags Media search
// @Param tmdbid path string true "TMDB id of the movie" example(603)
// @Param lang path string true "ISO 639-1 two-letter language code" example(en)
// @Success 200 {object} TmdbMovieInfoResponse
// @Failure 404 {object} MessageResponse
func GetMovieInfo() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		log.Println("Fetching movie info:", vars)

		tmdbid, err := strconv.Atoi(vars["tmdbid"])
		if err != nil {
			http.Error(w, noTmdbDataFound(), http.StatusNotFound)
			return
		}

		result := mediainfo.GetMovieInfo(tmdbid, vars["lang"])
		if result.Id == 0 {
			http.Error(w, noTmdbDataFound(), http.StatusNotFound)
			return
		}

		io.WriteString(w, movieInfo(result))
	}
}

// @Router /tmdbinfo/type/tv/tmdbid/{tmdbid}/lang/{lang} [get]
// @Summary Get show details
// @Description
// @Tags Media search
// @Param tmdbid path string true "TMDB id of the show" example(66732)
// @Param lang path string true "ISO 639-1 two-letter language code" example(en)
// @Success 200 {object} TmdbShowInfoResponse
// @Failure 404 {object} MessageResponse
func GetShowInfo() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		log.Println("Fetching show info:", vars)

		tmdbid, err := strconv.Atoi(vars["tmdbid"])
		if err != nil {
			http.Error(w, noTmdbDataFound(), http.StatusNotFound)
			return
		}

		result := mediainfo.GetShowInfo(tmdbid, vars["lang"])
		if result.Id == 0 {
			http.Error(w, noTmdbDataFound(), http.StatusNotFound)
			return
		}

		io.WriteString(w, showInfo(result))
	}
}

func movieInfo(result mediainfotypes.MovieInfo) string {
	response := TmdbMovieInfoResponse{
		Success: true,
		Result:  result,
	}

	log.Println("Returning movie info.")

	json, _ := json.Marshal(response)
	return string(json)
}

func showInfo(result mediainfotypes.ShowInfo) string {
	response := TmdbShowInfoResponse{
		Success: true,
		Result:  result,
	}

	log.Println("Returning show info.")

	json, _ := json.Marshal(response)
	return string(json)
}
