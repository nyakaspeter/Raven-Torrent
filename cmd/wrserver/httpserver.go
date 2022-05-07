package main

import (
	"archive/zip"
	"bytes"
	"crypto/tls"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/url"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"github.com/anacrolix/torrent/metainfo"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"github.com/koron/go-ssdp"
	"github.com/martinlindhe/subtitles"
	"github.com/oz/osdb"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/charmap"

	"github.com/silentmurdock/wrserver/pkg/dlnacast"
	"github.com/silentmurdock/wrserver/pkg/metadata/tmdb"
	"github.com/silentmurdock/wrserver/pkg/metadata/tvmaze"
	"github.com/silentmurdock/wrserver/pkg/torrents"
)

const apiPrefix = "/api/"

var serverHost string
var serverPort int
var dlnaServerPort int
var dlnaServer *dlnacast.HTTPserver = nil
var dlnaTvPayload *dlnacast.TVPayload = nil
var openSubtitlesUserAgent string
var webSocket *websocket.Conn
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func startHTTPServer(host string, port int, cors bool) *http.Server {
	newsrv := &http.Server{
		Addr:         fmt.Sprintf("%s:%d", host, port),
		ReadTimeout:  38 * time.Second,
		WriteTimeout: 38 * time.Second,
		Handler:      handleAPI(cors),
	}

	localIP := host
	if localIP == "" {
		localIP = getLocalIP()
	}

	serverHost = localIP
	serverPort = port

	address := fmt.Sprintf("http://%s:%d", localIP, port)

	// Must appear
	fmt.Printf("White Raven Server Version %s Started On Address: %s\n", version, address)

	go func() {
		if err := newsrv.ListenAndServe(); err != nil {
			// cannot panic, because this probably is an intentional close
			if err == http.ErrServerClosed {
				fmt.Printf("HTTP Server Closed\n")
			} else {
				fmt.Printf("HTTP Server Error: %s\n", err)
			}
			time.Sleep(1 * time.Nanosecond)
			procQuit <- true
		}
	}()

	// returning reference so caller can call Shutdown()
	return newsrv
}

