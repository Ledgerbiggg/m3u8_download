package service

import (
	"Down_m3u8/config"
	"Down_m3u8/logs"
	"Down_m3u8/util"
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"strings"
	"sync"
)

var (
	c = make(chan int, config.Configs.Mu8.Thread)
	t = make(chan int, 100000)
	// 名字和字节映射
	nameMap = make(map[string][]byte)
	// 顺序名字
	nameList []string
	// map写锁
	mu sync.RWMutex
	// 并发下载
	wg sync.WaitGroup
	// 并发保存
	wg2 sync.WaitGroup
)

func init() {
	for i := 0; i < config.Configs.Mu8.Thread; i++ {
		c <- i
	}
}

func DownloadMu8AndDownFile(key string, iv string, ts []string) {
	kb, err := util.NewHttpDos(key, nil, nil, nil).Get()
	if err != nil {
		logs.Println("下载失败key", err)
		return
	}
	// 按照顺序存放名称
	for _, v := range ts {
		split := strings.Split(v, "/")
		s := split[len(split)-1]
		nameList = append(nameList, s)
	}

	go func() {
		wg2.Add(1)
		err = splicingTS(nameList)
		if err != nil {
			logs.Println("splicingTS", err)
			return
		}
	}()

	for _, v := range ts {
		wg.Add(1)
		<-c
		go func(v string) {
			defer func() {
				c <- 1
				wg.Done()
			}()
			split := strings.Split(v, "/")
			logs.Println("正在下载", split[len(split)-1])
			get, err := util.NewHttpDos(v, nil, nil, nil).Get()
			if err != nil {
				logs.Println("NewHttpDos", err)
				return
			}
			k := string(kb)
			if iv != "" && k != "" {
				get, err = decrypt(k, iv, get)
				if err != nil {
					logs.Println("decrypt", err)
					return
				}
			}
			if err != nil {
				logs.Println("Copy", err)
				return
			}
			mu.Lock()
			defer func() {
				mu.Unlock()
			}()
			nameMap[split[len(split)-1]] = get
			if len(nameMap) >= 20 {
				t <- 1
			}
		}(v)
	}

	wg.Wait()
	for i := 0; i < len(nameMap); i++ {
		t <- 1
	}
	wg2.Wait()
}

// Splicing
func splicingTS(nameList []string) error {
	// 目标文件
	target, err := os.OpenFile("target.ts", os.O_RDWR|os.O_CREATE, 0666)
	target.Truncate(0)
	if err != nil {
		logs.Println("OpenFile", err)
		return err
	}
	if len(nameList) == 0 {
		<-t
	}
	defer target.Close()
	for _, v := range nameList {
		// 每次重试3次不行就走人
		for i := 0; i < 3; i++ {
			mu.RLock()
			b := nameMap[v]
			mu.RUnlock()
			if b == nil {
				<-t
				continue
			}
			// 追加文件到目标文件
			logs.Println("正在拼接", v)
			err = SaveIntoTarget(target, b)
			if err != nil {
				logs.Println("SaveIntoTarget", err)
				return err
			}
			mu.Lock()
			delete(nameMap, v)
			mu.Unlock()
			break
		}
	}
	wg2.Done()
	return nil
}

func SaveIntoTarget(f *os.File, b []byte) error {
	_, err := f.Seek(0, io.SeekEnd)
	if err != nil {
		logs.Println("Seek", err)
		return err
	}
	_, err = io.Copy(f, bytes.NewReader(b))
	if err != nil {
		logs.Println("Copy", err)
		return err
	}
	return nil
}

// decrypt 解密ts
func decrypt(key, ivHex string, ts []byte) ([]byte, error) {
	// 设置密钥和 IV
	keyBytes := []byte(key)

	// 去除开头的 "0x"
	if len(ivHex) > 2 && ivHex[0:2] == "0x" {
		ivHex = ivHex[2:]
	}

	// 将十六进制字符串转换为字节切片
	ivBytes, err := hex.DecodeString(ivHex)
	if err != nil {
		fmt.Println("Error decoding hex:", err)
		return nil, err
	}
	// 解密文件
	decryptedData, err := decryptAES128(ts, keyBytes, ivBytes)
	if err != nil {
		fmt.Println("Error decrypting file:", err)
		return nil, err
	}
	return decryptedData, nil
}

// 解密函数
func decryptAES128(data []byte, key []byte, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(data, data)

	// 去除填充内容
	paddingLen := int(data[len(data)-1])
	return data[:len(data)-paddingLen], nil
}
