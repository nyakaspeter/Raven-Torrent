package utils

import (
	"archive/zip"
	"bytes"
	"crypto/tls"
	"errors"
	"io"
	"io/ioutil"
	"net"
	"net/http"

	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/charmap"
)

func FetchZip(zipurl string, useragent string) (*zip.Reader, error) {
	req, err := http.NewRequest("GET", zipurl, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", useragent)
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

func DecodeData(encData []byte, enc string) string {
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

func GetLocalIP() string {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return "127.0.0.1"
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP.String()
}
