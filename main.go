package main

import (
	"fmt"
	"net"

	"gons/handler"
)

func main() {
	udpAddress, err := net.ResolveUDPAddr("udp", "127.0.0.1:53")
	if err != nil {
		fmt.Println(err)
		return
	}

	conn, err := net.ListenUDP("udp", udpAddress)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	buffer := make([]byte, 512)

	for {
		n, addr, err := conn.ReadFromUDP(buffer)
		if err != nil {
			fmt.Println(err)
			continue
		}

		go handler.HandleDNSRequest(conn, addr, buffer[:n])
	}
}
