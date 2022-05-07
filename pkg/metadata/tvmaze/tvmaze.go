package tvmaze

import (
	"crypto/tls"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

type tvmazeIdResponse struct {
	Id int `json:"id"`
}

type tvmazeErrorResponse struct {
	Code   int `json:"code"`
	Status int `json:"status"`
}

func GetEpisodes(tvdbid string, imdbid string) string {
	tvmazeid := ""

	if tvdbid != "" {
		tvmazeid = getTvMazeId("tvdb", tvdbid)
	}

	if tvmazeid == "" && imdbid != "" {
		tvmazeid = getTvMazeId("imdb", imdbid)
	}

	if tvmazeid == "" {
		return ""
	}

	requesturl := "https://api.tvmaze.com/shows/" + tvmazeid + "/episodes"

	req, err := http.NewRequest("GET", requesturl, nil)
	if err != nil {
		return ""
	}

	//req.Header.Set("User-Agent", UserAgent)
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{Transport: tr, Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return ""
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return ""
	}

	var message tvmazeErrorResponse
	err = json.Unmarshal(body, &message)
	if err == nil || string(body) == "null" {
		return ""
	}

	return string(body)
}

func getTvMazeId(qtype string, id string) string {
	requesturl := ""

	if qtype == "tvdb" {
		requesturl = "https://api.tvmaze.com/lookup/shows?thetvdb=" + id
	} else {
		requesturl = "https://api.tvmaze.com/lookup/shows?imdb=" + id
	}

	req, err := http.NewRequest("GET", requesturl, nil)
	if err != nil {
		return ""
	}

	//req.Header.Set("User-Agent", UserAgent)
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{Transport: tr, Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return ""
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return ""
	}

	var message tvmazeIdResponse
	err = json.Unmarshal(body, &message)
	if err != nil || string(body) == "null" || message.Id == 0 {
		return ""
	}

	return strconv.Itoa(message.Id)
}
