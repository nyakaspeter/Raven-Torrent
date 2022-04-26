package providers

import (
	"net/url"
	"sort"
	"strconv"
	"strings"

	"github.com/silentmurdock/wrserver/providers/eztv"
	"github.com/silentmurdock/wrserver/providers/itorrent"
	"github.com/silentmurdock/wrserver/providers/jackett"
	out "github.com/silentmurdock/wrserver/providers/output"
	"github.com/silentmurdock/wrserver/providers/pt"
	"github.com/silentmurdock/wrserver/providers/pto"
	"github.com/silentmurdock/wrserver/providers/rarbg"
	"github.com/silentmurdock/wrserver/providers/tmdb"
	"github.com/silentmurdock/wrserver/providers/tvmaze"
	"github.com/silentmurdock/wrserver/providers/x1337x"
	"github.com/silentmurdock/wrserver/providers/yts"
)

func GetMovieMagnet(imdbid string, query string, sources []string) []out.OutputMovieStruct {
	outputMovieData := []out.OutputMovieStruct{}

	ch := make(chan []out.OutputMovieStruct)

	counter := 0
	if imdbid != "" {
		for _, source := range sources {
			switch strings.ToLower(source) {
			case "jackett":
				go jackett.GetMovieMagnetByImdb(imdbid, ch)
				counter++
			case "pt":
				go pt.GetMovieMagnetByImdb(imdbid, ch)
				counter++
			case "yts":
				go yts.GetMovieMagnetByImdb(imdbid, ch)
				counter++
			case "rarbg":
				go rarbg.GetMovieMagnetByImdb(imdbid, ch)
				counter++
			case "pto":
				go pto.GetMovieMagnetByImdb(imdbid, ch)
				counter++
			case "itorrent":
				go itorrent.GetMovieMagnetByImdb(imdbid, ch)
				counter++
			}
		}
	}

	if query != "" {
		params, err := url.ParseQuery(query)
		if err == nil {
			for _, source := range sources {
				switch strings.ToLower(source) {
				case "jackett":
					go jackett.GetMovieMagnetByQuery(params, ch)
					counter++
				case "1337x":
					go x1337x.GetMovieMagnetByQuery(params, ch)
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

func GetShowMagnet(imdbid string, query string, season string, episode string, sources []string) []out.OutputShowStruct {
	outputShowData := []out.OutputShowStruct{}

	ch := make(chan []out.OutputShowStruct)

	counter := 0
	if imdbid != "" {
		for _, source := range sources {
			switch strings.ToLower(source) {
			case "jackett":
				go jackett.GetShowMagnetByImdb(imdbid, season, episode, ch)
				counter++
			case "pt":
				go pt.GetShowMagnetByImdb(imdbid, season, episode, ch)
				counter++
			case "eztv":
				go eztv.GetShowMagnetByImdb(imdbid, season, episode, ch)
				counter++
			case "rarbg":
				go rarbg.GetShowMagnetByImdb(imdbid, season, episode, ch)
				counter++
			case "itorrent":
				go itorrent.GetShowMagnetByImdb(imdbid, season, episode, ch)
				counter++
			}
		}
	}

	if query != "" {
		params, err := url.ParseQuery(query)
		if err == nil {
			for _, source := range sources {
				switch strings.ToLower(source) {
				case "jackett":
					go jackett.GetShowMagnetByQuery(params, season, episode, ch)
					counter++
				case "1337x":
					go x1337x.GetShowMagnetByQuery(params, season, episode, ch)
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

func SetTMDBKey(tmdbKey string) {
	tmdb.SetTMDBKey(tmdbKey)
}

func SetJackettAddressAndKey(jackettAddress string, jackettKey string) {
	jackett.SetJackettAddressAndKey(jackettAddress, jackettKey)
}

func MirrorTmdbDiscover(qtype string, genretype string, sort string, date string, lang string, cpage string) string {
	return tmdb.MirrorTmdbDiscover(qtype, genretype, sort, date, lang, cpage)
}

func MirrorTmdbSearch(qtype string, lang string, cpage string, typedtext string) string {
	return tmdb.MirrorTmdbSearch(qtype, lang, cpage, typedtext)
}

func MirrorTmdbInfo(qtype string, tmdbid string, lang string) string {
	return tmdb.MirrorTmdbInfo(qtype, tmdbid, lang)
}

func GetTvMazeEpisodes(tvdb string, imdb string) string {
	return tvmaze.GetTvMazeEpisodes(tvdb, imdb)
}
