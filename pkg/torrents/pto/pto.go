package pto

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
	Items []struct {
		Id           string `json:"id"`
		Quality      string `json:"quality"`
		SizeBytes    int64  `json:"size_bytes"`
		Lang         string `json:"language"`
		File         string `json:"file"`
		TorrentSeeds int64  `json:"torrent_seeds"`
		TorrentPeers int64  `json:"torrent_peers"`
	} `json:"items"`
	ItemsLang []struct {
		Id           string `json:"id"`
		Quality      string `json:"quality"`
		SizeBytes    int64  `json:"size_bytes"`
		Lang         string `json:"language"`
		File         string `json:"file"`
		TorrentSeeds int64  `json:"torrent_seeds"`
		TorrentPeers int64  `json:"torrent_peers"`
	} `json:"items_lang"`
}

func GetMovieTorrentsByImdbId(imdb string, ch chan<- []MovieTorrent) {
	req, err := http.NewRequest("GET", "https://api.apiumadomain.com/movie?cb=&quality=360p,480p,720p,1080p,2160p&page=1&imdb="+imdb, nil)
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

	outputMovieData := []MovieTorrent{}

	for _, thistorrent := range response.Items {
		temp := MovieTorrent{
			Hash:     thistorrent.Id,
			Quality:  thistorrent.Quality,
			Size:     strconv.FormatInt(thistorrent.SizeBytes, 10),
			Provider: "PTO",
			Lang:     thistorrent.Lang,
			Title:    RemoveFileExtension(thistorrent.File),
			Seeds:    strconv.FormatInt(thistorrent.TorrentSeeds, 10),
			Peers:    strconv.FormatInt(thistorrent.TorrentPeers, 10),
			Magnet:   GetMagnetLinkFromInfoHash(thistorrent.Id),
		}
		outputMovieData = append(outputMovieData, temp)
	}

	for _, thistorrent := range response.ItemsLang {
		temp := MovieTorrent{
			Hash:     thistorrent.Id,
			Quality:  thistorrent.Quality,
			Size:     strconv.FormatInt(thistorrent.SizeBytes, 10),
			Provider: "PTO",
			Lang:     thistorrent.Lang,
			Title:    RemoveFileExtension(thistorrent.File),
			Seeds:    strconv.FormatInt(thistorrent.TorrentSeeds, 10),
			Peers:    strconv.FormatInt(thistorrent.TorrentPeers, 10),
			Magnet:   GetMagnetLinkFromInfoHash(thistorrent.Id),
		}
		outputMovieData = append(outputMovieData, temp)
	}

	ch <- outputMovieData
}
