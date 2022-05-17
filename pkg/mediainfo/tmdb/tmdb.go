package tmdb

import (
	"crypto/tls"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/nyakaspeter/raven-torrent/internal/settings"
	"github.com/nyakaspeter/raven-torrent/pkg/mediainfo/types"
)

func DiscoverMovies(params types.MovieDiscoverParams, language string, page int) (types.MovieResults, error) {
	requesturl := "https://api.themoviedb.org/3/discover/movie?api_key=" + *settings.TMDBKey +
		"&with_original_language=en" +
		"&region=US&with_release_type=5" +
		"&language=" + language +
		"&page=" + strconv.Itoa(page)

	if params.SortBy != "" {
		requesturl += "&sort_by=" + params.SortBy
	} else {
		requesturl += "&sort_by=popularity.desc"
	}

	if params.MaxReleaseDate != "" {
		requesturl += "&release_date.lte=" + params.MaxReleaseDate
	}

	if params.MinReleaseDate != "" {
		requesturl += "&release_date.gte=" + params.MinReleaseDate
	}

	if len(params.GenreIds) > 0 {
		genreIdsJson, _ := json.Marshal(params.GenreIds)
		genresIdsString := strings.Trim(string(genreIdsJson), "[]")
		requesturl += "&with_genres=" + genresIdsString
	}

	req, err := http.NewRequest("GET", requesturl, nil)
	if err != nil {
		return types.MovieResults{}, err
	}

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{Transport: tr, Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return types.MovieResults{}, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return types.MovieResults{}, err
	}

	var results types.MovieResults
	err = json.Unmarshal(body, &results)
	if err != nil {
		return types.MovieResults{}, err
	}

	return results, nil
}

func DiscoverShows(params types.ShowDiscoverParams, language string, page int) (types.ShowResults, error) {
	requesturl := "https://api.themoviedb.org/3/discover/tv?api_key=" + *settings.TMDBKey +
		"&with_original_language=en" +
		"&language=" + language +
		"&page=" + strconv.Itoa(page)

	if params.SortBy != "" {
		requesturl += "&sort_by=" + params.SortBy
	} else {
		requesturl += "&sort_by=popularity.desc"
	}

	if params.MaxAirDate != "" {
		requesturl += "&air_date.lte=" + params.MaxAirDate
	}

	if params.MinAirDate != "" {
		requesturl += "&air_date.gte=" + params.MinAirDate
	}

	if len(params.GenreIds) > 0 {
		genreIdsJson, _ := json.Marshal(params.GenreIds)
		genresIdsString := strings.Trim(string(genreIdsJson), "[]")
		requesturl += "&with_genres=" + genresIdsString
	}

	req, err := http.NewRequest("GET", requesturl, nil)
	if err != nil {
		return types.ShowResults{}, err
	}

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{Transport: tr, Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return types.ShowResults{}, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return types.ShowResults{}, err
	}

	var results types.ShowResults
	err = json.Unmarshal(body, &results)
	if err != nil {
		return types.ShowResults{}, err
	}

	return results, nil
}

func SearchMovies(title string, language string, page int) (types.MovieResults, error) {
	req, err := http.NewRequest("GET", "https://api.themoviedb.org/3/search/movie?api_key="+*settings.TMDBKey+"&language="+language+"&page="+strconv.Itoa(page)+"&query="+url.QueryEscape(title), nil)
	if err != nil {
		return types.MovieResults{}, err
	}

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{Transport: tr, Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return types.MovieResults{}, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return types.MovieResults{}, err
	}

	var results types.MovieResults
	err = json.Unmarshal(body, &results)
	if err != nil {
		return types.MovieResults{}, err
	}

	return results, nil
}

func SearchShows(title string, language string, page int) (types.ShowResults, error) {
	req, err := http.NewRequest("GET", "https://api.themoviedb.org/3/search/tv?api_key="+*settings.TMDBKey+"&language="+language+"&page="+strconv.Itoa(page)+"&query="+url.QueryEscape(title), nil)
	if err != nil {
		return types.ShowResults{}, err
	}

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{Transport: tr, Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return types.ShowResults{}, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return types.ShowResults{}, err
	}

	var results types.ShowResults
	err = json.Unmarshal(body, &results)
	if err != nil {
		return types.ShowResults{}, err
	}

	return results, nil
}

func GetMovieInfo(tmdbId int, language string) (types.MovieInfo, error) {
	requesturl := "https://api.themoviedb.org/3/movie/" + strconv.Itoa(tmdbId) + "?api_key=" + *settings.TMDBKey + "&language=" + language

	req, err := http.NewRequest("GET", requesturl, nil)
	if err != nil {
		return types.MovieInfo{}, err
	}

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{Transport: tr, Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return types.MovieInfo{}, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return types.MovieInfo{}, err
	}

	var result types.MovieInfo
	err = json.Unmarshal(body, &result)
	if err != nil {
		return types.MovieInfo{}, err
	}

	return result, nil
}

func GetShowInfo(tmdbId int, language string) (types.ShowInfo, error) {
	requesturl := "https://api.themoviedb.org/3/tv/" + strconv.Itoa(tmdbId) + "?api_key=" + *settings.TMDBKey + "&append_to_response=external_ids&language=" + language

	req, err := http.NewRequest("GET", requesturl, nil)
	if err != nil {
		return types.ShowInfo{}, err
	}

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{Transport: tr, Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return types.ShowInfo{}, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return types.ShowInfo{}, err
	}

	var result types.ShowInfo
	err = json.Unmarshal(body, &result)
	if err != nil {
		return types.ShowInfo{}, err
	}

	return result, nil
}

func GetShowSeason(tmdbId int, seasonNumber int, language string) (types.SeasonInfo, error) {
	requesturl := "https://api.themoviedb.org/3/tv/" + strconv.Itoa(tmdbId) + "/season/" + strconv.Itoa(seasonNumber) + "?api_key=" + *settings.TMDBKey + "&language=" + language

	req, err := http.NewRequest("GET", requesturl, nil)
	if err != nil {
		return types.SeasonInfo{}, err
	}

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{Transport: tr, Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return types.SeasonInfo{}, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return types.SeasonInfo{}, err
	}

	var result types.SeasonInfo
	err = json.Unmarshal(body, &result)
	if err != nil {
		return types.SeasonInfo{}, err
	}

	return result, nil
}
