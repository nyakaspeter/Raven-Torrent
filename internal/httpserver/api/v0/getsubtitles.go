package v0

import (
	"encoding/base64"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"sort"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	"github.com/nyakaspeter/raven-torrent/internal/torrentclient"
	"github.com/nyakaspeter/raven-torrent/pkg/subtitles"
	subtitlestypes "github.com/nyakaspeter/raven-torrent/pkg/subtitles/types"
	"github.com/oz/osdb"
)

type SubtitleFilesResponse struct {
	Lang         string `json:"lang"`
	SubtitleName string `json:"subtitlename"`
	ReleaseName  string `json:"releasename"`
	SubFormat    string `json:"subformat"`
	SubEncoding  string `json:"subencoding"`
	SubData      string `json:"subdata"`
	VttData      string `json:"vttdata"`
}

type SubtitleFilesResultsResponse struct {
	Success bool                    `json:"success"`
	Results []SubtitleFilesResponse `json:"results"`
}

func GetSubtitlesByImdb() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		log.Println("Search subtitle by imdbid...")

		season, err := strconv.ParseInt(vars["season"], 10, 64)
		if err != nil {
			http.Error(w, noSubtitlesFound(), http.StatusNotFound)
			return
		}

		episode, err := strconv.ParseInt(vars["episode"], 10, 64)
		if err != nil {
			http.Error(w, noSubtitlesFound(), http.StatusNotFound)
			return
		}

		langs := strings.Split(vars["lang"], ",")

		var output []osdb.Subtitle
		if season == 0 && episode == 0 {
			params := subtitlestypes.MediaParams{}
			params.ImdbId = vars["imdb"]
			output = subtitles.GetSubtitles(params, langs)
		} else {
			params := subtitlestypes.MediaParams{}
			params.ImdbId = vars["imdb"]
			epParams := subtitlestypes.EpisodeParams{}
			epParams.Season = season
			epParams.Episode = episode
			output = subtitles.GetSubtitlesForEpisode(params, epParams, langs)
		}

		if len(output) == 0 {
			http.Error(w, noSubtitlesFound(), http.StatusNotFound)
			return
		}

		log.Println("Subtitle found.")
		io.WriteString(w, subtitleFilesList(r.Host, output, langs[0]))
	}
}

func GetSubtitlesByText() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		log.Println("Search subtitle by text...")

		season, err := strconv.ParseInt(vars["season"], 10, 64)
		if err != nil {
			http.Error(w, noSubtitlesFound(), http.StatusNotFound)
			return
		}

		episode, err := strconv.ParseInt(vars["episode"], 10, 64)
		if err != nil {
			http.Error(w, noSubtitlesFound(), http.StatusNotFound)
			return
		}

		langs := strings.Split(vars["lang"], ",")

		var output []osdb.Subtitle
		if season == 0 && episode == 0 {
			params := subtitlestypes.MediaParams{}
			params.Title = vars["text"]
			output = subtitles.GetSubtitles(params, langs)
		} else {
			params := subtitlestypes.MediaParams{}
			params.Title = vars["text"]
			epParams := subtitlestypes.EpisodeParams{}
			epParams.Season = season
			epParams.Episode = episode
			output = subtitles.GetSubtitlesForEpisode(params, epParams, langs)
		}

		if len(output) == 0 {
			http.Error(w, noSubtitlesFound(), http.StatusNotFound)
			return
		}

		log.Println("Subtitle found.")
		io.WriteString(w, subtitleFilesList(r.Host, output, langs[0]))
	}
}

func GetSubtitlesByFileHash() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		if d, err := base64.StdEncoding.DecodeString(vars["base64path"]); err == nil {
			if t, ok := torrentclient.ActiveTorrents[vars["hash"]]; ok {
				idx := torrentclient.GetFileByPath(string(d), t.Torrent.Files())
				file := t.Torrent.Files()[idx]

				path := file.DisplayPath()
				log.Println("Calculate Opensubtitles hash...")

				torrentclient.IncreaseFileClients(path, t)

				fileSize := file.Length()
				fileHash := torrentclient.CalculateOpensubtitlesHash(file)
				log.Println("Opensubtitles hash calculated:", fileHash)

				//stop downloading the file when no connections left
				if torrentclient.DecreaseFileClients(path, t) <= 0 {
					torrentclient.StopDownloadFile(file)
				}

				langs := strings.Split(vars["lang"], ",")

				params := subtitlestypes.MediaParams{}
				params.FileHash = fileHash
				params.FileSize = fileSize
				output := subtitles.GetSubtitles(params, langs)

				if len(output) == 0 {
					http.Error(w, noSubtitlesFound(), http.StatusNotFound)
					return
				}

				log.Println("Subtitle found.")
				io.WriteString(w, subtitleFilesList(r.Host, output, langs[0]))
			} else {
				http.Error(w, noSubtitlesFound(), http.StatusNotFound)
				log.Println("Unknown torrent:", vars["hash"])
				return
			}
		} else {
			http.Error(w, noSubtitlesFound(), http.StatusNotFound)
			log.Println(err)
			return
		}
	}
}

func noSubtitlesFound() string {
	message := MessageResponse{
		Success: false,
		Message: "No subtitles found.",
	}

	messageString, _ := json.Marshal(message)

	return string(messageString)
}

func subtitleFilesList(address string, files osdb.Subtitles, lang string) string {
	sortSubtitleFiles(files, lang)

	var results []SubtitleFilesResponse

	for _, f := range files {
		if f.SubFormat == "srt" {
			workSubFileName := strings.ReplaceAll(f.SubFileName, "\"", "")
			workSubFileName = strings.ReplaceAll(workSubFileName, "\\", "")

			workMovieReleaseName := strings.ReplaceAll(f.MovieReleaseName, "\"", "")
			workMovieReleaseName = strings.ReplaceAll(workMovieReleaseName, "\\", "")

			result := SubtitleFilesResponse{
				Lang:         f.ISO639,
				SubtitleName: workSubFileName,
				ReleaseName:  workMovieReleaseName,
				SubFormat:    f.SubFormat,
				SubEncoding:  f.SubEncoding,
				SubData:      "http://" + address + "/api/v0/getsubtitle/" + base64.URLEncoding.EncodeToString([]byte(f.ZipDownloadLink)) + "/encode/" + f.SubEncoding + "/subtitle.srt",
				VttData:      "http://" + address + "/api/v0/getsubtitle/" + base64.URLEncoding.EncodeToString([]byte(f.ZipDownloadLink)) + "/encode/" + f.SubEncoding + "/subtitle.vtt",
			}

			results = append(results, result)
		}
	}

	message := SubtitleFilesResultsResponse{
		Success: true,
		Results: results,
	}

	messageString, _ := json.Marshal(message)

	return string(messageString)
}

func sortSubtitleFiles(files osdb.Subtitles, lang string) {
	sort.Slice(files, func(i, j int) bool {
		return files[i].SubLanguageID == lang
	})
}
