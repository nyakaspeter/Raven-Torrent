package x1337x

import (
	"crypto/tls"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/nyakaspeter/raven-torrent/pkg/torrents/types"
	"github.com/nyakaspeter/raven-torrent/pkg/torrents/utils"
)

func GetMovieTorrentsByText(searchText string, ch chan<- []types.MovieTorrent) {
	query := url.QueryEscape(searchText)

	// Disable security
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	http.DefaultTransport.(*http.Transport).ResponseHeaderTimeout = 10 * time.Second

	doc, err := goquery.NewDocument("https://www.1337x.to/category-search/" + query + "/Movies/1/")
	if err != nil {
		ch <- []types.MovieTorrent{}
		return
	}

	outputMovieData := []types.MovieTorrent{}

	innerCh := make(chan types.MovieTorrent)

	counter := 0
	doc.Find("tbody tr").Each(func(_ int, item *goquery.Selection) {
		seedsClass := item.Find("td.seeds")
		seeds := seedsClass.Text()

		if seeds != "0" {
			nameClass := item.Find("td.name")
			linkTag := nameClass.Find("a").Next()
			link, _ := linkTag.Attr("href")

			go scrapeMovieData(link, innerCh)
			counter++
		}
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

func GetShowTorrentsByText(searchText string, season string, episode string, ch chan<- []types.ShowTorrent) {
	query := searchText + " "

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

	// Disable security
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	http.DefaultTransport.(*http.Transport).ResponseHeaderTimeout = 10 * time.Second

	doc, err := goquery.NewDocument("https://www.1337x.to/category-search/" + query + "/TV/1/")
	if err != nil {
		ch <- []types.ShowTorrent{}
		return
	}

	outputShowData := []types.ShowTorrent{}

	innerCh := make(chan types.ShowTorrent)

	counter := 0
	doc.Find("tbody tr").Each(func(_ int, item *goquery.Selection) {
		seedsClass := item.Find("td.seeds")
		seeds := seedsClass.Text()

		if seeds != "0" {
			nameClass := item.Find("td.name")
			linkTag := nameClass.Find("a").Next()
			link, _ := linkTag.Attr("href")

			go scrapeShowData(link, season, episode, innerCh)
			counter++
		}
	})

	for counter > 0 {
		temp := <-innerCh
		if temp.Hash != "" {
			outputShowData = append(outputShowData, temp)
		}
		counter--
	}

	ch <- outputShowData
}

func scrapeMovieData(movieUrl string, innerCh chan<- types.MovieTorrent) {
	doc, err := goquery.NewDocument("https://www.1337x.to" + movieUrl)
	if err != nil {
		innerCh <- types.MovieTorrent{}
	}

	// Find title for raw magnet selection
	title := doc.Find("title").Text()
	title = strings.TrimPrefix(title, "Download")
	title = strings.TrimSuffix(title, "Torrent | 1337x")
	title = utils.CleanString(title)

	// Trimmed title
	//title := doc.Find(".box-info-heading h1").Text()
	//title = strings.TrimSpace(title)

	// Try to decode quality information from movieUrl
	quality := utils.GuessQualityFromString(title)

	// Find Magnet link and decode infohash
	link := ""
	infoHash := ""
	doc.Find(".torrent-detail-page ul li a").Each(func(_ int, item *goquery.Selection) {
		if item.Text() == "Magnet Download" {
			link, _ = item.Attr("href")
			infoHash = utils.GetInfoHashFromMagnetLink(link)
		}
	})

	size := ""
	language := ""
	seeders := ""
	leechers := ""
	doc.Find(".torrent-detail-page ul.list li").Each(func(_ int, item *goquery.Selection) {
		textNode := item.ChildrenFiltered("strong").Text()
		if textNode == "Total size" {
			size = utils.DecodeSize(item.ChildrenFiltered("span").Text())
		} else if textNode == "Language" {
			language = utils.DecodeLanguage(item.ChildrenFiltered("span").Text(), "en")
		} else if textNode == "Seeders" {
			seeders = item.ChildrenFiltered("span").Text()
		} else if textNode == "Leechers" {
			leechers = item.ChildrenFiltered("span").Text()
		}
	})

	innerCh <- types.MovieTorrent{
		Hash:     infoHash,
		Quality:  quality,
		Size:     size,
		Provider: "1337X",
		Lang:     language,
		Title:    title,
		Seeds:    seeders,
		Peers:    leechers,
		Magnet:   link,
	}
}

func scrapeShowData(movieUrl string, season string, episode string, innerCh chan<- types.ShowTorrent) {
	doc, err := goquery.NewDocument("https://www.1337x.to" + movieUrl)
	if err != nil {
		innerCh <- types.ShowTorrent{}
	}

	// Find title for raw magnet selection
	title := doc.Find("title").Text()
	title = strings.TrimPrefix(title, "Download")
	title = strings.TrimSuffix(title, "Torrent | 1337x")
	title = utils.CleanString(title)

	// Trimmed title
	//title := doc.Find(".box-info-heading h1").Text()
	//title = strings.TrimSpace(title)

	// Try to decode quality information from movieUrl
	quality := utils.GuessQualityFromString(title)

	// Find Magnet link and decode infohash
	link := ""
	infoHash := ""
	doc.Find(".torrent-detail-page ul li a").Each(func(_ int, item *goquery.Selection) {
		if item.Text() == "Magnet Download" {
			link, _ = item.Attr("href")
			infoHash = utils.GetInfoHashFromMagnetLink(link)
		}
	})

	size := ""
	language := ""
	seeders := ""
	leechers := ""
	doc.Find(".torrent-detail-page ul.list li").Each(func(_ int, item *goquery.Selection) {
		textNode := item.ChildrenFiltered("strong").Text()
		if textNode == "Total size" {
			size = utils.DecodeSize(item.ChildrenFiltered("span").Text())
		} else if textNode == "Language" {
			language = utils.DecodeLanguage(item.ChildrenFiltered("span").Text(), "en")
		} else if textNode == "Seeders" {
			seeders = item.ChildrenFiltered("span").Text()
		} else if textNode == "Leechers" {
			leechers = item.ChildrenFiltered("span").Text()
		}
	})

	seasonNum, epNum := utils.GuessSeasonEpisodeNumberFromString(title)

	innerCh <- types.ShowTorrent{
		Hash:     infoHash,
		Quality:  quality,
		Size:     size,
		Provider: "1337X",
		Lang:     language,
		Title:    title,
		Seeds:    seeders,
		Peers:    leechers,
		Season:   seasonNum,
		Episode:  epNum,
		Magnet:   link,
	}
}
