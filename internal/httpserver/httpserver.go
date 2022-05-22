package httpserver

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/nyakaspeter/raven-torrent/internal/settings"
	"github.com/nyakaspeter/raven-torrent/pkg/utils"
)

var quitSignal chan os.Signal

var serverHost string
var serverPort int
var httpServer *http.Server

func StartHttpServer(appQuitSignal chan os.Signal) *http.Server {
	quitSignal = appQuitSignal

	httpServer = &http.Server{
		Addr:         fmt.Sprintf("%s:%d", *settings.Host, *settings.Port),
		ReadTimeout:  38 * time.Second,
		WriteTimeout: 38 * time.Second,
		Handler:      routesHandler(),
	}

	localIP := *settings.Host
	if localIP == "" {
		localIP = utils.GetLocalIP()
	}

	serverHost = localIP
	serverPort = *settings.Port

	address := fmt.Sprintf("http://%s:%d", serverHost, serverPort)

	go func() {
		if err := httpServer.ListenAndServe(); err != nil {
			// cannot panic, because this probably is an intentional close
			if err != http.ErrServerClosed {
				fmt.Printf("HTTP Server Error: %s\n", err)
			}
			time.Sleep(1 * time.Nanosecond)
			quitSignal <- os.Kill
		}
	}()

	// Must appear
	fmt.Printf("Raven Torrent started on address: %s\n", address)

	return httpServer
}

func StopHttpServer() {
	if httpServer == nil {
		return
	}

	httpServer.Close()
}
