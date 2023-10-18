package v0

import (
	"encoding/base64"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/gorilla/mux"
	"github.com/nyakaspeter/raven-torrent/pkg/torrents"
	torrentsTypes "github.com/nyakaspeter/raven-torrent/pkg/torrents/types"
)

type MovieMagnetLinksResponse struct {
	Success bool                         `json:"success"`
	Results []torrentsTypes.MovieTorrent `json:"results"`
}

// @Router /getmoviemagnet/imdb/{imdb}/providers/{providers} [get]
// @Summary Get movie torrents by IMDB id
// @Description
// @Tags Torrent search
// @Param imdb path string true "IMDB id of the movie" example(tt0133093)
// @Param providers path string true "Torrent providers to use, separated by comma. Possible values: jackett, yts, 1337x, itorrent" example(jackett,yts)
// @Success 200 {object} MovieMagnetLinksResponse
// @Failure 404 {object} MessageResponse
func GetMovieTorrentsByImdb() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		log.Println("Searching torrents:", vars)

		output := torrents.GetMovieTorrents(getMovieParams(vars["imdb"], ""), getSourceParams(vars["providers"]))
		if len(output) > 0 {
			io.WriteString(w, movieTorrentsList(output))

		} else {
			http.Error(w, noMovieTorrentsFound(), http.StatusNotFound)
		}
	}
}

// @Router /getmoviemagnet/query/{query}/providers/{providers} [get]
// @Summary Get movie torrents by query string
// @Description
// @Tags Torrent search
// @Param query path string true "URI encoded query string. Supported parameters: title, releaseyear" example(title=The+Matrix&releaseyear=1999)
// @Param providers path string true "Torrent providers to use, separated by comma. Possible values: jackett, yts, 1337x, itorrent" example(jackett,yts)
// @Success 200 {object} MovieMagnetLinksResponse
// @Failure 404 {object} MessageResponse
func GetMovieTorrentsByQuery() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		log.Println("Searching torrents:", vars)

		output := torrents.GetMovieTorrents(getMovieParams("", vars["query"]), getSourceParams(vars["providers"]))
		if len(output) > 0 {
			io.WriteString(w, movieTorrentsList(output))
		} else {
			http.Error(w, noMovieTorrentsFound(), http.StatusNotFound)
		}
	}
}

// @Router /getmoviemagnet/imdb/{imdb}/query/{query}/providers/{providers} [get]
// @Summary Get movie torrents by IMDB id and query string
// @Description
// @Tags Torrent search
// @Param imdb path string true "IMDB id of the movie" example(tt0133093)
// @Param query path string true "URI encoded query string. Supported parameters: title, releaseyear" example(title=The+Matrix&releaseyear=1999)
// @Param providers path string true "Torrent providers to use, separated by comma. Possible values: jackett, yts, 1337x, itorrent" example(jackett,yts)
// @Success 200 {object} MovieMagnetLinksResponse
// @Failure 404 {object} MessageResponse
func GetMovieTorrentsByImdbAndQuery() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		log.Println("Searching torrents:", vars)

		output := torrents.GetMovieTorrents(getMovieParams(vars["imdb"], vars["query"]), getSourceParams(vars["providers"]))
		if len(output) > 0 {
			io.WriteString(w, movieTorrentsList(output))
		} else {
			http.Error(w, noMovieTorrentsFound(), http.StatusNotFound)
		}
	}
}

func getMovieParams(imdb string, query string) torrentsTypes.MovieParams {
	movieParams := torrentsTypes.MovieParams{}

	movieParams.ImdbId = imdb
	movieParams.SearchText = ""

	params, err := url.ParseQuery(query)
	if err == nil {
		if params["title"] != nil {
			movieParams.SearchText += params["title"][0]
		}
		if params["releaseyear"] != nil {
			movieParams.SearchText += " " + params["releaseyear"][0]
		}
	}

	return movieParams
}

func getSourceParams(providers string) torrentsTypes.SourceParams {
	sourceParams := torrentsTypes.SourceParams{}

	sources := strings.Split(providers, ",")

	for _, source := range sources {
		sourceName, sourceArgs := getSourceArgs(source)

		switch sourceName {
		case "jackett":
			sourceParams.Jackett.Enabled = true
			if len(sourceArgs) == 2 {
				sourceParams.Jackett.Address = sourceArgs[0]
				sourceParams.Jackett.ApiKey = sourceArgs[1]
			}
		case "yts":
			sourceParams.Yts.Enabled = true
		case "itorrent":
			sourceParams.Itorrent.Enabled = true
		case "1337x":
			sourceParams.X1337x.Enabled = true
		case "eztv":
			sourceParams.Eztv.Enabled = true
		}
	}

	return sourceParams
}

func getSourceArgs(source string) (string, []string) {
	split := strings.Split(source, ":")
	sourceName := strings.ToLower(split[0])
	var decodedArgs []string

	for i := 1; i < len(split); i++ {
		if split[i] == "" {
			continue
		}

		decodedArg, err := base64.StdEncoding.DecodeString(split[i])

		if err == nil {
			strArg := string(decodedArg)
			decodedArgs = append(decodedArgs, strArg)
		}
	}

	return sourceName, decodedArgs
}

func movieTorrentsList(results []torrentsTypes.MovieTorrent) string {
	message := MovieMagnetLinksResponse{
		Success: true,
		Results: results,
	}

	messageString, _ := json.Marshal(message)

	log.Println("Found", len(results), "torrents.")

	return string(messageString)
}

func noMovieTorrentsFound() string {
	message := MessageResponse{
		Success: false,
		Message: "No torrents found.",
	}

	messageString, _ := json.Marshal(message)

	log.Println("No torrents found.")

	return string(messageString)
}
