package netsend

import (
	"log"
	"net"
)

func Recv(messageBox bool) {
	bindAddr, err := net.ResolveUDPAddr("udp", "0.0.0.0:8686")
	if err != nil {
		panic(err)
	}
	conn, err := net.ListenUDP("udp", bindAddr)
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	b := make([]byte, 1024)
	size, err := conn.Read(b)
	if err != nil {
		panic(err)
	}

	msg := string(b[:size])
	if messageBox {
		MessageBox("netsend", msg, MB_ICONINFORMATION|MB_OK)
	} else {
		log.Print(msg)
	}
}
