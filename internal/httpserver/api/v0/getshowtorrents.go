package v0

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"

	"github.com/gorilla/mux"
	"github.com/nyakaspeter/raven-torrent/pkg/torrents"
	torrentsTypes "github.com/nyakaspeter/raven-torrent/pkg/torrents/types"
)

type ShowMagnetLinksResponse struct {
	Success bool                        `json:"success"`
	Results []torrentsTypes.ShowTorrent `json:"results"`
}

func GetShowTorrentsByImdb() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		log.Printf("Getting tv show magnet link by this imdb id: %v, season: %v, episode: %v\n", vars["imdb"], vars["season"], vars["episode"])

		output := torrents.GetShowTorrents(getShowParams(vars["imdb"], "", vars["season"], vars["episode"]), getSourceParams(vars["providers"]))
		if len(output) > 0 {
			io.WriteString(w, showTorrentsList(output))
			log.Printf("Magnet link found.\n")
		} else {
			http.Error(w, noShowTorrentsFound(), http.StatusNotFound)
			log.Printf("Not found any magnet link.\n")
		}
	}
}

func GetShowTorrentsByQuery() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		log.Printf("Getting tv show magnet link by this query: %v, season: %v, episode: %v\n", vars["query"], vars["season"], vars["episode"])

		output := torrents.GetShowTorrents(getShowParams("", vars["query"], vars["season"], vars["episode"]), getSourceParams(vars["providers"]))
		if len(output) > 0 {
			io.WriteString(w, showTorrentsList(output))
			log.Printf("Magnet link found.\n")
		} else {
			http.Error(w, noShowTorrentsFound(), http.StatusNotFound)
			log.Printf("Not found any magnet link.\n")
		}
	}
}

func GetShowTorrentsByImdbAndQuery() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		log.Printf("Getting tv show magnet link by this imdb id: %v, query: %v, season: %v, episode: %v\n", vars["imdb"], vars["query"], vars["season"], vars["episode"])

		output := torrents.GetShowTorrents(getShowParams(vars["imdb"], vars["query"], vars["season"], vars["episode"]), getSourceParams(vars["providers"]))
		if len(output) > 0 {
			io.WriteString(w, showTorrentsList(output))
			log.Printf("Magnet link found.\n")
		} else {
			http.Error(w, noShowTorrentsFound(), http.StatusNotFound)
			log.Printf("Not found any magnet link.\n")
		}
	}
}

func getShowParams(imdb string, query string, season string, episode string) torrentsTypes.ShowParams {
	showParams := torrentsTypes.ShowParams{}

	showParams.ImdbId = imdb
	showParams.SearchText = ""
	showParams.Season = season
	showParams.Episode = episode

	params, err := url.ParseQuery(query)
	if err == nil {
		if params["title"] != nil {
			showParams.SearchText += params["title"][0]
		}
	}

	return showParams
}

func showTorrentsList(results []torrentsTypes.ShowTorrent) string {
	message := ShowMagnetLinksResponse{
		Success: true,
		Results: results,
	}

	messageString, _ := json.Marshal(message)

	return string(messageString)
}

func noShowTorrentsFound() string {
	message := MessageResponse{
		Success: false,
		Message: "No torrents found.",
	}

	messageString, _ := json.Marshal(message)

	return string(messageString)
}
