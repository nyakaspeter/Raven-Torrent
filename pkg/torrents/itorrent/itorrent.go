package itorrent

import (
	"crypto/tls"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/nyakaspeter/raven-torrent/pkg/torrents/types"
	"github.com/nyakaspeter/raven-torrent/pkg/torrents/utils"
)

func GetMovieTorrentsByImdbId(imdb string, ch chan<- []types.MovieTorrent) {
	// Disable security
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	http.DefaultTransport.(*http.Transport).ResponseHeaderTimeout = 10 * time.Second

	doc, err := goquery.NewDocument("https://itorrent.ws/torrentek/category/3/title/" + imdb + "/view_mode/photos/")
	if err != nil {
		ch <- []types.MovieTorrent{}
		return
	}

	outputMovieData := []types.MovieTorrent{}

	innerCh := make(chan types.MovieTorrent)

	counter := 0
	doc.Find("#ajaxtable .text-container").Each(func(_ int, item *goquery.Selection) {
		linkTag := item.Find("a")
		link, _ := linkTag.Attr("href")
		go scrapeMovieData(link, innerCh)
		counter++
	})

	for counter > 0 {
		temp := <-innerCh
		if temp.Hash != "" {
			outputMovieData = append(outputMovieData, temp)
		}
		counter--
	}

	ch <- outputMovieData
}

func GetShowTorrentsByImdbId(imdb string, season string, episode string, ch chan<- []types.ShowTorrent) {
	// Disable security
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	http.DefaultTransport.(*http.Transport).ResponseHeaderTimeout = 10 * time.Second

	link := "https://itorrent.ws/torrentek/category/4/title/" + imdb + "/"

	if season != "0" {
		link += "series_season/" + season + "/"
	}

	if episode != "0" {
		link += "series_episode/" + episode + "/"
	}

	doc, err := goquery.NewDocument(link)
	if err != nil {
		ch <- []types.ShowTorrent{}
		return
	}

	outputMovieData := []types.ShowTorrent{}

	innerCh := make(chan types.ShowTorrent)

	counter := 0
	doc.Find("td.ellipse").Each(func(_ int, item *goquery.Selection) {
		linkTag := item.Find("a")
		link, _ := linkTag.Attr("href")
		go scrapeShowData(link, season, episode, innerCh)
		counter++
	})

	for counter > 0 {
		temp := <-innerCh
		if temp.Hash != "" {
			outputMovieData = append(outputMovieData, temp)
		}
		counter--
	}

	ch <- outputMovieData
}

func scrapeMovieData(movieUrl string, innerCh chan<- types.MovieTorrent) {
	doc, err := goquery.NewDocument("https://itorrent.ws" + movieUrl)
	if err != nil {
		innerCh <- types.MovieTorrent{}
		return
	}

	// Find title for raw magnet selection
	title := doc.Find("h1#torrent_title").Text()
	title = strings.TrimSpace(title)

	// Try to decode quality information from movieUrl
	quality := utils.GuessQualityFromString(movieUrl)

	// Find Magnet link and decode infohash
	magnet := ""
	infoHash := ""

	doc.Find(".btn.btn-success.seed-warning").Each(func(_ int, item *goquery.Selection) {
		magnet, _ = item.Attr("href")
		infoHash = utils.GetInfoHashFromMagnetLink(magnet)
	})

	// Find Torrent link
	torrent := ""
	doc.Find(".btn.btn-primary.seed-warning").Each(func(_ int, item *goquery.Selection) {
		torrent, _ = item.Attr("href")
		torrent = "https://itorrent.ws" + torrent
	})

	size := ""
	language := ""
	seeds := ""
	leech := ""
	seedInt := int64(0)
	doc.Find("#torrent_page .left1").Each(func(_ int, item *goquery.Selection) {

		dataType := item.Find(".type").Text()
		switch dataType {
		case "Méret":
			size = utils.DecodeSize(item.Next().Text())
		case "Peer":
			value := item.Next().Text()
			re := regexp.MustCompile("[0-9]+")
			stringsize := re.FindAllString(value, -1)
			seedInt, _ = strconv.ParseInt(stringsize[0], 10, 64)
			seeds = stringsize[0]
			leech = stringsize[1]
		case "Nyelv":
			language = utils.DecodeLanguage(item.Next().Text(), "hu")
		}
	})

	if seedInt == 0 {
		innerCh <- types.MovieTorrent{}
		return
	}

	/*intSize, _ := strconv.ParseInt(size, 10, 64)
	  if intSize > (5 * 1024 * 1024 * 1024) {
	      innerCh <- OutputMovieStruct{}
	  }*/

	innerCh <- types.MovieTorrent{
		Hash:     infoHash,
		Quality:  quality,
		Size:     size,
		Provider: "ITORRENT",
		Lang:     language,
		Title:    title,
		Seeds:    seeds,
		Peers:    leech,
		Magnet:   magnet,
		Torrent:  torrent,
	}
}

func scrapeShowData(movieUrl string, season string, episode string, innerCh chan<- types.ShowTorrent) {
	doc, err := goquery.NewDocument("https://itorrent.ws" + movieUrl)
	if err != nil {
		innerCh <- types.ShowTorrent{}
		return
	}

	// Find title for raw magnet selection
	title := doc.Find("h1#torrent_title").Text()
	title = strings.TrimSpace(title)

	// Try to find episode number from title
	season, episode = utils.GuessSeasonEpisodeNumberFromString(title)

	// Try to decode quality information from movieUrl
	quality := utils.GuessQualityFromString(movieUrl)

	// Find Magnet link and decode infohash
	magnet := ""
	infoHash := ""

	doc.Find(".btn.btn-success.seed-warning").Each(func(_ int, item *goquery.Selection) {
		magnet, _ = item.Attr("href")
		infoHash = utils.GetInfoHashFromMagnetLink(magnet)
	})

	// Find Torrent link
	torrent := ""
	doc.Find(".btn.btn-primary.seed-warning").Each(func(_ int, item *goquery.Selection) {
		torrent, _ = item.Attr("href")
		torrent = "http://itorrent.ws" + torrent
	})

	size := ""
	language := ""
	seeds := ""
	leech := ""
	seedInt := int64(0)
	doc.Find("#torrent_page .left1").Each(func(_ int, item *goquery.Selection) {

		dataType := item.Find(".type").Text()
		switch dataType {
		case "Méret":
			size = utils.DecodeSize(item.Next().Text())
		case "Peer":
			value := item.Next().Text()
			re := regexp.MustCompile("[0-9]+")
			stringsize := re.FindAllString(value, -1)
			seedInt, _ = strconv.ParseInt(stringsize[0], 10, 64)
			seeds = stringsize[0]
			leech = stringsize[1]
		case "Nyelv":
			language = utils.DecodeLanguage(item.Next().Text(), "hu")
		}
	})

	if seedInt == 0 {
		innerCh <- types.ShowTorrent{}
		return
	}

	innerCh <- types.ShowTorrent{
		Hash:     infoHash,
		Quality:  quality,
		Size:     size,
		Provider: "ITORRENT",
		Lang:     language,
		Title:    title,
		Seeds:    seeds,
		Peers:    leech,
		Season:   season,
		Episode:  episode,
		Magnet:   magnet,
		Torrent:  torrent,
	}
}
