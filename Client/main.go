package main

import (
	"LiteSender/Common"
	"log"
)

func main() {
	conf := Common.ReadConfig()
	ip := GetLocalIp()

	//创建client
	client := NewClient(ip, conf.Port)
	if client == nil {
		log.Fatal("连接服务器失败")
		return
	}

	client.StartClient()
}
