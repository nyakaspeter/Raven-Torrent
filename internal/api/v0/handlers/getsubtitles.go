package handlers

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
	"github.com/nyakaspeter/raven-torrent/internal/api/v0/responses"
	"github.com/nyakaspeter/raven-torrent/internal/torrentclient"
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

// TODO: Extract to package

func GetSubtitlesByImdb(useragent string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		// Create Opensubtitles client
		c, err := osdb.NewClient()
		if err != nil {
			http.Error(w, failedToConnectToOpenSubtitles(), http.StatusNotFound)
			return
		}

		c.UserAgent = useragent

		// Anonymous Login with UserAgent string will set c.Token when successful
		if err = c.LogIn("", "", ""); err != nil {
			http.Error(w, failedToConnectToOpenSubtitles(), http.StatusNotFound)
			return
		}

		ids := make([]string, 1)
		ids[0] = strings.TrimPrefix(vars["imdb"], "tt")
		langs := strings.Split(vars["lang"], ",")

		// Fallback language always English
		if len(langs) == 0 {
			langs = append(langs, "eng")
		}

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

		params := []interface{}{}
		if season == 0 && episode == 0 {
			params = []interface{}{
				c.Token,
				[]struct {
					Imdb  string `xmlrpc:"imdbid"`
					Langs string `xmlrpc:"sublanguageid"`
				}{{
					ids[0],
					strings.Join(langs, ","),
				}},
			}
		} else {
			params = []interface{}{
				c.Token,
				[]struct {
					Imdb    string `xmlrpc:"imdbid"`
					Langs   string `xmlrpc:"sublanguageid"`
					Season  int64  `xmlrpc:"season"`
					Episode int64  `xmlrpc:"episode"`
				}{{
					ids[0],
					strings.Join(langs, ","),
					season,
					episode,
				}},
			}
		}

		res, err := c.SearchSubtitles(&params)
		if err != nil {
			http.Error(w, noSubtitlesFound(), http.StatusNotFound)
			return
		}

		if len(res) == 0 {
			http.Error(w, noSubtitlesFound(), http.StatusNotFound)
			return
		}

		found := false
		for _, f := range res {
			if f.SubFormat == "srt" {
				found = true
				break
			}
		}

		if !found {
			http.Error(w, noSubtitlesFound(), http.StatusNotFound)
			return
		}

		log.Println("Subtitle found.")
		io.WriteString(w, subtitleFilesList(r.Host, res, langs[0]))
	}
}

func GetSubtitlesByText(useragent string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		// Create Opensubtitles client
		c, err := osdb.NewClient()
		if err != nil {
			http.Error(w, failedToConnectToOpenSubtitles(), http.StatusNotFound)
			return
		}

		c.UserAgent = useragent

		// Anonymous Login with UserAgent string will set c.Token when successful
		if err = c.LogIn("", "", ""); err != nil {
			http.Error(w, failedToConnectToOpenSubtitles(), http.StatusNotFound)
			return
		}

		text := vars["text"]
		langs := strings.Split(vars["lang"], ",")

		// Fallback language always English
		if len(langs) == 0 {
			langs = append(langs, "eng")
		}

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

		params := []interface{}{}
		if season == 0 && episode == 0 {
			params = []interface{}{
				c.Token,
				[]struct {
					Query string `xmlrpc:"query"`
					Langs string `xmlrpc:"sublanguageid"`
				}{{
					text,
					strings.Join(langs, ","),
				}},
			}
		} else {
			params = []interface{}{
				c.Token,
				[]struct {
					Query   string `xmlrpc:"query"`
					Langs   string `xmlrpc:"sublanguageid"`
					Season  int64  `xmlrpc:"season"`
					Episode int64  `xmlrpc:"episode"`
				}{{
					text,
					strings.Join(langs, ","),
					season,
					episode,
				}},
			}
		}

		res, err := c.SearchSubtitles(&params)
		if err != nil {
			http.Error(w, noSubtitlesFound(), http.StatusNotFound)
			return
		}

		if len(res) == 0 {
			http.Error(w, noSubtitlesFound(), http.StatusNotFound)
			return
		}

		found := false
		for _, f := range res {
			if f.SubFormat == "srt" {
				found = true
				break
			}
		}

		if !found {
			http.Error(w, noSubtitlesFound(), http.StatusNotFound)
			return
		}

		log.Println("Subtitle found.")
		io.WriteString(w, subtitleFilesList(r.Host, res, langs[0]))
	}
}

func GetSubtitlesByFileHash(useragent string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		if d, err := base64.StdEncoding.DecodeString(vars["base64path"]); err == nil {
			if t, ok := torrentclient.ActiveTorrents[vars["hash"]]; ok {

				idx := torrentclient.GetFileByPath(string(d), t.Torrent.Files())
				file := t.Torrent.Files()[idx]

				path := file.DisplayPath()
				log.Println("Calculate Opensubtitles hash...")

				torrentclient.IncreaseFileClients(path, t)

				fileHash := torrentclient.CalculateOpensubtitlesHash(file)
				log.Println("Opensubtitles hash calculated:", fileHash)

				//stop downloading the file when no connections left
				if torrentclient.DecreaseFileClients(path, t) <= 0 {
					torrentclient.StopDownloadFile(file)
				}

				// Create Opensubtitles client
				c, err := osdb.NewClient()
				if err != nil {
					http.Error(w, failedToConnectToOpenSubtitles(), http.StatusNotFound)
					return
				}

				c.UserAgent = useragent

				// Anonymous Login with UserAgent string will set c.Token when successful
				if err = c.LogIn("", "", ""); err != nil {
					http.Error(w, failedToConnectToOpenSubtitles(), http.StatusNotFound)
					return
				}

				langs := strings.Split(vars["lang"], ",")

				// Fallback language always English
				if len(langs) == 0 {
					langs = append(langs, "eng")
				}

				params := []interface{}{
					c.Token,
					[]struct {
						Hash  string `xmlrpc:"moviehash"`
						Size  int64  `xmlrpc:"moviebytesize"`
						Langs string `xmlrpc:"sublanguageid"`
					}{{
						fileHash,
						file.Length(),
						strings.Join(langs, ","),
					}},
				}

				res, err := c.SearchSubtitles(&params)
				if err != nil {
					http.Error(w, noSubtitlesFound(), http.StatusNotFound)
					return
				}

				if len(res) == 0 {
					http.Error(w, noSubtitlesFound(), http.StatusNotFound)
					return
				}

				found := false
				for _, f := range res {
					if f.SubFormat == "srt" {
						found = true
						break
					}
				}

				if !found {
					http.Error(w, noSubtitlesFound(), http.StatusNotFound)
					return
				}

				log.Println("Subtitle found.")
				io.WriteString(w, subtitleFilesList(r.Host, res, langs[0]))
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

func failedToConnectToOpenSubtitles() string {
	message := responses.MessageResponse{
		Success: false,
		Message: "Failed to connect to opensubtitles.",
	}

	messageString, _ := json.Marshal(message)

	return string(messageString)
}

func noSubtitlesFound() string {
	message := responses.MessageResponse{
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
				SubData:      "http://" + address + "/api/getsubtitle/" + base64.URLEncoding.EncodeToString([]byte(f.ZipDownloadLink)) + "/encode/" + f.SubEncoding + "/subtitle.srt",
				VttData:      "http://" + address + "/api/getsubtitle/" + base64.URLEncoding.EncodeToString([]byte(f.ZipDownloadLink)) + "/encode/" + f.SubEncoding + "/subtitle.vtt",
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
