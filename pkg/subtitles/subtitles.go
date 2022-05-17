package subtitles

import (
	"archive/zip"
	"bytes"
	"crypto/tls"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"strings"

	subs "github.com/martinlindhe/subtitles"
	"github.com/nyakaspeter/raven-torrent/internal/settings"
	"github.com/nyakaspeter/raven-torrent/pkg/subtitles/types"
	"github.com/nyakaspeter/raven-torrent/pkg/utils"
	"github.com/oz/osdb"
)

func GetSubtitles(movie types.MediaParams, languages []string) []osdb.Subtitle {
	c, err := osdb.NewClient()
	if err != nil {
		return []osdb.Subtitle{}
	}

	c.UserAgent = *settings.OpenSubtitlesUserAgent

	if err = c.LogIn("", "", ""); err != nil {
		return []osdb.Subtitle{}
	}

	// Fallback language always English
	if len(languages) == 0 {
		languages = append(languages, "eng")
	}

	params := []interface{}{}
	if movie.FileHash != "" && movie.FileSize != 0 {
		params = []interface{}{
			c.Token,
			[]struct {
				Hash  string `xmlrpc:"moviehash"`
				Size  int64  `xmlrpc:"moviebytesize"`
				Langs string `xmlrpc:"sublanguageid"`
			}{{
				movie.FileHash,
				movie.FileSize,
				strings.Join(languages, ","),
			}},
		}
	} else if movie.ImdbId != "" {
		params = []interface{}{
			c.Token,
			[]struct {
				Imdb  string `xmlrpc:"imdbid"`
				Langs string `xmlrpc:"sublanguageid"`
			}{{
				strings.TrimPrefix(movie.ImdbId, "tt"),
				strings.Join(languages, ","),
			}},
		}
	} else if movie.Title != "" {
		params = []interface{}{
			c.Token,
			[]struct {
				Query string `xmlrpc:"query"`
				Langs string `xmlrpc:"sublanguageid"`
			}{{
				movie.Title,
				strings.Join(languages, ","),
			}},
		}
	}

	res, err := c.SearchSubtitles(&params)
	if err != nil {
		return []osdb.Subtitle{}
	}

	foundSrt := false
	for _, f := range res {
		if f.SubFormat == "srt" {
			foundSrt = true
			break
		}
	}

	if !foundSrt {
		return []osdb.Subtitle{}
	}

	return res
}

func GetSubtitlesForEpisode(show types.MediaParams, episode types.EpisodeParams, languages []string) []osdb.Subtitle {
	c, err := osdb.NewClient()
	if err != nil {
		return []osdb.Subtitle{}
	}

	c.UserAgent = *settings.OpenSubtitlesUserAgent

	if err = c.LogIn("", "", ""); err != nil {
		return []osdb.Subtitle{}
	}

	// Fallback language always English
	if len(languages) == 0 {
		languages = append(languages, "eng")
	}

	params := []interface{}{}
	if show.ImdbId != "" {
		params = []interface{}{
			c.Token,
			[]struct {
				Imdb    string `xmlrpc:"imdbid"`
				Langs   string `xmlrpc:"sublanguageid"`
				Season  int64  `xmlrpc:"season"`
				Episode int64  `xmlrpc:"episode"`
			}{{
				strings.TrimPrefix(show.ImdbId, "tt"),
				strings.Join(languages, ","),
				episode.Season,
				episode.Episode,
			}},
		}
	} else if show.Title != "" {
		params = []interface{}{
			c.Token,
			[]struct {
				Query   string `xmlrpc:"query"`
				Langs   string `xmlrpc:"sublanguageid"`
				Season  int64  `xmlrpc:"season"`
				Episode int64  `xmlrpc:"episode"`
			}{{
				show.Title,
				strings.Join(languages, ","),
				episode.Season,
				episode.Episode,
			}},
		}
	}

	res, err := c.SearchSubtitles(&params)
	if err != nil {
		return []osdb.Subtitle{}
	}

	foundSrt := false
	for _, f := range res {
		if f.SubFormat == "srt" {
			foundSrt = true
			break
		}
	}

	if !foundSrt {
		return []osdb.Subtitle{}
	}

	return res
}

func GetSubtitleContents(params types.SubtitleParams) types.SubtitleContents {
	zipContent, err := fetchZip(params.Url, *settings.OpenSubtitlesUserAgent)
	if err != nil {
		return types.SubtitleContents{}
	}

	contents := types.SubtitleContents{}
	for _, f := range zipContent.File {
		if strings.HasSuffix(strings.ToLower(f.Name), ".srt") {
			fileHandler, err := f.Open()
			if err != nil {
				return types.SubtitleContents{}
			}
			data, err := ioutil.ReadAll(fileHandler)
			if err != nil {
				return types.SubtitleContents{}
			}
			fileHandler.Close()

			// Remove UTF BOM
			if data[0] == 0xef && data[1] == 0xbb && data[2] == 0xbf {
				data = bytes.Trim(data, "\xef\xbb\xbf")
			}

			srt := utils.DecodeData(data, params.Encoding)

			subtitle, err := subs.NewFromSRT(srt)
			if err != nil {
				return types.SubtitleContents{}
			}

			if params.TargetType == "srt" {
				contents.Text = subtitle.RemoveAds().AsSRT()
				contents.ContentType = "text/plain; charset=utf-8"
				contents.ContentDisposition = "filename=subtitle.srt"
				return contents
			} else if params.TargetType == "vtt" {
				contents.Text = subtitle.RemoveAds().AsVTT()
				contents.ContentType = "text/vtt; charset=utf-8"
				contents.ContentDisposition = "filename=subtitle.vtt"
				return contents
			} else {
				break
			}
		}
	}
	return contents
}

func fetchZip(zipurl string, useragent string) (*zip.Reader, error) {
	req, err := http.NewRequest("GET", zipurl, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", useragent)
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{Transport: tr}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		b, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, errors.New(resp.Status)
		}
		return nil, errors.New(string(b))
	}

	buf := &bytes.Buffer{}

	_, err = io.Copy(buf, resp.Body)
	if err != nil {
		return nil, err
	}

	b := bytes.NewReader(buf.Bytes())
	return zip.NewReader(b, int64(b.Len()))
}
