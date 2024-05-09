package main

import (
	"log"
	"net"
	"os"
)

type FileHandler struct {
	conn net.Conn
}

func (h *FileHandler) Handle(conn net.Conn) {
	if h == nil {
		log.Println("<nil>")
		return
	}

	b := make([]byte, 1024)
	h.conn = conn
	n, err := conn.Read(b)
	if err != nil {
		log.Println("read fail", err)
		return
	}
	filename := string(b[:n])
	log.Printf("receive from client: %s\n", filename)

	if _, err := conn.Write([]byte("ready")); err != nil {
		log.Fatal("connect Write err:", err)
		return
	}

	if file, err := os.Create(filename); err != nil {
		log.Fatal("Create err:", err)
		return
	} else {
		buf := make([]byte, 4096)
		for {
			var n int
			if n, err = conn.Read(buf); n == 0 {
				log.Println("read data from connect finish!")
				conn.Close()
				file.Close()
				return
			} else if err != nil {
				log.Fatal("read data from connect fail")
			}
			if n, err = file.Write(buf[:n]); err != nil {
				log.Fatal("write data to file fail")
			}
		}

	}

}
