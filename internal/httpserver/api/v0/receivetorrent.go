package v0

import (
	"bytes"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/anacrolix/torrent/metainfo"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"github.com/nyakaspeter/raven-torrent/internal/torrentclient"
)

var webSocket *websocket.Conn
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func ReceiveTorrent() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		io.WriteString(w, torrentclient.CheckReceiver(vars["todo"]))
	}
}

func Websocket(quitSignal chan os.Signal) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
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
						quitSignal <- os.Kill
					}()
				} else {
					value := torrentclient.SetReceivedMagnet(string(message))
					if err = webSocket.WriteMessage(1, []byte("{\"function\":\"sendmagnet\",\"data\":\""+value+"\"}")); err != nil {
						return
					}
				}
			} else if messageType == 2 {
				metaData, error := metainfo.Load(bytes.NewReader(message))
				if error == nil {
					value := torrentclient.SetReceivedTorrent(metaData)
					if err = webSocket.WriteMessage(1, []byte("{\"function\":\"sendfile\",\"data\":\""+value+"\"}")); err != nil {
						return
					}
				}
			}

		}
	}
}
