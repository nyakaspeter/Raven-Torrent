package dlnacast

import (
	"encoding/xml"
	"errors"
	"io"
	"net/http"
	"net/url"
)

type MediaRenderer struct {
	XMLName     xml.Name `xml:"root"`
	Text        string   `xml:",chardata"`
	Xmlns       string   `xml:"xmlns,attr"`
	Dlna        string   `xml:"dlna,attr"`
	Pnpx        string   `xml:"pnpx,attr"`
	Df          string   `xml:"df,attr"`
	SpecVersion struct {
		Text  string `xml:",chardata"`
		Major string `xml:"major"`
		Minor string `xml:"minor"`
	} `xml:"specVersion"`
	Device struct {
		Text             string `xml:",chardata"`
		DeviceType       string `xml:"deviceType"`
		FriendlyName     string `xml:"friendlyName"`
		Manufacturer     string `xml:"manufacturer"`
		ManufacturerURL  string `xml:"manufacturerURL"`
		ModelDescription string `xml:"modelDescription"`
		ModelName        string `xml:"modelName"`
		ModelURL         string `xml:"modelURL"`
		ModelNumber      string `xml:"modelNumber"`
		SerialNumber     string `xml:"serialNumber"`
		UDN              string `xml:"UDN"`
		IconList         struct {
			Text string `xml:",chardata"`
			Icon []struct {
				Text     string `xml:",chardata"`
				Mimetype string `xml:"mimetype"`
				Width    string `xml:"width"`
				Height   string `xml:"height"`
				Depth    string `xml:"depth"`
				URL      string `xml:"url"`
			} `xml:"icon"`
		} `xml:"iconList"`
		ServiceList struct {
			Text    string `xml:",chardata"`
			Service []struct {
				Text        string `xml:",chardata"`
				ServiceType string `xml:"serviceType"`
				ServiceId   string `xml:"serviceId"`
				SCPDURL     string `xml:"SCPDURL"`
				ControlURL  string `xml:"controlURL"`
				EventSubURL string `xml:"eventSubURL"`
			} `xml:"service"`
		} `xml:"serviceList"`
	} `xml:"device"`
}

type EventPropertySet struct {
	XMLName       xml.Name `xml:"propertyset"`
	EventInstance struct {
		XMLName                 xml.Name `xml:"InstanceID"`
		Value                   string   `xml:"val,attr"`
		CurrentTransportActions struct {
			Value string `xml:"val,attr"`
		} `xml:"CurrentTransportActions"`
		TransportState struct {
			Value string `xml:"val,attr"`
		} `xml:"TransportState"`
	} `xml:"property>LastChange>Event>InstanceID"`
}

// Get the device friendly name from the main DMR xml.
func GetDeviceFriendlyName(dmrurl string) (string, error) {
	var renderer MediaRenderer

	client := &http.Client{}
	req, err := http.NewRequest("GET", dmrurl, nil)
	if err != nil {
		return "", err
	}

	xmlresp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer xmlresp.Body.Close()

	xmlbody, err := io.ReadAll(xmlresp.Body)
	if err != nil {
		return "", err
	}

	err = xml.Unmarshal(xmlbody, &renderer)
	if err != nil {
		return "", err
	}

	return renderer.Device.FriendlyName, nil
}

// Get the AVTransport URL from the main DMR xml.
func GetAvTransportUrl(dmrurl string) (string, string, error) {
	var renderer MediaRenderer

	parsedURL, err := url.Parse(dmrurl)
	if err != nil {
		return "", "", err
	}

	client := &http.Client{}
	req, err := http.NewRequest("GET", dmrurl, nil)
	if err != nil {
		return "", "", err
	}

	xmlresp, err := client.Do(req)
	if err != nil {
		return "", "", err
	}
	defer xmlresp.Body.Close()

	xmlbody, err := io.ReadAll(xmlresp.Body)
	if err != nil {
		return "", "", err
	}
	xml.Unmarshal(xmlbody, &renderer)
	for i := 0; i < len(renderer.Device.ServiceList.Service); i++ {
		if renderer.Device.ServiceList.Service[i].ServiceId == "urn:upnp-org:serviceId:AVTransport" {
			avtransportControlURL := parsedURL.Scheme + "://" + parsedURL.Host + renderer.Device.ServiceList.Service[i].ControlURL
			avtransportEventSubURL := parsedURL.Scheme + "://" + parsedURL.Host + renderer.Device.ServiceList.Service[i].EventSubURL
			return avtransportControlURL, avtransportEventSubURL, nil
		}
	}

	return "", "", errors.New("something broke somewhere - wrong DMR URL?")
}

// Parse the Notify messages from the media renderer.
func ParseNotifyEvent(xmlbody string) (string, string, error) {
	var root EventPropertySet
	err := xml.Unmarshal([]byte(xmlbody), &root)
	if err != nil {
		return "", "", err
	}
	previousstate := root.EventInstance.CurrentTransportActions.Value
	newstate := root.EventInstance.TransportState.Value

	return previousstate, newstate, nil
}
