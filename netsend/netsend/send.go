package netsend

import (
	"io/ioutil"
	"log"
	"net"
	"os"
	"strings"
)

func Send(msg string) {
	if msg == "" {
		b, _ := ioutil.ReadAll(os.Stdin)
		msg = strings.TrimSpace(string(b))
	}

	if msg == "" {
		log.Print("Empty string, doing nothing")
		return
	}

	sendAddr, err := net.ResolveUDPAddr("udp", "255.255.255.255:8686")
	if err != nil {
		panic(err)
	}
	conn, err := net.DialUDP("udp", nil, sendAddr)
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	b := []byte(msg)
	if len(b) > 1024 {
		b = b[:1024]
	}
	conn.Write(b)
}
