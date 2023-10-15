package v0

import (
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/nyakaspeter/raven-torrent/pkg/mediainfo"
)

// @Router /tmdbsearch/type/movie/lang/{lang}/page/{page}/text/{text} [get]
// @Summary Search movies
// @Description
// @Tags Media search
// @Param text path string true " "
// @Param lang path string true " "
// @Param page path int true " "
// @Success 200 {object} TmdbMovieResultsResponse
// @Failure 404 {object} MessageResponse
func SearchMovies() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		log.Println("Searching movies:", vars)

		page, err := strconv.Atoi(vars["page"])
		if err != nil {
			http.Error(w, noTmdbDataFound(), http.StatusNotFound)
			return
		}

		results := mediainfo.SearchMovies(vars["text"], vars["lang"], page)
		if results.TotalResults == 0 {
			http.Error(w, noTmdbDataFound(), http.StatusNotFound)
			return
		}

		io.WriteString(w, movieResultsList(results))
	}
}

// @Router /tmdbsearch/type/tv/lang/{lang}/page/{page}/text/{text} [get]
// @Summary Search shows
// @Description
// @Tags Media search
// @Param text path string true " "
// @Param lang path string true " "
// @Param page path int true " "
// @Success 200 {object} TmdbShowResultsResponse
// @Failure 404 {object} MessageResponse
func SearchShows() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		log.Println("Searching shows:", vars)

		page, err := strconv.Atoi(vars["page"])
		if err != nil {
			http.Error(w, noTmdbDataFound(), http.StatusNotFound)
			return
		}

		results := mediainfo.SearchShows(vars["text"], vars["lang"], page)
		if results.TotalResults == 0 {
			http.Error(w, noTmdbDataFound(), http.StatusNotFound)
			return
		}

		io.WriteString(w, showResultsList(results))
	}
}