func handleAPI(cors bool) http.Handler {
	router := mux.NewRouter()
	router.SkipClean(true)

	router.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, resourceNotFound(), http.StatusNotFound)
	})

	router.HandleFunc(apiPrefix+"about", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, serverInfo())
	})

	router.HandleFunc(apiPrefix+"stop", func(w http.ResponseWriter, r *http.Request) {
		_, err := io.WriteString(w, serverStopped())
		if err == nil {
			go func() {
				time.Sleep(1 * time.Nanosecond)
				procQuit <- true
			}()
		} else {
			go func() {
				time.Sleep(1 * time.Nanosecond)
				procQuit <- true
			}()
		}
	})

	router.HandleFunc(apiPrefix+"restart/downrate/{downrate}/uprate/{uprate}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		for _, thistorrent := range activeTorrents {
			log.Println("Delete torrent:", thistorrent.torrent.InfoHash().String())
			stopAllFileDownload(thistorrent.torrent.Files())
			thistorrent.torrent.Drop()
			delete(activeTorrents, thistorrent.torrent.InfoHash().String())
		}

		_, err := io.WriteString(w, torrentClientRestarting())
		if err == nil {
			go func() {
				time.Sleep(1 * time.Nanosecond)
				dr, _ := strconv.ParseInt(vars["downrate"], 10, 64)
				ur, _ := strconv.ParseInt(vars["uprate"], 10, 64)
				procRestart <- []int64{dr, ur}
			}()
		} else {
			go func() {
				time.Sleep(1 * time.Nanosecond)
				dr, _ := strconv.ParseInt(vars["downrate"], 10, 64)
				ur, _ := strconv.ParseInt(vars["uprate"], 10, 64)
				procRestart <- []int64{dr, ur}
			}()
		}
	})

	router.HandleFunc(apiPrefix+"add/{base64uri}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		base64uri := vars["base64uri"]

		uri, err := base64.StdEncoding.DecodeString(base64uri)

		if err != nil {
			http.Error(w, failedToAddTorrent(), http.StatusNotFound)
			return
		}

		for tryCount := 0; tryCount < 4; tryCount++ {
			if tryCount > 0 {
				time.Sleep(10 * time.Second)
			}

			t := addTorrent(string(uri))

			if t != nil {
				log.Println("Added torrent:", t.InfoHash().String())
				io.WriteString(w, torrentFilesList(r.Host, t.InfoHash().String(), t.Files()))
				return
			} else if len(activeTorrents) == 0 {
				http.Error(w, failedToAddTorrent(), http.StatusNotFound)
				return
			}
		}

		if len(activeTorrents) > 0 {
			http.Error(w, onlyOneTorrentAllowed(), http.StatusNotFound)
		}
	})

	router.HandleFunc(apiPrefix+"getmediarenderers", func(w http.ResponseWriter, r *http.Request) {
		devices := []mediaRenderer{}
		list, err := ssdp.Search(ssdp.All, 1, "")

		if err != nil || len(list) == 0 {
			http.Error(w, noMediaRenderersFound(), http.StatusNotFound)
			log.Println("No media renderers found.")
		}

		for _, srv := range list {
			if srv.Type == "urn:schemas-upnp-org:service:AVTransport:1" {
				duplicate := false
				for _, device := range devices {
					if device.Location == srv.Location {
						duplicate = true
					}
				}

				if !duplicate {
					if friendlyName, err := dlnacast.GetDeviceFriendlyName(srv.Location); err != nil {
						devices = append(devices, mediaRenderer{Name: srv.Server, Location: srv.Location})
					} else {
						devices = append(devices, mediaRenderer{Name: friendlyName, Location: srv.Location})
					}
				}
			}
		}

		if len(devices) > 0 {
			io.WriteString(w, mediaRenderersList(devices))
		} else {
			http.Error(w, noMediaRenderersFound(), http.StatusNotFound)
			log.Println("No media renderers found.")
		}
	})

	router.HandleFunc(apiPrefix+"startplayer/{base64path}/{base64args}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		var err error

		path, err := base64.StdEncoding.DecodeString(vars["base64path"])
		args, err := base64.StdEncoding.DecodeString(vars["base64args"])

		splitArgs := strings.Split(string(args), ",")

		if err != nil {
			http.Error(w, failedToOpenMediaPlayer(), http.StatusNotFound)
			return
		}

		cmd := exec.Command(string(path), splitArgs...)
		err = cmd.Run()

		if err != nil {
			http.Error(w, failedToOpenMediaPlayer(), http.StatusNotFound)
			return
		}

		io.WriteString(w, "{\"success\": true}")
	})

	router.HandleFunc(apiPrefix+"cast/{base64location}/{base64query}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		var err error
		var videoURL = ""
		var subtitlesURL = ""
		var videoTitle = "video"

		location, err := base64.StdEncoding.DecodeString(vars["base64location"])
		query, err := base64.StdEncoding.DecodeString(vars["base64query"])

		if err != nil {
			http.Error(w, failedCastingToDevice(), http.StatusNotFound)
			return
		}

		params, err := url.ParseQuery(string(query))

		if err != nil {
			http.Error(w, failedCastingToDevice(), http.StatusNotFound)
			return
		}

		if params["video"] != nil && params["video"][0] != "" {
			videoURL = params["video"][0]
		}

		if params["subtitle"] != nil && params["subtitle"][0] != "" {
			subtitlesURL = params["subtitle"][0]
		}

		if params["title"] != nil && params["title"][0] != "" {
			videoTitle = params["title"][0]
		}

		transportURL, controlURL, err := dlnacast.GetAvTransportUrl(string(location))

		if err != nil || videoURL == "" {
			http.Error(w, failedCastingToDevice(), http.StatusNotFound)
			return
		}

		videoURL = strings.Replace(videoURL, "localhost", serverHost, 1)
		subtitlesURL = strings.Replace(subtitlesURL, "localhost", serverHost, 1)

		whereToListen := fmt.Sprintf("%s:%d", serverHost, dlnaServerPort)
		callbackURL := fmt.Sprintf("http://%s/callback", whereToListen)

		newPayload := &dlnacast.TVPayload{
			TransportURL: transportURL,
			ControlURL:   controlURL,
			CallbackURL:  callbackURL,
			VideoURL:     videoURL,
			SubtitlesURL: subtitlesURL,
			VideoTitle:   videoTitle,
		}

		serverStarted := make(chan struct{})

		if dlnaServer == nil {
			srv := dlnacast.CreateServer(whereToListen)
			dlnaServer = &srv
		} else {
			dlnaTvPayload.SendtoTV("Stop")
			dlnaServer.StopServer()
		}

		dlnaTvPayload = newPayload

		go func() {
			dlnaServer.StartServer(serverStarted, newPayload)
		}()
		// Wait for HTTP server to properly initialize
		<-serverStarted

		err = newPayload.SendtoTV("Play1")

		if err != nil || videoURL == "" {
			http.Error(w, failedCastingToDevice(), http.StatusNotFound)
			return
		}

		io.WriteString(w, "{\"success\": true}")
	})

	router.HandleFunc(apiPrefix+"delete/{hash}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		if t, ok := activeTorrents[vars["hash"]]; ok {

			log.Println("Delete torrent:", vars["hash"])
			stopAllFileDownload(t.torrent.Files())
			t.torrent.Drop()
			delete(activeTorrents, vars["hash"])

			io.WriteString(w, deletedTorrent())
		} else {
			http.Error(w, torrentNotFound(), http.StatusNotFound)
			log.Println("Torrent not found:", vars["hash"])
		}
	})

	router.HandleFunc(apiPrefix+"deleteall", func(w http.ResponseWriter, r *http.Request) {
		if len(activeTorrents) > 0 {
			for _, thistorrent := range activeTorrents {
				log.Println("Delete torrent:", thistorrent.torrent.InfoHash().String())
				stopAllFileDownload(thistorrent.torrent.Files())
				thistorrent.torrent.Drop()
				delete(activeTorrents, thistorrent.torrent.InfoHash().String())
			}
			io.WriteString(w, deletedAllTorrents())
		} else {
			http.Error(w, noActiveTorrentsFound(), http.StatusNotFound)
			log.Println("No active torrents found.")
		}
	})

	router.HandleFunc(apiPrefix+"get/{hash}/{base64path}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		if d, err := base64.StdEncoding.DecodeString(vars["base64path"]); err == nil {
			if t, ok := activeTorrents[vars["hash"]]; ok {

				idx := getFileByPath(string(d), t.torrent.Files())
				if idx != -1 {
					file := t.torrent.Files()[idx]

					path := file.DisplayPath()
					log.Println("Downloading torrent:", vars["hash"])

					increaseFileClients(path, t)

					/*log.Println("Calculate Opensubtitles hash...")
					fileHash := calculateOpensubtitlesHash(file)
					log.Println("Opensubtitles hash calculated:", fileHash)*/

					serveTorrentFile(w, r, file)
					//stop downloading the file when no connections left
					if decreaseFileClients(path, t) <= 0 {
						stopDownloadFile(file)
					}
				} else {
					http.Error(w, invalidBase64Path(), http.StatusNotFound)
					return
				}
			} else {
				http.Error(w, torrentNotFound(), http.StatusNotFound)
				log.Println("Unknown torrent:", vars["hash"])
				return
			}
		} else {
			http.Error(w, invalidBase64Path(), http.StatusNotFound)
			log.Println(err)
			return
		}
	})

	router.HandleFunc(apiPrefix+"stats/{hash}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		if t, ok := activeTorrents[vars["hash"]]; ok {
			log.Println("Check torrent stats:", vars["hash"])
			io.WriteString(w, downloadStats(r.Host, t.torrent))
		} else {
			http.Error(w, torrentNotFound(), http.StatusNotFound)
			log.Println("Unknown torrent:", vars["hash"])
		}
	})

	router.HandleFunc(apiPrefix+"subtitlesbyimdb/{imdb}/lang/{lang}/season/{season}/episode/{episode}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		// Create Opensubtitles client
		c, err := osdb.NewClient()
		if err != nil {
			http.Error(w, failedToConnectToOpenSubtitles(), http.StatusNotFound)
			return
		}

		c.UserAgent = openSubtitlesUserAgent

		// Anonymous Login with UserAgent string will set c.Token when successful
		if err = c.LogIn("", "", ""); err != nil {
			http.Error(w, failedToConnectToOpenSubtitles(), http.StatusNotFound)
			return
		}

		ids := make([]string, 1)
		ids[0] = strings.TrimPrefix(vars["imdb"], "tt")
		langs := strings.Split(vars["lang"], ",")

		// Fallback language always English
		if len(langs) == 0 {
			langs = append(langs, "eng")
		}

		log.Println("Search subtitle by imdbid...")

		season, err := strconv.ParseInt(vars["season"], 10, 64)
		episode, err := strconv.ParseInt(vars["episode"], 10, 64)

		params := []interface{}{}
		if season == 0 && episode == 0 {
			params = []interface{}{
				c.Token,
				[]struct {
					Imdb  string `xmlrpc:"imdbid"`
					Langs string `xmlrpc:"sublanguageid"`
				}{{
					ids[0],
					strings.Join(langs, ","),
				}},
			}
		} else {
			params = []interface{}{
				c.Token,
				[]struct {
					Imdb    string `xmlrpc:"imdbid"`
					Langs   string `xmlrpc:"sublanguageid"`
					Season  int64  `xmlrpc:"season"`
					Episode int64  `xmlrpc:"episode"`
				}{{
					ids[0],
					strings.Join(langs, ","),
					season,
					episode,
				}},
			}
		}

		res, err := c.SearchSubtitles(&params)
		if err != nil {
			http.Error(w, noSubtitlesFound(), http.StatusNotFound)
			return
		}

		if len(res) == 0 {
			http.Error(w, noSubtitlesFound(), http.StatusNotFound)
			return
		}

		found := false
		for _, f := range res {
			if f.SubFormat == "srt" {
				found = true
				break
			}
		}

		if found == false {
			http.Error(w, noSubtitlesFound(), http.StatusNotFound)
			return
		}

		log.Println("Subtitle found.")
		io.WriteString(w, subtitleFilesList(r.Host, res, langs[0]))
	})

	router.HandleFunc(apiPrefix+"subtitlesbytext/{text}/lang/{lang}/season/{season}/episode/{episode}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		// Create Opensubtitles client
		c, err := osdb.NewClient()
		if err != nil {
			http.Error(w, failedToConnectToOpenSubtitles(), http.StatusNotFound)
			return
		}

		c.UserAgent = openSubtitlesUserAgent

		// Anonymous Login with UserAgent string will set c.Token when successful
		if err = c.LogIn("", "", ""); err != nil {
			http.Error(w, failedToConnectToOpenSubtitles(), http.StatusNotFound)
			return
		}

		text := vars["text"]
		langs := strings.Split(vars["lang"], ",")

		// Fallback language always English
		if len(langs) == 0 {
			langs = append(langs, "eng")
		}

		log.Println("Search subtitle by text...")

		season, err := strconv.ParseInt(vars["season"], 10, 64)
		episode, err := strconv.ParseInt(vars["episode"], 10, 64)

		params := []interface{}{}
		if season == 0 && episode == 0 {
			params = []interface{}{
				c.Token,
				[]struct {
					Query string `xmlrpc:"query"`
					Langs string `xmlrpc:"sublanguageid"`
				}{{
					text,
					strings.Join(langs, ","),
				}},
			}
		} else {
			params = []interface{}{
				c.Token,
				[]struct {
					Query   string `xmlrpc:"query"`
					Langs   string `xmlrpc:"sublanguageid"`
					Season  int64  `xmlrpc:"season"`
					Episode int64  `xmlrpc:"episode"`
				}{{
					text,
					strings.Join(langs, ","),
					season,
					episode,
				}},
			}
		}

		res, err := c.SearchSubtitles(&params)
		if err != nil {
			http.Error(w, noSubtitlesFound(), http.StatusNotFound)
			return
		}

		if len(res) == 0 {
			http.Error(w, noSubtitlesFound(), http.StatusNotFound)
			return
		}

		found := false
		for _, f := range res {
			if f.SubFormat == "srt" {
				found = true
				break
			}
		}

		if found == false {
			http.Error(w, noSubtitlesFound(), http.StatusNotFound)
			return
		}

		log.Println("Subtitle found.")
		io.WriteString(w, subtitleFilesList(r.Host, res, langs[0]))
	})

	router.HandleFunc(apiPrefix+"subtitlesbyfile/{hash}/{base64path}/lang/{lang}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		if d, err := base64.StdEncoding.DecodeString(vars["base64path"]); err == nil {
			if t, ok := activeTorrents[vars["hash"]]; ok {

				idx := getFileByPath(string(d), t.torrent.Files())
				file := t.torrent.Files()[idx]

				path := file.DisplayPath()
				log.Println("Calculate Opensubtitles hash...")

				increaseFileClients(path, t)

				fileHash := calculateOpensubtitlesHash(file)
				log.Println("Opensubtitles hash calculated:", fileHash)

				//stop downloading the file when no connections left
				if decreaseFileClients(path, t) <= 0 {
					stopDownloadFile(file)
				}

				// Create Opensubtitles client
				c, err := osdb.NewClient()
				if err != nil {
					http.Error(w, failedToConnectToOpenSubtitles(), http.StatusNotFound)
					return
				}

				c.UserAgent = openSubtitlesUserAgent

				// Anonymous Login with UserAgent string will set c.Token when successful
				if err = c.LogIn("", "", ""); err != nil {
					http.Error(w, failedToConnectToOpenSubtitles(), http.StatusNotFound)
					return
				}

				langs := strings.Split(vars["lang"], ",")

				// Fallback language always English
				if len(langs) == 0 {
					langs = append(langs, "eng")
				}

				params := []interface{}{
					c.Token,
					[]struct {
						Hash  string `xmlrpc:"moviehash"`
						Size  int64  `xmlrpc:"moviebytesize"`
						Langs string `xmlrpc:"sublanguageid"`
					}{{
						fileHash,
						file.Length(),
						strings.Join(langs, ","),
					}},
				}

				res, err := c.SearchSubtitles(&params)
				if err != nil {
					http.Error(w, noSubtitlesFound(), http.StatusNotFound)
					return
				}

				if len(res) == 0 {
					http.Error(w, noSubtitlesFound(), http.StatusNotFound)
					return
				}

				found := false
				for _, f := range res {
					if f.SubFormat == "srt" {
						found = true
						break
					}
				}

				if found == false {
					http.Error(w, noSubtitlesFound(), http.StatusNotFound)
					return
				}

				log.Println("Subtitle found.")
				io.WriteString(w, subtitleFilesList(r.Host, res, langs[0]))
			} else {
				http.Error(w, noSubtitlesFound(), http.StatusNotFound)
				log.Println("Unknown torrent:", vars["hash"])
				return
			}
		} else {
			http.Error(w, noSubtitlesFound(), http.StatusNotFound)
			log.Println(err)
			return
		}
	})

	router.HandleFunc(apiPrefix+"getsubtitle/{base64path}/encode/{encode}/subtitle.{subtype}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		if subtitleurl, err := base64.StdEncoding.DecodeString(vars["base64path"]); err == nil {

			zipContent, err := fetchZip(string(subtitleurl))
			if err != nil {
				http.Error(w, failedToLoadSubtitle(), http.StatusNotFound)
				return
			}

			for _, f := range zipContent.File {
				if strings.HasSuffix(strings.ToLower(f.Name), ".srt") == true {
					fileHandler, err := f.Open()
					if err != nil {
						http.Error(w, failedToLoadSubtitle(), http.StatusNotFound)
						return
					}
					data, err := ioutil.ReadAll(fileHandler)
					if err != nil {
						http.Error(w, failedToLoadSubtitle(), http.StatusNotFound)
						return
					}
					fileHandler.Close()

					// Remove UTF BOM
					if data[0] == 0xef && data[1] == 0xbb && data[2] == 0xbf {
						data = bytes.Trim(data, "\xef\xbb\xbf")
					}

					srt := decodeData(data, vars["encode"])

					subtitle, err := subtitles.NewFromSRT(srt)
					if err != nil {
						http.Error(w, failedToLoadSubtitle(), http.StatusNotFound)
						return
					}

					if vars["subtype"] == "srt" {
						w.Header().Set("Content-Disposition", "filename=subtitle.srt")
						w.Header().Set("Content-Type", "text/plain; charset=utf-8")
						io.WriteString(w, subtitle.RemoveAds().AsSRT())
					} else if vars["subtype"] == "vtt" {
						w.Header().Set("Content-Disposition", "filename=subtitle.vtt")
						w.Header().Set("Content-Type", "text/vtt; charset=utf-8")
						io.WriteString(w, subtitle.RemoveAds().AsVTT())
					} else {
						http.Error(w, failedToLoadSubtitle(), http.StatusNotFound)
						return
					}

					break
				}
			}
		} else {
			http.Error(w, failedToLoadSubtitle(), http.StatusNotFound)
		}
	})

	router.HandleFunc(apiPrefix+"torrents", func(w http.ResponseWriter, r *http.Request) {
		if len(activeTorrents) > 0 {
			io.WriteString(w, activeTorrentsList())
		} else {
			http.Error(w, noActiveTorrentsFound(), http.StatusNotFound)
			log.Println("No active torrents found.")
		}
	})

	router.HandleFunc(apiPrefix+"getmoviemagnet/imdb/{imdb}/providers/{providers}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		log.Printf("Getting movie magnet link by this imdb id: %v\n", vars["imdb"])

		output := torrents.GetMovieTorrents(vars["imdb"], "", strings.Split(vars["providers"], ","))
		if len(output) > 0 {
			io.WriteString(w, movieTorrentsList(output))
			log.Printf("Magnet link found.\n")
		} else {
			http.Error(w, noTorrentsFound(), http.StatusNotFound)
			log.Printf("Not found any magnet link.\n")
		}
	})

	router.HandleFunc(apiPrefix+"getmoviemagnet/query/{query}/providers/{providers}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		log.Printf("Getting movie magnet link by this query: %v\n", vars["query"])

		output := torrents.GetMovieTorrents("", vars["query"], strings.Split(vars["providers"], ","))
		if len(output) > 0 {
			io.WriteString(w, movieTorrentsList(output))
			log.Printf("Magnet link found.\n")
		} else {
			http.Error(w, noTorrentsFound(), http.StatusNotFound)
			log.Printf("Not found any magnet link.\n")
		}
	})

	router.HandleFunc(apiPrefix+"getmoviemagnet/imdb/{imdb}/query/{query}/providers/{providers}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		log.Printf("Getting movie magnet link by this imdb id: %v, query: %v\n", vars["imdb"], vars["query"])

		output := torrents.GetMovieTorrents(vars["imdb"], vars["query"], strings.Split(vars["providers"], ","))
		if len(output) > 0 {
			io.WriteString(w, movieTorrentsList(output))
			log.Printf("Magnet link found.\n")
		} else {
			http.Error(w, noTorrentsFound(), http.StatusNotFound)
			log.Printf("Not found any magnet link.\n")
		}
	})

	router.HandleFunc(apiPrefix+"getshowmagnet/imdb/{imdb}/season/{season}/episode/{episode}/providers/{providers}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		log.Printf("Getting tv show magnet link by this imdb id: %v, season: %v, episode: %v\n", vars["imdb"], vars["season"], vars["episode"])

		output := torrents.GetShowTorrents(vars["imdb"], "", vars["season"], vars["episode"], strings.Split(vars["providers"], ","))
		if len(output) > 0 {
			io.WriteString(w, showTorrentsList(output))
			log.Printf("Magnet link found.\n")
		} else {
			http.Error(w, noTorrentsFound(), http.StatusNotFound)
			log.Printf("Not found any magnet link.\n")
		}
	})

	router.HandleFunc(apiPrefix+"getshowmagnet/query/{query}/season/{season}/episode/{episode}/providers/{providers}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		log.Printf("Getting tv show magnet link by this query: %v, season: %v, episode: %v\n", vars["query"], vars["season"], vars["episode"])

		output := torrents.GetShowTorrents("", vars["query"], vars["season"], vars["episode"], strings.Split(vars["providers"], ","))
		if len(output) > 0 {
			io.WriteString(w, showTorrentsList(output))
			log.Printf("Magnet link found.\n")
		} else {
			http.Error(w, noTorrentsFound(), http.StatusNotFound)
			log.Printf("Not found any magnet link.\n")
		}
	})

	router.HandleFunc(apiPrefix+"getshowmagnet/imdb/{imdb}/query/{query}/season/{season}/episode/{episode}/providers/{providers}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		log.Printf("Getting tv show magnet link by this imdb id: %v, query: %v, season: %v, episode: %v\n", vars["imdb"], vars["query"], vars["season"], vars["episode"])

		output := torrents.GetShowTorrents(vars["imdb"], vars["query"], vars["season"], vars["episode"], strings.Split(vars["providers"], ","))
		if len(output) > 0 {
			io.WriteString(w, showTorrentsList(output))
			log.Printf("Magnet link found.\n")
		} else {
			http.Error(w, noTorrentsFound(), http.StatusNotFound)
			log.Printf("Not found any magnet link.\n")
		}
	})

	router.HandleFunc(apiPrefix+"tmdbdiscover/type/{type}/genretype/{genretype}/sort/{sort}/date/{date}/lang/{lang}/page/{page}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		log.Println("Get TMDB list by genre")

		output := tmdb.Discover(vars["type"], vars["genretype"], vars["sort"], vars["date"], vars["lang"], vars["page"])
		if output != "" {
			io.WriteString(w, discoveredList(output))
		} else {
			http.Error(w, noTmdbDataFound(), http.StatusNotFound)
		}
	})

	router.HandleFunc(apiPrefix+"tmdbsearch/type/{type}/lang/{lang}/page/{page}/text/{text}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		log.Println("Get TMDB search")

		output := tmdb.Search(vars["type"], vars["lang"], vars["page"], vars["text"])
		if output != "" {
			io.WriteString(w, discoveredList(output))
		} else {
			http.Error(w, noTmdbDataFound(), http.StatusNotFound)
		}
	})

	router.HandleFunc(apiPrefix+"tmdbinfo/type/{type}/tmdbid/{tmdbid}/lang/{lang}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		log.Println("Get TMDB info")

		output := tmdb.GetInfo(vars["type"], vars["tmdbid"], vars["lang"])
		if output != "" {
			io.WriteString(w, discoveredList(output))
		} else {
			http.Error(w, noTmdbDataFound(), http.StatusNotFound)
		}
	})

	router.HandleFunc(apiPrefix+"tvmazeepisodes/tvdb/{tvdb}/imdb/{imdb}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		log.Println("Get TVMaze episodes")

		output := tvmaze.GetEpisodes(vars["tvdb"], vars["imdb"])
		if output != "" {
			io.WriteString(w, showEpisodeList(output))
		} else {
			http.Error(w, noTvMazeDataFound(), http.StatusNotFound)
		}
	})

	router.HandleFunc(apiPrefix+"tvmazeepisodes/tvdb/{tvdb}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		log.Println("Get TVMaze episodes")

		output := tvmaze.GetEpisodes(vars["tvdb"], "")
		if output != "" {
			io.WriteString(w, showEpisodeList(output))
		} else {
			http.Error(w, noTvMazeDataFound(), http.StatusNotFound)
		}
	})

	router.HandleFunc(apiPrefix+"tvmazeepisodes/imdb/{imdb}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		log.Println("Get TVMaze episodes")

		output := tvmaze.GetEpisodes("", vars["imdb"])
		if output != "" {
			io.WriteString(w, showEpisodeList(output))
		} else {
			http.Error(w, noTvMazeDataFound(), http.StatusNotFound)
		}
	})

	router.HandleFunc(apiPrefix+"receivemagnet/{todo}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		io.WriteString(w, checkReceiver(vars["todo"]))
	})

	router.HandleFunc(apiPrefix+"websocket", func(w http.ResponseWriter, r *http.Request) {
		upgrader.CheckOrigin = func(r *http.Request) bool { return true }

		webSocket, _ = upgrader.Upgrade(w, r, nil) // Error ignored

		for {
			// Read message from ws
			messageType, message, err := webSocket.ReadMessage()
			if err != nil {
				return
			}

			if messageType == 1 {
				if string(message) == "stop" {
					if err = webSocket.WriteMessage(1, []byte("{\"function\":\"stopserver\",\"data\": \"ok\"}")); err != nil {
						return
					}

					go func() {
						time.Sleep(1 * time.Nanosecond)
						procQuit <- true
					}()
				} else {
					value := setReceivedMagnet(string(message))
					if err = webSocket.WriteMessage(1, []byte("{\"function\":\"sendmagnet\",\"data\":\""+value+"\"}")); err != nil {
						return
					}
				}
			} else if messageType == 2 {
				metaData, error := metainfo.Load(bytes.NewReader(message))
				if error == nil {
					value := setReceivedTorrent(metaData)
					if err = webSocket.WriteMessage(1, []byte("{\"function\":\"sendfile\",\"data\":\""+value+"\"}")); err != nil {
						return
					}
				}
			}

		}
	})

	// Create torrent magnet send page from main page
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, createReceiverPage())
	})

	// Enable CORS for api urls if required
	if cors == false {
		return router
	} else {
		return handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}), handlers.AllowedOrigins([]string{"*"}))(router)
	}
}

