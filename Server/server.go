package main

import (
	"fmt"
	"net"
	"strings"
)

type Server struct {
	ip      net.IP
	port    int
	handler IHandler
}

func NewServer(ip net.IP, port int) *Server {
	server := &Server{
		ip:      ip,
		port:    port,
		handler: &StrHandler{},
	}

	return server
}

func (server *Server) SetHandler(handler IHandler) {
	server.handler = handler
}

func (server *Server) Listen() {
	//listen start
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", server.ip, server.port))
	if err != nil {
		fmt.Println("listener start fail ", err)
		return
	}
	//close listen
	defer listener.Close()

	for {
		fmt.Println("addr:", listener.Addr())
		//accept
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("accept fail ", err)
			continue
		}
		//handler
		go server.handler.Handle(conn)
	}

}

func GetLocalIp() (ip string) {
	conn, err := net.Dial("udp", "8.8.8.8:53")
	if err != nil {
		fmt.Errorf("dial fail,err:%s", err)
		return
	}

	addr := conn.LocalAddr().String()
	ip = strings.Split(addr, ":")[0]
	return
}
