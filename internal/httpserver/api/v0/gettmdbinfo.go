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

func TmdbDiscover() func(w http.ResponseWriter, r *http.Request) {
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

		output := ""

		switch vars["type"] {
		case "movie":
			params := mediainfotypes.MovieDiscoverParams{}
			params.SortBy = vars["sort"]
			params.MaxReleaseDate = vars["date"]
			params.GenreIds = genreIds

			results := mediainfo.DiscoverMovies(params, vars["lang"], page)
			if results.TotalResults != 0 {
				resultsJson, _ := json.Marshal(results)
				output = string(resultsJson)
			}
		case "tv":
			params := mediainfotypes.ShowDiscoverParams{}
			params.SortBy = vars["sort"]
			params.MaxAirDate = vars["date"]
			params.GenreIds = genreIds

			results := mediainfo.DiscoverShows(params, vars["lang"], page)
			if results.TotalResults != 0 {
				resultsJson, _ := json.Marshal(results)
				output = string(resultsJson)
			}
		}

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

		page, err := strconv.Atoi(vars["page"])
		if err != nil {
			http.Error(w, noTmdbDataFound(), http.StatusNotFound)
			return
		}

		output := ""

		switch vars["type"] {
		case "movie":
			results := mediainfo.SearchMovies(vars["text"], vars["lang"], page)
			if results.TotalResults != 0 {
				resultsJson, _ := json.Marshal(results)
				output = string(resultsJson)
			}
		case "tv":
			results := mediainfo.SearchShows(vars["text"], vars["lang"], page)
			if results.TotalResults != 0 {
				resultsJson, _ := json.Marshal(results)
				output = string(resultsJson)
			}
		}

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

		tmdbid, err := strconv.Atoi(vars["tmdbid"])
		if err != nil {
			http.Error(w, noTmdbDataFound(), http.StatusNotFound)
			return
		}

		output := ""

		switch vars["type"] {
		case "movie":
			result := mediainfo.GetMovieInfo(tmdbid, vars["lang"])
			if result.Id != 0 {
				resultsJson, _ := json.Marshal(result)
				output = string(resultsJson)
			}
		case "tv":
			result := mediainfo.GetShowInfo(tmdbid, vars["lang"])
			if result.Id != 0 {
				resultsJson, _ := json.Marshal(result)
				output = string(resultsJson)
			}
		}

		if output != "" {
			io.WriteString(w, discoveredList(output))
		} else {
			http.Error(w, noTmdbDataFound(), http.StatusNotFound)
		}
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

func discoveredList(data string) string {
	return "{\"success\":true,\"results\":[" + data + "]}"
}

func noTmdbDataFound() string {
	message := MessageResponse{
		Success: false,
		Message: "No TMDB data found.",
	}

	messageString, _ := json.Marshal(message)

	return string(messageString)
}
