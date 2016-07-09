package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"time"

	"./sync"
)

type SyncConfig struct {
	From     *sync.Outline
	To       *sync.Outline
	Interval int
}

type Config struct {
	Sync []SyncConfig
}

func doSync(interval int, from *sync.Outline, to *sync.Outline) {
	for {
		sync.LightSync(from, to)
		time.Sleep(time.Second * time.Duration(interval))
	}
}

func main() {
	config := Config{}
	confdata, err := ioutil.ReadFile(".config/json.conf")
	if err != nil {
		log.Fatalln("read config failed:", err)
	}

	err = json.Unmarshal(confdata, &config)
	if err != nil {
		log.Fatalln("Unmarshal config failed:", err)
	}
	for _, v := range config.Sync {
		go doSync(v.Interval, v.From, v.To)
	}

	time.Sleep(time.Second * 1000)
}
