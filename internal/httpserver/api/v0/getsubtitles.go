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

// @Router /subtitlesbyimdb/{imdb}/lang/{lang}/season/{season}/episode/{episode} [get]
// @Summary Get subtitles by IMDB id
// @Description
// @Tags Subtitle search
// @Param imdb path string true " "
// @Param lang path string true " "
// @Param season path int true " "
// @Param episode path int true " "
// @Success 200 {object} SubtitleFilesResultsResponse
// @Failure 404 {object} MessageResponse
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

// @Router /subtitlesbytext/{text}/lang/{lang}/season/{season}/episode/{episode} [get]
// @Summary Get subtitles by text
// @Description
// @Tags Subtitle search
// @Param text path string true " "
// @Param lang path string true " "
// @Param season path int true " "
// @Param episode path int true " "
// @Success 200 {object} SubtitleFilesResultsResponse
// @Failure 404 {object} MessageResponse
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

// @Router /subtitlesbyfile/{hash}/{base64path}/lang/{lang} [get]
// @Summary Get subtitles by file hash
// @Description
// @Tags Subtitle search
// @Param hash path string true " "
// @Param base64path path string true " "
// @Param lang path string true " "
// @Success 200 {object} SubtitleFilesResultsResponse
// @Failure 404 {object} MessageResponse
func GetSubtitlesByFileHash() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		if d, err := base64.StdEncoding.DecodeString(vars["base64path"]); err == nil {
			if t, ok := torrentclient.ActiveTorrents[vars["hash"]]; ok {
				idx := torrentclient.GetFileIndexByPath(string(d), t.Torrent.Files())
				file := t.Torrent.Files()[idx]

				path := file.DisplayPath()
				log.Println("Calculate Opensubtitles hash...")

				torrentclient.IncreaseConnections(path, t)

				fileSize := file.Length()
				fileHash := torrentclient.CalculateOpensubtitlesHash(file)
				log.Println("Opensubtitles hash calculated:", fileHash)

				//stop downloading the file when no connections left
				if torrentclient.DecreaseConnections(path, t) <= 0 {
					torrentclient.StopFileDownload(file)
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

			baseLink := "http://" + address + "/subtitle/" + base64.URLEncoding.EncodeToString([]byte(f.ZipDownloadLink)) + "/" + f.SubEncoding

			result := SubtitleFilesResponse{
				Lang:         f.ISO639,
				SubtitleName: workSubFileName,
				ReleaseName:  workMovieReleaseName,
				SubFormat:    f.SubFormat,
				SubEncoding:  f.SubEncoding,
				SubData:      baseLink + "/srt",
				VttData:      baseLink + "/vtt",
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
