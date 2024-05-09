package main

import (
	"LiteSender/Common"
	"net"
)

func main() {
	ip := net.ParseIP(GetLocalIp())
	conf := Common.ReadConfig()
	server := NewServer(ip, conf.Port)
	server.SetHandler(&FileHandler{})
	server.Listen()
}
