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
// @Param genretype path string true "Genre ids separated by comma, or 'all' to search for all genres. Possible values: 28 (Action), 12	(Adventure), 16	(Animation), 35	(Comedy), 80 (Crime), 99 (Documentary), 18 (Drama), 10751 (Family), 14 (Fantasy), 36 (History), 27 (Horror), 10402 (Music), 9648 (Mystery), 10749 (Romance), 878 (Sci-fi), 53 (Thriller), 10752 (War), 37 (Western)" example(all)
// @Param sort path string true "Sort order. Possible values: popularity.asc, popularity.desc, release_date.asc, release_date.desc, revenue.asc, revenue.desc, primary_release_date.asc, primary_release_date.desc, original_title.asc, original_title.desc, vote_average.asc, vote_average.desc, vote_count.asc, vote_count.desc" example(popularity.desc)
// @Param date path string true "Filter and only include movies or tv shows that have a release or air date that is less than or equal to the specified value. Standard date format: YYYY-MM-DD" example(2020-01-01)
// @Param lang path string true "ISO 639-1 two-letter language code" example(en)
// @Param page path int true "Specify the page of results to query" example(1)
// @Success 200 {object} TmdbMovieResultsResponse
// @Failure 404 {object} MessageResponse
func DiscoverMovies() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		log.Println("Fetching movie list:", vars)

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
// @Param genretype path string true "Genre ids separated by comma, or 'all' to search for all genres. Possible values: 10759 (Action & Adventure), 16 (Animation), 35 (Comedy), 80 (Crime), 99 (Documentary), 18 (Drama), 10751 (Family), 10762 (Kids), 9648 (Mystery), 10763 (News), 10764 (Reality), 10765 (Sci-fi & Fantasy), 10766 (Soap), 10767 (Talk), 10768 (War & Politics), 37 (Western)" example(all)
// @Param sort path string true "Sort order. Possible values: popularity.asc, popularity.desc, release_date.asc, release_date.desc, revenue.asc, revenue.desc, primary_release_date.asc, primary_release_date.desc, original_title.asc, original_title.desc, vote_average.asc, vote_average.desc, vote_count.asc, vote_count.desc" example(popularity.desc)
// @Param date path string true "Filter and only include movies or tv shows that have a release or air date that is less than or equal to the specified value. Standard date format: YYYY-MM-DD" example(2020-01-01)
// @Param lang path string true "ISO 639-1 two-letter language code" example(en)
// @Param page path int true "Specify the page of results to query" example(1)
// @Success 200 {object} TmdbShowResultsResponse
// @Failure 404 {object} MessageResponse
func DiscoverShows() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		log.Println("Fetching show list:", vars)

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

	log.Println("Found", len(results.Results), "movies.")

	json, _ := json.Marshal(response)
	return string(json)
}

func showResultsList(results mediainfotypes.ShowResults) string {
	response := TmdbShowResultsResponse{
		Success: true,
		Results: results,
	}

	log.Println("Found", len(results.Results), "shows.")

	json, _ := json.Marshal(response)
	return string(json)
}

func noTmdbDataFound() string {
	message := MessageResponse{
		Success: false,
		Message: "No TMDB data found.",
	}

	messageString, _ := json.Marshal(message)

	log.Println("No TMDB data found.")

	return string(messageString)
}
