package jackett

import (
	"crypto/tls"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"time"

	. "github.com/silentmurdock/wrserver/pkg/torrents/output"
)

type apiResponse struct {
	TorrentResults []struct {
		Title        string `json:"Title"`
		Tracker      string `json:"Tracker"`
		CategoryDesc string `json:"CategoryDesc"`
		PublishDate  string `json:"PublishDate"`
		Description  string `json:"Description"`
		Imdb         int64  `json:"Imdb"`
		Size         int64  `json:"Size"`
		Seeders      int64  `json:"Seeders"`
		Peers        int64  `json:"Peers"`
		Details      string `json:"Details"`
		Guid         string `json:"Guid"`
		Link         string `json:"Link"`
		MagnetUri    string `json:"MagnetUri"`
		InfoHash     string `json:"InfoHash"`
	} `json:"Results"`
}

var jackettAddress = ""
var jackettKey = ""
var movieCategories = [...]string{"2000", "2030", "2040"}
var showCategories = [...]string{"5000", "5030", "5040"}

func GetMovieTorrentsByImdbId(imdb string, ch chan<- []MovieTorrent) {
	outputMovieData := []MovieTorrent{}

	for _, category := range movieCategories {
		req, err := http.NewRequest("GET", (jackettAddress + "/api/v2.0/indexers/all/results?apikey=" + jackettKey + "&category=" + category + "&query=" + imdb), nil)
		if err != nil {
			break
		}

		//req.Header.Set("User-Agent", UserAgent)
		tr := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}

		client := &http.Client{Transport: tr, Timeout: 10 * time.Second}
		resp, err := client.Do(req)
		if err != nil {
			break
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			break
		}

		response := apiResponse{}
		err = json.Unmarshal(body, &response)
		if err != nil {
			break
		}

		if len(response.TorrentResults) == 0 {
			break
		}

		for _, thistorrent := range response.TorrentResults {
			temp := MovieTorrent{
				Hash:     GetInfoHashFromMagnetLink(thistorrent.MagnetUri),
				Quality:  GuessQualityFromString(thistorrent.Title),
				Size:     strconv.FormatInt(thistorrent.Size, 10),
				Provider: thistorrent.Tracker,
				Lang:     GuessLanguageFromString(thistorrent.Title),
				Title:    thistorrent.Title,
				Seeds:    strconv.FormatInt(thistorrent.Seeders, 10),
				Peers:    strconv.FormatInt(thistorrent.Peers, 10),
				Magnet:   thistorrent.MagnetUri,
				Torrent:  thistorrent.Link,
			}
			outputMovieData = append(outputMovieData, temp)
		}
	}

	ch <- outputMovieData
	return
}

func GetMovieTorrentsByQuery(params map[string][]string, ch chan<- []MovieTorrent) {
	// Decode params data
	query := ""

	if params["title"] != nil && params["title"][0] != "" {
		query += params["title"][0]
	} else {
		ch <- []MovieTorrent{}
		return
	}

	if params["releaseyear"] != nil {
		query += " " + params["releaseyear"][0]
	}

	query = url.QueryEscape(query)

	outputMovieData := []MovieTorrent{}

	for _, category := range movieCategories {
		req, err := http.NewRequest("GET", (jackettAddress + "/api/v2.0/indexers/all/results?apikey=" + jackettKey + "&category=" + category + "&query=" + query), nil)
		if err != nil {
			break
		}

		tr := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}

		client := &http.Client{Transport: tr, Timeout: 10 * time.Second}
		resp, err := client.Do(req)
		if err != nil {
			break
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			break
		}

		response := apiResponse{}
		err = json.Unmarshal(body, &response)
		if err != nil {
			break
		}

		if len(response.TorrentResults) == 0 {
			break
		}

		for _, thistorrent := range response.TorrentResults {
			temp := MovieTorrent{
				Hash:     GetInfoHashFromMagnetLink(thistorrent.MagnetUri),
				Quality:  GuessQualityFromString(thistorrent.Title),
				Size:     strconv.FormatInt(thistorrent.Size, 10),
				Provider: thistorrent.Tracker,
				Lang:     GuessLanguageFromString(thistorrent.Title),
				Title:    thistorrent.Title,
				Seeds:    strconv.FormatInt(thistorrent.Seeders, 10),
				Peers:    strconv.FormatInt(thistorrent.Peers, 10),
				Magnet:   thistorrent.MagnetUri,
				Torrent:  thistorrent.Link,
			}
			outputMovieData = append(outputMovieData, temp)
		}
	}

	ch <- outputMovieData
	return
}

