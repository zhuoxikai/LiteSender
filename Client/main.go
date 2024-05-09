package main

import (
	"LiteSender/Common"
	"log"
)

func main() {
	conf := Common.ReadConfig()

	//创建client
	client := NewClient(conf.Ip, conf.Port)
	if client == nil {
		log.Fatal("连接服务器失败")
		return
	}

	client.SendFile(conf)
}
