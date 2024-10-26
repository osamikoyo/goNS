package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run dns_client.go <domain>")
		return
	}

	domain := os.Args[1]

	udpAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:1053")
	if err != nil {
		fmt.Println(err)
		return
	}

	conn, err := net.DialUDP("udp", nil, udpAddr)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	_, err = conn.Write([]byte(domain))
	if err != nil {
		fmt.Println(err)
		return
	}

	buffer := make([]byte, 512)
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(buffer[:n]))
}