package main

import (
	"fmt"
	"log"
	"net"
	"strings"
	"sync"
)

type Client struct {
	mu         sync.Mutex
	routerIp   string
	routerPort int
	conn       net.Conn
}

func NewClient(routerIp string, routerPort int) *Client {
	client := new(Client)
	client.routerIp = routerIp
	client.routerPort = routerPort
	//开始连接
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", routerIp, routerPort))
	if err != nil {
		log.Fatalf("连接 %s:%d 失败: %v\n", routerIp, routerPort, err)
		return nil
	}
	client.conn = conn
	return client
}

func GetLocalIp() (ip string) {
	conn, err := net.Dial("udp", "8.8.8.8:53")
	if err != nil {
		log.Fatalf("dial fail,err:%s", err)
		return
	}

	addr := conn.LocalAddr().String()
	ip = strings.Split(addr, ":")[0]
	return
}

func (c *Client) StartClient() {
	defer c.conn.Close()
	buf := make([]byte, 1024)

	for {
		c.conn.Write([]byte("Alice"))
		for {
			n, err := c.conn.Read(buf)
			if err != nil {
				log.Fatalf("read fail:%s", err)
			}
			log.Println("receive data:", string(buf[:n]))
		}
	}
}
