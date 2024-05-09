package Common

import (
	"encoding/json"
	"log"
	"os"
)

type Config struct {
	Ip   string
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
		Ip:   "127.0.0.1",
		Port: 8899,
	}
}
