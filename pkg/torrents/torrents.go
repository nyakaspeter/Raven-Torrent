package torrents

import (
	"encoding/base64"
	"net/url"
	"sort"
	"strconv"
	"strings"

	"github.com/nyakaspeter/raven-torrent/pkg/torrents/eztv"
	"github.com/nyakaspeter/raven-torrent/pkg/torrents/itorrent"
	"github.com/nyakaspeter/raven-torrent/pkg/torrents/jackett"
	. "github.com/nyakaspeter/raven-torrent/pkg/torrents/output"
	"github.com/nyakaspeter/raven-torrent/pkg/torrents/pt"
	"github.com/nyakaspeter/raven-torrent/pkg/torrents/rarbg"
	"github.com/nyakaspeter/raven-torrent/pkg/torrents/x1337x"
	"github.com/nyakaspeter/raven-torrent/pkg/torrents/yts"
)

func GetMovieTorrents(imdbid string, query string, sources []string) []MovieTorrent {
	outputMovieData := []MovieTorrent{}

	ch := make(chan []MovieTorrent)

	counter := 0
	if imdbid != "" {
		for _, source := range sources {
			sourceName, sourceArgs := getSourceArgs(source)
			switch sourceName {
			case "jackett":
				if len(sourceArgs) == 2 {
					jackett.JackettAddress = sourceArgs[0]
					jackett.JackettKey = sourceArgs[1]
				}
				go jackett.GetMovieTorrentsByImdbId(imdbid, ch)
				counter++
			case "pt":
				go pt.GetMovieTorrentsByImdbId(imdbid, ch)
				counter++
			case "yts":
				go yts.GetMovieTorrentsByImdbId(imdbid, ch)
				counter++
			case "rarbg":
				go rarbg.GetMovieTorrentsByImdbId(imdbid, ch)
				counter++
			case "itorrent":
				go itorrent.GetMovieTorrentsByImdbId(imdbid, ch)
				counter++
			}
		}
	}

	if query != "" {
		params, err := url.ParseQuery(query)
		if err == nil {
			for _, source := range sources {
				sourceName, sourceArgs := getSourceArgs(source)
				switch sourceName {
				case "jackett":
					if len(sourceArgs) == 2 {
						jackett.JackettAddress = sourceArgs[0]
						jackett.JackettKey = sourceArgs[1]
					}
					go jackett.GetMovieTorrentsByQuery(params, ch)
					counter++
				case "1337x":
					go x1337x.GetMovieTorrentsByQuery(params, ch)
					counter++
				}
			}
		}
	}

	for counter > 0 {
		temp := <-ch
		if len(temp) > 0 {
			for _, item := range temp {
				duplicate := false
				for i, output := range outputMovieData {
					if (output.Hash != "" && strings.EqualFold(output.Hash, item.Hash)) ||
						(output.Torrent != "" && output.Provider == item.Provider && output.Title == item.Title) {
						duplicate = true
						if outputMovieData[i].Size == "0" && item.Size != "0" {
							outputMovieData[i].Size = item.Size
							outputMovieData[i].Title = item.Title
						}
					}
				}

				if !duplicate {
					outputMovieData = append(outputMovieData, item)
				}
			}
		}
		counter--
	}

	// Sort by seeds in descending order
	sort.Slice(outputMovieData, func(i, j int) bool {
		si, _ := strconv.ParseInt(outputMovieData[i].Seeds, 10, 64)
		sj, _ := strconv.ParseInt(outputMovieData[j].Seeds, 10, 64)
		return si > sj
	})

	return outputMovieData
}

func GetShowTorrents(imdbid string, query string, season string, episode string, sources []string) []ShowTorrent {
	outputShowData := []ShowTorrent{}

	ch := make(chan []ShowTorrent)

	counter := 0
	if imdbid != "" {
		for _, source := range sources {
			sourceName, sourceArgs := getSourceArgs(source)
			switch sourceName {
			case "jackett":
				if len(sourceArgs) == 2 {
					jackett.JackettAddress = sourceArgs[0]
					jackett.JackettKey = sourceArgs[1]
				}
				go jackett.GetShowTorrentsByImdbId(imdbid, season, episode, ch)
				counter++
			case "pt":
				go pt.GetShowTorrentsByImdbId(imdbid, season, episode, ch)
				counter++
			case "eztv":
				go eztv.GetShowTorrentsByImdbId(imdbid, season, episode, ch)
				counter++
			case "rarbg":
				go rarbg.GetShowTorrentsByImdbId(imdbid, season, episode, ch)
				counter++
			case "itorrent":
				go itorrent.GetShowTorrentsByImdbId(imdbid, season, episode, ch)
				counter++
			}
		}
	}

	if query != "" {
		params, err := url.ParseQuery(query)
		if err == nil {
			for _, source := range sources {
				sourceName, sourceArgs := getSourceArgs(source)
				switch sourceName {
				case "jackett":
					if len(sourceArgs) == 2 {
						jackett.JackettAddress = sourceArgs[0]
						jackett.JackettKey = sourceArgs[1]
					}
					go jackett.GetShowTorrentsByQuery(params, season, episode, ch)
					counter++
				case "1337x":
					go x1337x.GetShowTorrentsByQuery(params, season, episode, ch)
					counter++
				}
			}
		}
	}

	for counter > 0 {
		temp := <-ch
		if len(temp) > 0 {
			for _, item := range temp {
				duplicate := false
				for i, output := range outputShowData {
					if (output.Hash != "" && strings.EqualFold(output.Hash, item.Hash)) ||
						(output.Torrent != "" && output.Provider == item.Provider && output.Title == item.Title) {
						duplicate = true
						if outputShowData[i].Size == "0" && item.Size != "0" {
							outputShowData[i].Size = item.Size
							outputShowData[i].Title = item.Title
						}
					}
				}

				if !duplicate {
					outputShowData = append(outputShowData, item)
				}
			}
		}
		counter--
	}

	// Sort by seeds in descending order
	sort.Slice(outputShowData, func(i, j int) bool {
		si, _ := strconv.ParseInt(outputShowData[i].Seeds, 10, 64)
		sj, _ := strconv.ParseInt(outputShowData[j].Seeds, 10, 64)
		return si > sj
	})

	return outputShowData
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
