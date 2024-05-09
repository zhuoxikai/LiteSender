package main

import (
	"net"
)

func main() {
	ip := net.ParseIP(GetLocalIp())
	server := NewServer(ip, 8899)
	server.SetHandler(&StrHandler{})
	server.Listen()
}
