package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"os/signal"

	"github.com/nyakaspeter/raven-torrent/internal/httpserver"
	"github.com/nyakaspeter/raven-torrent/internal/settings"
	"github.com/nyakaspeter/raven-torrent/internal/torrentclient"
)

var quitSignal = make(chan os.Signal, 1)

func main() {
	settings.Init()
	log.SetFlags(0)
	signal.Notify(quitSignal, os.Interrupt)

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

	// Disable or enable the log in production mode
	if !*settings.EnableLog {
		log.SetOutput(ioutil.Discard)
		defer log.SetOutput(os.Stderr)
	}

	// Check if need to run in the background
	if *settings.Background {
		args := os.Args[1:]
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

	_, err := torrentclient.StartTorrentClient()
	if err != nil {
		quit()
	}

	httpserver.StartHttpServer(quitSignal)

	<-quitSignal
	quit()
}

func quit() {
	log.Println("Quitting.")
	httpserver.StopHttpServer()
	torrentclient.StopTorrentClient()
}
