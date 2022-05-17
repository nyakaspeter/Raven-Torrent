package utils

import (
	"net"

	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/charmap"
)

func GetLocalIP() string {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return "127.0.0.1"
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP.String()
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
