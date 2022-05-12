package v0

import (
	"fmt"
	"net/http"
	"time"

	"github.com/nyakaspeter/raven-torrent/pkg/utils"
)

var OpenSubtitlesUserAgent string

var procQuit chan bool
var procRestart chan []int64

var serverHost string
var serverPort int
var httpServer *http.Server

func StartHttpServer(host string, port int, cors bool, quit chan bool, restart chan []int64) *http.Server {
	procQuit = quit
	procRestart = restart

	httpServer = &http.Server{
		Addr:         fmt.Sprintf("%s:%d", host, port),
		ReadTimeout:  38 * time.Second,
		WriteTimeout: 38 * time.Second,
		Handler:      handleAPI(cors),
	}

	localIP := host
	if localIP == "" {
		localIP = utils.GetLocalIP()
	}

	serverHost = localIP
	serverPort = port

	address := fmt.Sprintf("http://%s:%d", serverHost, serverPort)

	// Must appear
	fmt.Printf("Raven Torrent started on address: %s\n", address)

	go func() {
		if err := httpServer.ListenAndServe(); err != nil {
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

	return httpServer
}

func StopHttpServer() {
	httpServer.Close()
}
