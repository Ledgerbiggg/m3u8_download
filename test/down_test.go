package test

import (
	"Down_m3u8/config"
	"Down_m3u8/logs"
	"Down_m3u8/service"
	"Down_m3u8/util"
	"testing"
)

func init() {
	logs.InitLogStyle()
	err := config.LoadConfig()
	if err != nil {
		panic(err)
	}
}

func TestDownloadMu8AndDownFile(t *testing.T) {

	mu8, vi, strings, err := util.ReadMU8(config.Configs.Mu8.Url)
	if err != nil {
		t.Error(err)
	}
	// 读取 mu8 文件
	service.DownloadMu8AndDownFile(mu8, vi, strings)

}

func TestDownKey(t *testing.T) {
	util.NewHttpDos("https://cdn1.xlzys.com/play/9aAGQA9d/enc.key", nil, nil, nil).Get()
}

func TestM(t *testing.T) {

	get, err := util.NewHttpDos("https://m3u.haiwaikan.com/xm3u8/286a8da263780575721b36a490557330fe5fcbc155e95a999f864a2f50c98f919921f11e97d0da21.m3u8", nil, nil, nil).Get()
	if err != nil {
		t.Error(err)
	}

	t.Log(string(get))
}
