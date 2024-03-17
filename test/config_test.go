package test

import (
	"Down_m3u8/config"
	"Down_m3u8/logs"
	"log"
	"testing"
)

func init() {
	logs.InitLogStyle()
	err := config.LoadConfig()
	if err != nil {
		panic(err)
	}
}

func TestConfig(t *testing.T) {
	configs := config.Configs
	log.Println(configs.Mu8.Url)
}
