package v0

import (
	"encoding/base64"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	"github.com/nyakaspeter/raven-torrent/internal/torrentclient"
	"github.com/nyakaspeter/raven-torrent/pkg/subtitles"
	subtitlestypes "github.com/nyakaspeter/raven-torrent/pkg/subtitles/types"
)

type SubtitleFilesResultsResponse struct {
	Success bool                          `json:"success"`
	Results []subtitlestypes.SubtitleFile `json:"results"`
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

		log.Println("Searching subtitles:", vars)

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

		var output []subtitlestypes.SubtitleFile
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

		io.WriteString(w, subtitleFilesListResponse(output))
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

		log.Println("Searching subtitles:", vars)

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

		var output []subtitlestypes.SubtitleFile
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

		io.WriteString(w, subtitleFilesListResponse(output))
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

		log.Println("Searching subtitles:", vars)

		if d, err := base64.StdEncoding.DecodeString(vars["base64path"]); err == nil {
			if t, ok := torrentclient.ActiveTorrents[vars["hash"]]; ok {
				idx := torrentclient.GetFileIndexByPath(string(d), t.Torrent.Files())
				file := t.Torrent.Files()[idx]

				path := file.DisplayPath()

				torrentclient.IncreaseConnections(path, t)

				fileSize := file.Length()
				fileHash := torrentclient.CalculateOpensubtitlesHash(file)

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

				io.WriteString(w, subtitleFilesListResponse(output))
			} else {
				http.Error(w, noSubtitlesFound(), http.StatusNotFound)
				return
			}
		} else {
			http.Error(w, noSubtitlesFound(), http.StatusNotFound)
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

	log.Println("No subtitles found.")

	return string(messageString)
}

func subtitleFilesListResponse(results []subtitlestypes.SubtitleFile) string {
	message := SubtitleFilesResultsResponse{
		Success: true,
		Results: results,
	}

	messageString, _ := json.Marshal(message)

	log.Println("Found", len(results), "subtitles.")

	return string(messageString)
}
