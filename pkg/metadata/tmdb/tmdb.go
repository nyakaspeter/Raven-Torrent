package tmdb

import (
	"crypto/tls"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

type tmdbResponse struct {
	TotalResults int `json:"total_results"`
	Id           int `json:"id"`
}

var TmdbKey string = ""

func Discover(qtype string, genretype string, sort string, date string, lang string, cpage string) string {
	requesturl := ""

	if qtype == "movie" {
		if genretype == "all" {
			requesturl = "https://api.themoviedb.org/3/discover/" + qtype + "?api_key=" + TmdbKey + "&sort_by=" + sort + "&release_date.lte=" + date + "&with_original_language=en&region=US&with_release_type=5&language=" + lang + "&page=" + cpage
		} else {
			requesturl = "https://api.themoviedb.org/3/discover/" + qtype + "?api_key=" + TmdbKey + "&sort_by=" + sort + "&release_date.lte=" + date + "&with_original_language=en&region=US&with_release_type=5&with_genres=" + genretype + "&language=" + lang + "&page=" + cpage
		}
	} else {
		if genretype == "all" {
			requesturl = "https://api.themoviedb.org/3/discover/" + qtype + "?api_key=" + TmdbKey + "&sort_by=" + sort + "&air_date.lte=" + date + "&with_original_language=en&language=" + lang + "&page=" + cpage
		} else {
			requesturl = "https://api.themoviedb.org/3/discover/" + qtype + "?api_key=" + TmdbKey + "&sort_by=" + sort + "&air_date.lte=" + date + "&with_original_language=en&with_genres=" + genretype + "&language=" + lang + "&page=" + cpage
		}
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

	var message tmdbResponse
	err = json.Unmarshal(body, &message)
	if err != nil || message.TotalResults == 0 {
		return ""
	}

	return string(body)
}

func Search(qtype string, lang string, cpage string, typedtext string) string {
	req, err := http.NewRequest("GET", "https://api.themoviedb.org/3/search/"+qtype+"?api_key="+TmdbKey+"&language="+lang+"&page="+cpage+"&query="+typedtext, nil)
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

	var message tmdbResponse
	err = json.Unmarshal(body, &message)
	if err != nil || message.TotalResults == 0 {
		return ""
	}

	return string(body)
}

func GetInfo(qtype string, tmdbid string, lang string) string {
	requesturl := "https://api.themoviedb.org/3/" + qtype + "/" + tmdbid + "?api_key=" + TmdbKey + "&language=" + lang
	if qtype == "tv" {
		requesturl = "https://api.themoviedb.org/3/" + qtype + "/" + tmdbid + "?api_key=" + TmdbKey + "&append_to_response=external_ids&language=" + lang
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

	var message tmdbResponse
	err = json.Unmarshal(body, &message)
	if err != nil || message.Id == 0 {
		return ""
	}

	return string(body)
}
