package tvmaze

import (
	"crypto/tls"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

type tvmazeResponse struct {
	Id int64 `json:"id"`
}

func GetTvmazeIdByImdbOrTvdbId(params map[string][]string) int64 {
	// Decode params data
	requesturl := ""

	if params["imdb"] != nil && params["imdb"][0] != "" {
		requesturl = "http://api.tvmaze.com/lookup/shows?imdb=" + params["imdb"][0]
	} else if params["tvdb"] != nil && params["tvdb"][0] != "" {
		requesturl = "http://api.tvmaze.com/lookup/shows?thetvdb=" + params["tvdb"][0]
	}

	req, err := http.NewRequest("GET", requesturl, nil)
	if err != nil {
		return 0
	}

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{Transport: tr, Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return 0
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0
	}

	var message tvmazeResponse
	err = json.Unmarshal(body, &message)
	if err != nil {
		return 0
	}

	return message.Id
}

func GetEpisodesByTvmazeId(id int64) string {
	if id == 0 {
		return "[]"
	}

	requesturl := "http://api.tvmaze.com/shows/" + strconv.FormatInt(id, 10) + "/episodes"

	req, err := http.NewRequest("GET", requesturl, nil)
	if err != nil {
		return "[]"
	}

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{Transport: tr, Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return "[]"
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "[]"
	}

	return string(body)
}
