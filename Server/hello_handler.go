package main

import (
	"fmt"
	"log"
	"net"
)

type HelloHandler struct {
	conn net.Conn
}

func (h *HelloHandler) Handle(conn net.Conn) {
	if h == nil {
		log.Println("<nil>")
		return
	}

	b := make([]byte, 1024)
	h.conn = conn
	for {
		n, err := conn.Read(b)
		if err != nil {
			log.Println("read fail", err)
			return
		}
		uname := string(b[:n])
		log.Printf("receive from %s\n", uname)
		msg := fmt.Sprintf("hello! %s.I am Bob!", uname)
		h.conn.Write([]byte(msg))
	}

}
