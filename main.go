package main

import (
	"Down_m3u8/config"
	"Down_m3u8/logs"
	"Down_m3u8/service"
	"Down_m3u8/util"
)

func init() {
	logs.InitLogStyle()
	err := config.LoadConfig()
	if err != nil {
		panic(err)
	}
}

func main() {
	mu8, vi, strings, err := util.ReadMU8(config.Configs.Mu8.Url)
	if err != nil {
		logs.Println("读取mu8文件失败", err)
		return
	}
	// 读取 mu8 文件
	service.DownloadMu8AndDownFile(mu8, vi, strings)
	logs.Println("完成下载请打开target.ts观看")
}