func fetchZip(zipurl string) (*zip.Reader, error) {
	req, err := http.NewRequest("GET", zipurl, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", openSubtitlesUserAgent)
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{Transport: tr}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		b, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, errors.New(resp.Status)
		}
		return nil, errors.New(string(b))
	}

	buf := &bytes.Buffer{}

	_, err = io.Copy(buf, resp.Body)
	if err != nil {
		return nil, err
	}

	b := bytes.NewReader(buf.Bytes())
	return zip.NewReader(b, int64(b.Len()))
}

func decodeData(encData []byte, enc string) string {
	var dec *encoding.Decoder

	switch enc {
	case "CP1250":
		dec = charmap.Windows1250.NewDecoder()
	case "CP1251":
		dec = charmap.Windows1251.NewDecoder()
	case "CP1252":
		dec = charmap.Windows1252.NewDecoder()
	case "CP1253":
		dec = charmap.Windows1253.NewDecoder()
	case "CP1254":
		dec = charmap.Windows1254.NewDecoder()
	case "CP1255":
		dec = charmap.Windows1255.NewDecoder()
	case "CP1256":
		dec = charmap.Windows1256.NewDecoder()
	case "CP1257":
		dec = charmap.Windows1257.NewDecoder()
	case "CP1258":
		dec = charmap.Windows1258.NewDecoder()
	default:
		return string(encData)
	}

	out, _ := dec.Bytes(encData)
	return string(out)
}

func getLocalIP() string {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return "127.0.0.1"
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP.String()
}
