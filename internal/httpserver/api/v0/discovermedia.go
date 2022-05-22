package v0

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	"github.com/nyakaspeter/raven-torrent/pkg/mediainfo"
	mediainfotypes "github.com/nyakaspeter/raven-torrent/pkg/mediainfo/types"
)

type TmdbMovieResultsResponse struct {
	Success bool                        `json:"success"`
	Results mediainfotypes.MovieResults `json:"results"`
}

type TmdbShowResultsResponse struct {
	Success bool                       `json:"success"`
	Results mediainfotypes.ShowResults `json:"results"`
}

// @Router /tmdbdiscover/type/movie/genretype/{genretype}/sort/{sort}/date/{date}/lang/{lang}/page/{page} [get]
// @Summary Discover movies by genre
// @Description
// @Tags Media search
// @Param genretype path string true " "
// @Param sort path string true " "
// @Param date path string true " "
// @Param lang path string true " "
// @Param page path int true " "
// @Success 200 {object} TmdbMovieResultsResponse
// @Failure 404 {object} MessageResponse
func DiscoverMovies() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		log.Println("Get TMDB list by genre")

		page, err := strconv.Atoi(vars["page"])
		if err != nil {
			http.Error(w, noTmdbDataFound(), http.StatusNotFound)
			return
		}

		genreIds := []int{}
		if vars["genretype"] != "all" {
			genreIdStrings := strings.Split(vars["genretype"], ",")
			genreIds, err = sliceAtoi(genreIdStrings)
			if err != nil {
				http.Error(w, noTmdbDataFound(), http.StatusNotFound)
				return
			}
		}

		params := mediainfotypes.MovieDiscoverParams{}
		params.SortBy = vars["sort"]
		params.MaxReleaseDate = vars["date"]
		params.GenreIds = genreIds

		results := mediainfo.DiscoverMovies(params, vars["lang"], page)
		if results.TotalResults == 0 {
			http.Error(w, noTmdbDataFound(), http.StatusNotFound)
			return
		}

		io.WriteString(w, movieResultsList(results))
	}
}

// @Router /tmdbdiscover/type/tv/genretype/{genretype}/sort/{sort}/date/{date}/lang/{lang}/page/{page} [get]
// @Summary Discover shows by genre
// @Description
// @Tags Media search
// @Param genretype path string true " "
// @Param sort path string true " "
// @Param date path string true " "
// @Param lang path string true " "
// @Param page path int true " "
// @Success 200 {object} TmdbShowResultsResponse
// @Failure 404 {object} MessageResponse
func DiscoverShows() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		log.Println("Get TMDB list by genre")

		page, err := strconv.Atoi(vars["page"])
		if err != nil {
			http.Error(w, noTmdbDataFound(), http.StatusNotFound)
			return
		}

		genreIds := []int{}
		if vars["genretype"] != "all" {
			genreIdStrings := strings.Split(vars["genretype"], ",")
			genreIds, err = sliceAtoi(genreIdStrings)
			if err != nil {
				http.Error(w, noTmdbDataFound(), http.StatusNotFound)
				return
			}
		}

		params := mediainfotypes.ShowDiscoverParams{}
		params.SortBy = vars["sort"]
		params.MaxAirDate = vars["date"]
		params.GenreIds = genreIds

		results := mediainfo.DiscoverShows(params, vars["lang"], page)
		if results.TotalResults == 0 {
			http.Error(w, noTmdbDataFound(), http.StatusNotFound)
			return
		}

		io.WriteString(w, showResultsList(results))
	}
}

func sliceAtoi(sa []string) ([]int, error) {
	si := make([]int, 0, len(sa))
	for _, a := range sa {
		i, err := strconv.Atoi(a)
		if err != nil {
			return si, err
		}
		si = append(si, i)
	}
	return si, nil
}

func movieResultsList(results mediainfotypes.MovieResults) string {
	response := TmdbMovieResultsResponse{
		Success: true,
		Results: results,
	}

	json, _ := json.Marshal(response)
	return string(json)
}

func showResultsList(results mediainfotypes.ShowResults) string {
	response := TmdbShowResultsResponse{
		Success: true,
		Results: results,
	}

	json, _ := json.Marshal(response)
	return string(json)
}

func noTmdbDataFound() string {
	message := MessageResponse{
		Success: false,
		Message: "No TMDB data found.",
	}

	messageString, _ := json.Marshal(message)

	return string(messageString)
}
