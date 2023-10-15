// DLNA casting based on Go2TV (https://github.com/alexballas/Go2TV)

package dlnacast

import (
	"fmt"
	"html"
	"io"
	"log"
	"net"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type HTTPserver struct {
	http *http.Server
	mux  *http.ServeMux
}

var Server *HTTPserver = nil
var TvPayload *TVPayload = nil

func CreateServer(a string) HTTPserver {
	srv := HTTPserver{
		http: &http.Server{Addr: a},
		mux:  nil,
	}

	return srv
}

func (s *HTTPserver) StartServer(serverStarted chan<- struct{}, tvPayload *TVPayload) {
	s.mux = http.NewServeMux()
	s.mux.HandleFunc("/callback", tvPayload.callbackHandler)
	s.http.Handler = s.mux

	ln, err := net.Listen("tcp", s.http.Addr)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Encountered error(s): %s\n", err)
	}

	serverStarted <- struct{}{}
	s.http.Serve(ln)
}

func (s *HTTPserver) StopServer() {
	s.http.Close()
}

func (payload *TVPayload) callbackHandler(w http.ResponseWriter, req *http.Request) {
	reqParsed, _ := io.ReadAll(req.Body)
	sidVal, sidExists := req.Header["Sid"]

	if !sidExists {
		http.Error(w, "", 404)
		return
	}

	if sidVal[0] == "" {
		http.Error(w, "", 404)
		return
	}

	uuid := sidVal[0]
	uuid = strings.TrimLeft(uuid, "[")
	uuid = strings.TrimLeft(uuid, "]")
	uuid = strings.TrimLeft(uuid, "uuid:")

	// Apparently we should ignore the first message
	// On some media renderers we receive a STOPPED message
	// even before we start streaming.
	seq, err := GetSequence(uuid)
	if err != nil {
		http.Error(w, "", 404)
		return
	}

	if seq == 0 {
		IncreaseSequence(uuid)
		fmt.Fprintf(w, "OK\n")
		return
	}

	reqParsedUnescape := html.UnescapeString(string(reqParsed))
	previousstate, newstate, err := ParseNotifyEvent(reqParsedUnescape)
	if err != nil {
		http.Error(w, "", 404)
		return
	}

	if !UpdateMRstate(previousstate, newstate, uuid) {
		http.Error(w, "", 404)
		return
	}

	parsedURLtransport, err := url.Parse(payload.TransportURL)
	if err == nil {
		if newstate == "PLAYING" {
			log.Println(payload.VideoTitle + " playback started on " + parsedURLtransport.Hostname())
		}
		if newstate == "PAUSED_PLAYBACK" {
			log.Println(payload.VideoTitle + " playback paused on " + parsedURLtransport.Hostname())
		}
		if newstate == "STOPPED" {
			log.Println(payload.VideoTitle + " playback stopped on " + parsedURLtransport.Hostname())
			payload.UnsubscribeSoapCall(uuid)
		}
	}

	// We could just not send anything here
	// as the core server package would still
	// default to a 200 OK empty response.
	w.WriteHeader(http.StatusOK)
}
