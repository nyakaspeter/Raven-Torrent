package dlnacast

import (
	"fmt"
	"strings"

	"github.com/koron/go-ssdp"
	"github.com/nyakaspeter/raven-torrent/internal/settings"
	"github.com/nyakaspeter/raven-torrent/pkg/dlnacast/types"
	"github.com/nyakaspeter/raven-torrent/pkg/utils"
)

func GetMediaDevices() []types.MediaDevice {
	devices := []types.MediaDevice{}
	list, err := ssdp.Search(ssdp.All, 1, "")

	if err != nil {
		return []types.MediaDevice{}
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
				if friendlyName, err := GetDeviceFriendlyName(srv.Location); err != nil {
					devices = append(devices, types.MediaDevice{Name: srv.Server, Location: srv.Location})
				} else {
					devices = append(devices, types.MediaDevice{Name: friendlyName, Location: srv.Location})
				}
			}
		}
	}

	return devices
}

func CastMediaToDevice(media types.MediaParams, deviceLocation string) error {
	host := utils.GetLocalIP()

	transportURL, controlURL, err := GetAvTransportUrl(deviceLocation)
	if err != nil {
		return err
	}

	media.VideoUrl = strings.Replace(media.VideoUrl, "localhost", host, 1)
	media.SubtitleUrl = strings.Replace(media.SubtitleUrl, "localhost", host, 1)

	whereToListen := fmt.Sprintf("%s:%d", host, *settings.DlnaPort)
	callbackURL := fmt.Sprintf("http://%s/callback", whereToListen)

	newPayload := &TVPayload{
		TransportURL: transportURL,
		ControlURL:   controlURL,
		CallbackURL:  callbackURL,
		VideoURL:     media.VideoUrl,
		SubtitlesURL: media.SubtitleUrl,
		VideoTitle:   media.Title,
	}

	if Server == nil {
		srv := CreateServer(whereToListen)
		Server = &srv
	} else {
		TvPayload.SendtoTV("Stop")
		Server.StopServer()
	}

	TvPayload = newPayload

	serverStarted := make(chan struct{})
	go Server.StartServer(serverStarted, newPayload)

	// Wait for HTTP server to properly initialize
	<-serverStarted

	err = newPayload.SendtoTV("Play1")

	if err != nil {
		return err
	}

	return nil
}
