package main

import (
	"flag"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"os/signal"

	"github.com/anacrolix/torrent"
)

type serviceSettings struct {
	Host           *string
	Port           *int
	DownloadDir    *string
	DownloadRate   *int
	UploadRate     *int
	MaxConnections *int
	NoDHT          *bool
	EnableLog      *bool
	StorageType    *string
	MemorySize     *int64
	Background     *bool
	CORS           *bool
	TMDBKey        *string
	OSUserAgent    *string
	JackettAddress *string
	JackettKey     *string
}

var procQuit chan bool
var procError chan string
var procRestart chan []int64

var cl *torrent.Client

var originalArgs []string

func quit(srv *http.Server) {
	log.Println("Quitting")

	srv.Close()

	if cl != nil {
		cl.Close()
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

func waitingForSignals(saveSettings serviceSettings, receivedArgs []int64, srv *http.Server) {
	if receivedArgs[0] != -1 && receivedArgs[1] != -1 {
		*saveSettings.DownloadRate = int(receivedArgs[0])
		*saveSettings.UploadRate = int(receivedArgs[1])
	}

	select {
	case err := <-procError:
		log.Println("Error:", err)
		quit(srv)

	case <-procQuit:
		quit(srv)

	case receivedArgs := <-procRestart:
		if receivedArgs[0] != -1 && receivedArgs[1] != -1 {
			log.Println("Restarting torrent client with new settings.")
		} else {
			log.Println("Restarting torrent client because torrent deletion.")
		}

		cl.Close()

		select {
		case err := <-procError:
			log.Println("Error:", err)
			quit(srv)

		case <-procQuit:
			quit(srv)

		case <-cl.Closed():
			cl = nil
			cl = startTorrentClient(saveSettings)

			waitingForSignals(saveSettings, receivedArgs, srv)
		}
	}
}

func main() {
	procQuit = make(chan bool)
	procError = make(chan string)
	procRestart = make(chan []int64, 2)

	var settings serviceSettings

	settings.Host = flag.String("host", "", "listening server ip")
	settings.Port = flag.Int("port", 9000, "listening port")
	settings.DownloadDir = flag.String("dir", "", "specify the directory where files will be downloaded to if storagetype is set to \"piecefile\" or \"file\"")
	settings.DownloadRate = flag.Int("downrate", 0, "download speed rate in Kbps")
	settings.UploadRate = flag.Int("uprate", 0, "upload speed rate in Kbps")
	settings.MaxConnections = flag.Int("maxconn", 50, "max connections per torrent")
	settings.NoDHT = flag.Bool("nodht", false, "disable dht")
	settings.EnableLog = flag.Bool("log", false, "enable log messages")
	settings.StorageType = flag.String("storagetype", "memory", "select storage type (must be set to \"memory\" or \"piecefile\" or \"file\")")
	settings.Background = flag.Bool("background", false, "run the server in the background")
	settings.CORS = flag.Bool("cors", false, "enable CORS")
	settings.MemorySize = flag.Int64("memorysize", 128, "specify the storage memory size in MB if storagetype is set to \"memory\" (minimum 64)") // 64MB is optimal for TVs
	settings.TMDBKey = flag.String("tmdbkey", "", "set external TMDB API key")
	settings.JackettAddress = flag.String("jackettaddress", "", "set external Jackett API address")
	settings.JackettKey = flag.String("jackettkey", "", "set external Jackett API key")
	settings.OSUserAgent = flag.String("osuseragent", "", "set external OpenSubtitles user agent")

	// Set Opensubtitles server address to http because https not working on Samsung Smart TV
	os.Setenv("OSDB_SERVER", "http://api.opensubtitles.org/xml-rpc")

	flag.Parse()

	handleSignals()

	log.SetFlags(0)
	// Check storage type
	if *settings.StorageType != "memory" && *settings.StorageType != "piecefile" && *settings.StorageType != "file" {
		log.Printf("missing or invalid -storagetype value: \"%s\" (must be set to \"memory\" or \"piecefile\" or \"file\")\nUsage of %s:\n", *settings.StorageType, os.Args[0])
		flag.PrintDefaults()
		os.Exit(2)
	}

	// Check memory size if StorageType is memory
	if *settings.StorageType == "memory" && *settings.MemorySize < 64 {
		log.Printf("the memory size is too small: \"%dMB\" (must be set to minimum 64MB)\nUsage of %s:\n", *settings.MemorySize, os.Args[0])
		flag.PrintDefaults()
		os.Exit(2)
	}

	// Check dir flag settings if storage type is piecefile or file
	if *settings.StorageType == "piecefile" || *settings.StorageType == "file" {
		if *settings.DownloadDir == "" {
			log.Printf("empty -dir value (must be set if selected -storagetype is \"piecefile\" or \"file\")\nUsage of %s:\n", os.Args[0])
			flag.PrintDefaults()
			os.Exit(2)
		}
	}

	// Set TMDB API key
	if *settings.TMDBKey != "" {
		setTMDBKey(*settings.TMDBKey)
	} else {
		setTMDBKey(TMDBKey)
	}

	// Configure Jackett API
	if *settings.JackettAddress != "" && *settings.JackettKey != "" {
		setJackettAddressAndKey(*settings.JackettAddress, *settings.JackettKey)
	}

	// Set OpenSubtitles user agent string
	if *settings.OSUserAgent != "" {
		setOSUserAgent(*settings.OSUserAgent)
	}

	// Disable or enable the log in production mode
	if *settings.EnableLog == false {
		log.SetOutput(ioutil.Discard)
		defer log.SetOutput(os.Stderr)
	}

	originalArgs = os.Args[1:]
	// Check if need to run in the background
	if *settings.Background == true {
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

	cl = startTorrentClient(settings)
	srv := startHTTPServer(*settings.Host, *settings.Port, *settings.CORS)

	waitingForSignals(settings, []int64{int64(*settings.DownloadRate), int64(*settings.UploadRate)}, srv)
}
