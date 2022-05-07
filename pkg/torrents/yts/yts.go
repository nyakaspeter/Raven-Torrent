package yts

import (
	"crypto/tls"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	. "github.com/silentmurdock/wrserver/pkg/torrents/output"
)

type apiMovieResponse struct {
	Status        string `json:"status"`
	StatusMessage string `json:"status_message"`
	Data          struct {
		MovieCount int64 `json:"movie_count"`
		Movies     []struct {
			TitleEnglish string `json:"title_english"`
			TitleLong    string `json:"title_long"`
			Lang         string `json:"language"`
			Torrents     []struct {
				Hash      string `json:"hash"`
				Quality   string `json:"quality"`
				SizeBytes int64  `json:"size_bytes"`
				Seeds     int64  `json:"seeds"`
				Peers     int64  `json:"peers"`
			} `json:"torrents"`
		} `json:"movies"`
	} `json:"data"`
}

func GetMovieTorrentsByImdbId(imdb string, ch chan<- []MovieTorrent) {
	req, err := http.NewRequest("GET", "https://yts.mx/api/v2/list_movies.json?query_term="+imdb, nil)
	if err != nil {
		ch <- []MovieTorrent{}
		return
	}

	//req.Header.Set("User-Agent", UserAgent)
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{Transport: tr, Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		ch <- []MovieTorrent{}
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		ch <- []MovieTorrent{}
		return
	}

	response := apiMovieResponse{}
	err = json.Unmarshal(body, &response)
	if err != nil {
		ch <- []MovieTorrent{}
		return
	}

	if response.Data.MovieCount == 0 {
		ch <- []MovieTorrent{}
		return
	}

	outputMovieData := []MovieTorrent{}

	for _, thistorrent := range response.Data.Movies[0].Torrents {
		temp := MovieTorrent{
			Hash:     thistorrent.Hash,
			Quality:  thistorrent.Quality,
			Size:     strconv.FormatInt(thistorrent.SizeBytes, 10),
			Provider: "YTS",
			Lang:     DecodeLanguage(response.Data.Movies[0].Lang, "en"),
			Title:    response.Data.Movies[0].TitleLong,
			Seeds:    strconv.FormatInt(thistorrent.Seeds, 10),
			Peers:    strconv.FormatInt(thistorrent.Peers, 10),
			Magnet:   GetMagnetLinkFromInfoHash(thistorrent.Hash),
		}
		outputMovieData = append(outputMovieData, temp)
	}

	ch <- outputMovieData
}
