package main

import (
	"LiteSender/Common"
	"fmt"
	"io"
	"log"
	"net"
	"os"
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

func (c *Client) SendFile(conf *Common.Config) {
	defer c.conn.Close()
	buf := make([]byte, 1024)

	_, err := c.conn.Write([]byte(conf.FileName))
	if err != nil {
		log.Fatal("write filename error")
	}

	n, err := c.conn.Read(buf)
	if err != nil {
		log.Fatalf("read fail:%s", err)
	}
	resp := string(buf[:n])
	log.Println("receive data from server:", resp)
	if resp == "ready" {
		SendData(c, conf)
	}
}

func SendData(c *Client, conf *Common.Config) {
	buf := make([]byte, 4096)
	file, err := os.Open(conf.FilePath)
	defer file.Close()
	if err != nil {
		log.Fatalf("open file fail:%s", err)
		return
	}
	for {
		n, err := file.Read(buf)
		if err == io.EOF {
			log.Println("read file done!")
			return
		} else if err != nil {
			log.Fatalf("read file fail:%s", err)
			return
		}
		_, err = c.conn.Write(buf[:n])
		if err != nil {
			log.Fatalf("connect write error:%s", err)
			return
		}
	}
}
