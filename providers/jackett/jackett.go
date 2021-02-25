package jackett

import (
	"crypto/tls"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"time"

	out "github.com/silentmurdock/wrserver/providers/output"
)

var jackettAddress = ""
var jackettKey = ""

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

func SetJackettAddressAndKey(address string, key string) {
	jackettAddress = address
	jackettKey = key
}

func GetMovieMagnetByImdb(imdb string, ch chan<- []out.OutputMovieStruct) {
	req, err := http.NewRequest("GET", (jackettAddress + "/api/v2.0/indexers/all/results?apikey=" + jackettKey + "&category=2030,2040&query=" + imdb), nil)
	if err != nil {
		ch <- []out.OutputMovieStruct{}
		return
	}

	//req.Header.Set("User-Agent", UserAgent)
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{Transport: tr, Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		ch <- []out.OutputMovieStruct{}
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		ch <- []out.OutputMovieStruct{}
		return
	}

	response := apiResponse{}
	err = json.Unmarshal(body, &response)
	if err != nil {
		ch <- []out.OutputMovieStruct{}
		return
	}

	if len(response.TorrentResults) == 0 {
		ch <- []out.OutputMovieStruct{}
		return
	}

	outputMovieData := []out.OutputMovieStruct{}

	for _, thistorrent := range response.TorrentResults {
		temp := out.OutputMovieStruct{
			Hash:     out.GetInfoHash(thistorrent.MagnetUri),
			Quality:  out.GuessQualityFromString(thistorrent.Title),
			Size:     strconv.FormatInt(thistorrent.Size, 10),
			Provider: thistorrent.Tracker,
			Lang:     out.GuessLanguageFromString(thistorrent.Title),
			Title:    thistorrent.Title,
			Seeds:    strconv.FormatInt(thistorrent.Seeders, 10),
			Peers:    strconv.FormatInt(thistorrent.Peers, 10),
			Magnet:   thistorrent.MagnetUri,
			Torrent:  thistorrent.Link,
		}
		outputMovieData = append(outputMovieData, temp)
	}

	ch <- outputMovieData
	return
}

func GetMovieMagnetByQuery(params map[string][]string, ch chan<- []out.OutputMovieStruct) {
	// Decode params data
	query := ""

	if params["title"] != nil && params["title"][0] != "" {
		query += params["title"][0]
	} else {
		ch <- []out.OutputMovieStruct{}
		return
	}

	if params["releaseyear"] != nil {
		query += " " + params["releaseyear"][0]
	}

	query = url.QueryEscape(query)

	req, err := http.NewRequest("GET", (jackettAddress + "/api/v2.0/indexers/all/results?apikey=" + jackettKey + "&category=2030,2040&query=" + query), nil)
	if err != nil {
		ch <- []out.OutputMovieStruct{}
		return
	}

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{Transport: tr, Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		ch <- []out.OutputMovieStruct{}
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		ch <- []out.OutputMovieStruct{}
		return
	}

	response := apiResponse{}
	err = json.Unmarshal(body, &response)
	if err != nil {
		ch <- []out.OutputMovieStruct{}
		return
	}

	if len(response.TorrentResults) == 0 {
		ch <- []out.OutputMovieStruct{}
		return
	}

	outputMovieData := []out.OutputMovieStruct{}

	for _, thistorrent := range response.TorrentResults {
		temp := out.OutputMovieStruct{
			Hash:     out.GetInfoHash(thistorrent.MagnetUri),
			Quality:  out.GuessQualityFromString(thistorrent.Title),
			Size:     strconv.FormatInt(thistorrent.Size, 10),
			Provider: thistorrent.Tracker,
			Lang:     out.GuessLanguageFromString(thistorrent.Title),
			Title:    thistorrent.Title,
			Seeds:    strconv.FormatInt(thistorrent.Seeders, 10),
			Peers:    strconv.FormatInt(thistorrent.Peers, 10),
			Magnet:   thistorrent.MagnetUri,
			Torrent:  thistorrent.Link,
		}
		outputMovieData = append(outputMovieData, temp)
	}

	ch <- outputMovieData
	return
}

func GetShowMagnetByImdb(imdb string, season string, episode string, ch chan<- []out.OutputShowStruct) {
	req, err := http.NewRequest("GET", (jackettAddress + "/api/v2.0/indexers/all/results?apikey=" + jackettKey + "&category=5030,5040&query=" + imdb), nil)
	if err != nil {
		ch <- []out.OutputShowStruct{}
		return
	}

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{Transport: tr, Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		ch <- []out.OutputShowStruct{}
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		ch <- []out.OutputShowStruct{}
		return
	}

	response := apiResponse{}
	err = json.Unmarshal(body, &response)
	if err != nil {
		ch <- []out.OutputShowStruct{}
		return
	}

	if len(response.TorrentResults) == 0 {
		ch <- []out.OutputShowStruct{}
		return
	}

	outputShowData := []out.OutputShowStruct{}

	for _, thistorrent := range response.TorrentResults {
		titleSeason, titleEpisode := out.GuessSeasonEpisodeNumberFromString(thistorrent.Title)

		if titleSeason == season && titleEpisode == episode {
			temp := out.OutputShowStruct{
				Hash:     out.GetInfoHash(thistorrent.MagnetUri),
				Quality:  out.GuessQualityFromString(thistorrent.Title),
				Season:   season,
				Episode:  episode,
				Size:     strconv.FormatInt(thistorrent.Size, 10),
				Provider: thistorrent.Tracker,
				Lang:     out.GuessLanguageFromString(thistorrent.Title),
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

func GetShowMagnetByQuery(params map[string][]string, season string, episode string, ch chan<- []out.OutputShowStruct) {
	// Decode params data
	query := ""

	if params["title"] != nil && params["title"][0] != "" {
		query += params["title"][0] + " "
	} else {
		ch <- []out.OutputShowStruct{}
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

	req, err := http.NewRequest("GET", (jackettAddress + "/api/v2.0/indexers/all/results?apikey=" + jackettKey + "&category=5030,5040&query=" + query), nil)
	if err != nil {
		ch <- []out.OutputShowStruct{}
		return
	}

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{Transport: tr, Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		ch <- []out.OutputShowStruct{}
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		ch <- []out.OutputShowStruct{}
		return
	}

	response := apiResponse{}
	err = json.Unmarshal(body, &response)
	if err != nil {
		ch <- []out.OutputShowStruct{}
		return
	}

	if len(response.TorrentResults) == 0 {
		ch <- []out.OutputShowStruct{}
		return
	}

	outputShowData := []out.OutputShowStruct{}

	for _, thistorrent := range response.TorrentResults {
		season, episode := out.GuessSeasonEpisodeNumberFromString(thistorrent.Title)

		temp := out.OutputShowStruct{
			Hash:     out.GetInfoHash(thistorrent.MagnetUri),
			Quality:  out.GuessQualityFromString(thistorrent.Title),
			Season:   season,
			Episode:  episode,
			Size:     strconv.FormatInt(thistorrent.Size, 10),
			Provider: thistorrent.Tracker,
			Lang:     out.GuessLanguageFromString(thistorrent.Title),
			Title:    thistorrent.Title,
			Seeds:    strconv.FormatInt(thistorrent.Seeders, 10),
			Peers:    strconv.FormatInt(thistorrent.Peers, 10),
			Magnet:   thistorrent.MagnetUri,
			Torrent:  thistorrent.Link,
		}
		outputShowData = append(outputShowData, temp)
	}

	ch <- outputShowData
	return
}
