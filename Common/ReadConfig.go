package Common

import (
	"encoding/json"
	"log"
	"net"
	"os"
)

type Config struct {
	Ip   net.IP
	Port int
}

func ReadConfig() (conf *Config) {
	cont, err := os.ReadFile("Config.json")
	if err != nil {
		log.Fatal("read file error:", err)
		return conf.DefaultConfig()
	}

	err = json.Unmarshal(cont, &conf)
	if err != nil {
		log.Fatal("json unmarshal error:", err)

	}

	log.Println("config:", conf)
	return
}

func (conf Config) DefaultConfig() *Config {
	return &Config{
		Ip:   net.ParseIP("127.0.0.1"),
		Port: 8899,
	}
}
