package handler

import (
	"fmt"
	"net"
	"strings"
)

func HandleDNSRequest(conn *net.UDPConn, addr *net.UDPAddr, buffer []byte) {
	domain := strings.TrimSpace(string(buffer))
	ip, err := net.LookupIP(domain)
	if err != nil {
		fmt.Println(err)
		return
	}

	response := []byte(ip[0].String())
	_, err = conn.WriteToUDP(response, addr)
	if err != nil {
		fmt.Println(err)
		return
	}
}