func GetShowTorrentsByImdbId(imdb string, season string, episode string, ch chan<- []ShowTorrent) {
	outputShowData := []ShowTorrent{}

	for _, category := range showCategories {
		req, err := http.NewRequest("GET", (jackettAddress + "/api/v2.0/indexers/all/results?apikey=" + jackettKey + "&category=" + category + "&query=" + imdb), nil)
		if err != nil {
			break
		}

		tr := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}

		client := &http.Client{Transport: tr, Timeout: 10 * time.Second}
		resp, err := client.Do(req)
		if err != nil {
			break
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			break
		}

		response := apiResponse{}
		err = json.Unmarshal(body, &response)
		if err != nil {
			break
		}

		if len(response.TorrentResults) == 0 {
			break
		}

		for _, thistorrent := range response.TorrentResults {
			titleSeason, titleEpisode := GuessSeasonEpisodeNumberFromString(thistorrent.Title)

			if titleSeason == season && titleEpisode == episode {
				temp := ShowTorrent{
					Hash:     GetInfoHashFromMagnetLink(thistorrent.MagnetUri),
					Quality:  GuessQualityFromString(thistorrent.Title),
					Season:   season,
					Episode:  episode,
					Size:     strconv.FormatInt(thistorrent.Size, 10),
					Provider: thistorrent.Tracker,
					Lang:     GuessLanguageFromString(thistorrent.Title),
					Title:    thistorrent.Title,
					Seeds:    strconv.FormatInt(thistorrent.Seeders, 10),
					Peers:    strconv.FormatInt(thistorrent.Peers, 10),
					Magnet:   thistorrent.MagnetUri,
					Torrent:  thistorrent.Link,
				}
				outputShowData = append(outputShowData, temp)
			}
		}
	}

	ch <- outputShowData
	return
}

func GetShowTorrentsByQuery(params map[string][]string, season string, episode string, ch chan<- []ShowTorrent) {
	// Decode params data
	query := ""

	if params["title"] != nil && params["title"][0] != "" {
		query += params["title"][0] + " "
	} else {
		ch <- []ShowTorrent{}
		return
	}

	if season != "0" {
		if len(season) == 1 {
			query += "s0" + season
		} else {
			query += "s" + season
		}
	}
	if episode != "0" {
		if len(episode) == 1 {
			query += "e0" + episode
		} else {
			query += "e" + episode
		}
	}

	query = url.QueryEscape(query)

	outputShowData := []ShowTorrent{}

	for _, category := range showCategories {
		req, err := http.NewRequest("GET", (jackettAddress + "/api/v2.0/indexers/all/results?apikey=" + jackettKey + "&category=" + category + "&query=" + query), nil)
		if err != nil {
			break
		}

		tr := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}

		client := &http.Client{Transport: tr, Timeout: 10 * time.Second}
		resp, err := client.Do(req)
		if err != nil {
			break
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			break
		}

		response := apiResponse{}
		err = json.Unmarshal(body, &response)
		if err != nil {
			break
		}

		if len(response.TorrentResults) == 0 {
			break
		}

		for _, thistorrent := range response.TorrentResults {
			season, episode := GuessSeasonEpisodeNumberFromString(thistorrent.Title)

			temp := ShowTorrent{
				Hash:     GetInfoHashFromMagnetLink(thistorrent.MagnetUri),
				Quality:  GuessQualityFromString(thistorrent.Title),
				Season:   season,
				Episode:  episode,
				Size:     strconv.FormatInt(thistorrent.Size, 10),
				Provider: thistorrent.Tracker,
				Lang:     GuessLanguageFromString(thistorrent.Title),
				Title:    thistorrent.Title,
				Seeds:    strconv.FormatInt(thistorrent.Seeders, 10),
				Peers:    strconv.FormatInt(thistorrent.Peers, 10),
				Magnet:   thistorrent.MagnetUri,
				Torrent:  thistorrent.Link,
			}
			outputShowData = append(outputShowData, temp)
		}
	}

	ch <- outputShowData
	return
}

func SetJackettAddressAndKey(address string, key string) {
	jackettAddress = address
	jackettKey = key
}
