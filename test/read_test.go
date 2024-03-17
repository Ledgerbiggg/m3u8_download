package test

import (
	"Down_m3u8/config"
	"Down_m3u8/logs"
	"Down_m3u8/util"
	"fmt"
	"testing"
)

func init() {
	logs.InitLogStyle()
	err := config.LoadConfig()
	if err != nil {
		panic(err)
	}
}

func TestReadMu8(t *testing.T) {
	// 读取 mu8 文件
	filename := config.Configs.Mu8.Url
	key, _, lines, err := util.ReadMU8(filename)
	if err != nil {
		fmt.Println("Error reading mu8 file:", err)
		return
	}

	// 打印结果
	fmt.Println("Lines read from", filename+":")
	for _, line := range lines {
		fmt.Println(line)
	}

	fmt.Println("Key read from", filename+":", key)
}

func TestReadKey(t *testing.T) {
	// 读取 enc.key 文件
	filename := "enc.key"
	key := util.ReadKey(filename)
	fmt.Println("Key read from", filename+":", key)
}
