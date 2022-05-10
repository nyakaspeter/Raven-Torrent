package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"os/signal"

	"github.com/nyakaspeter/raven-torrent/internal/server"
	"github.com/nyakaspeter/raven-torrent/pkg/metadata/tmdb"
	"github.com/nyakaspeter/raven-torrent/pkg/torrents/jackett"
)

type serviceSettings struct {
	Host                   *string
	Port                   *int
	DlnaPort               *int
	DownloadDir            *string
	DownloadRate           *int
	UploadRate             *int
	MaxConnections         *int
	NoDHT                  *bool
	EnableLog              *bool
	StorageType            *string
	MemorySize             *int64
	Background             *bool
	CORS                   *bool
	TMDBKey                *string
	OpenSubtitlesUserAgent *string
	JackettAddress         *string
	JackettKey             *string
}

var procQuit chan bool
var procRestart chan []int64
var originalArgs []string
var settings serviceSettings

func main() {
	settings.Host = flag.String("host", "", "listening server ip")
	settings.Port = flag.Int("port", 9000, "listening port")
	settings.DlnaPort = flag.Int("dlnaport", 3500, "DLNA server port")
	settings.DownloadDir = flag.String("dir", "", "specify the directory where files will be downloaded to if storagetype is set to \"file\"")
	settings.DownloadRate = flag.Int("downrate", 0, "download speed rate in Kbps")
	settings.UploadRate = flag.Int("uprate", 0, "upload speed rate in Kbps")
	settings.MaxConnections = flag.Int("maxconn", 50, "max connections per torrent")
	settings.NoDHT = flag.Bool("nodht", false, "disable dht")
	settings.EnableLog = flag.Bool("log", false, "enable log messages")
	settings.StorageType = flag.String("storagetype", "memory", "select storage type (must be set to \"memory\" or \"file\")")
	settings.Background = flag.Bool("background", false, "run the server in the background")
	settings.CORS = flag.Bool("cors", true, "enable CORS")
	settings.MemorySize = flag.Int64("memorysize", 128, "specify the storage memory size in MB if storagetype is set to \"memory\" (minimum 64)") // 64MB is optimal for TVs
	settings.TMDBKey = flag.String("tmdbkey", "a4d9ad8d2d072c50dc998cc0d1a508fa", "set external TMDB API key")
	settings.JackettAddress = flag.String("jackettaddress", "", "set external Jackett API address")
	settings.JackettKey = flag.String("jackettkey", "", "set external Jackett API key")
	settings.OpenSubtitlesUserAgent = flag.String("osuseragent", "White Raven v0.3", "set external OpenSubtitles user agent")

	procQuit = make(chan bool)
	procRestart = make(chan []int64, 2)
	handleSignals()

	flag.Parse()
	log.SetFlags(0)

	// Set Opensubtitles server address to http because https not working on Samsung Smart TV
	os.Setenv("OSDB_SERVER", "http://api.opensubtitles.org/xml-rpc")

	// Check storage type
	if *settings.StorageType != "memory" && *settings.StorageType != "file" {
		log.Printf("missing or invalid -storagetype value: \"%s\" (must be set to \"memory\" or \"file\")\nUsage of %s:\n", *settings.StorageType, os.Args[0])
		flag.PrintDefaults()
		os.Exit(2)
	}

	// Check memory size if StorageType is memory
	if *settings.StorageType == "memory" && *settings.MemorySize < 64 {
		log.Printf("the memory size is too small: \"%dMB\" (must be set to minimum 64MB)\nUsage of %s:\n", *settings.MemorySize, os.Args[0])
		flag.PrintDefaults()
		os.Exit(2)
	}

	// Check dir flag settings if storage type is file
	if *settings.StorageType == "file" {
		if *settings.DownloadDir == "" {
			log.Printf("empty -dir value (must be set if selected -storagetype is \"file\")\nUsage of %s:\n", os.Args[0])
			flag.PrintDefaults()
			os.Exit(2)
		}
	}

	// Set TMDB API key
	tmdb.TmdbKey = *settings.TMDBKey

	// Configure Jackett API
	jackett.JackettAddress = *settings.JackettAddress
	jackett.JackettKey = *settings.JackettKey

	// Set OpenSubtitles user agent string
	server.OpenSubtitlesUserAgent = *settings.OpenSubtitlesUserAgent

	// Set DLNA server port
	server.DlnaServerPort = *settings.DlnaPort

	// Disable or enable the log in production mode
	if !*settings.EnableLog {
		log.SetOutput(ioutil.Discard)
		defer log.SetOutput(os.Stderr)
	}

	originalArgs = os.Args[1:]
	// Check if need to run in the background
	if *settings.Background {
		args := originalArgs
		// Disable the background argument to false before the start
		for i := 0; i < len(args); i++ {
			if args[i] == "-background=true" || args[i] == "-background" || args[i] == "--background=true" || args[i] == "--background" {
				args[i] = "-background=false"
				break
			}
		}
		// Disable logs when running in the background
		for i := 0; i < len(args); i++ {
			if args[i] == "-log=true" || args[i] == "-log" || args[i] == "--log=true" || args[i] == "--log" {
				args[i] = "-log=false"
				break
			}
		}
		cmd := exec.Command(os.Args[0], args...)
		cmd.Start()
		log.Println("Running in the background with the following PID number:", cmd.Process.Pid)
		os.Exit(0)
	}

	startTorrentClient()
	startHttpServer()
	waitForSignals()
}

func startTorrentClient() {
	var err error = nil

	_, err = server.StartTorrentClient(
		*settings.StorageType,
		*settings.MemorySize,
		*settings.DownloadDir,
		*settings.DownloadRate,
		*settings.UploadRate,
		*settings.MaxConnections,
		*settings.NoDHT,
		*settings.EnableLog,
	)

	if err != nil {
		log.Println("Error:", err.Error())
		procQuit <- true
	}
}

func startHttpServer() {
	server.StartHttpServer(
		*settings.Host,
		*settings.Port,
		*settings.CORS,
		procQuit,
		procRestart,
	)
}

func waitForSignals() {
	select {

	case <-procQuit:
		quit()

	case receivedArgs := <-procRestart:
		if receivedArgs[0] != -1 && receivedArgs[1] != -1 {
			*settings.DownloadRate = int(receivedArgs[0])
			*settings.UploadRate = int(receivedArgs[1])
			log.Println("Restarting torrent client with new settings.")
		} else {
			log.Println("Restarting torrent client because torrent deletion.")
		}

		server.StopTorrentClient()
		startTorrentClient()
		waitForSignals()
	}
}

func handleSignals() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	go func() {
		for range c {
			// ^C, handle it
			procQuit <- true
		}
	}()
}

func quit() {
	log.Println("Quitting")

	server.StopHttpServer()
	server.StopTorrentClient()
}
