package util

import (
	"bufio"
	"bytes"
	"io"
	"os"
	"strings"
)

// ReadMU8 函数用于读取 mu8 文件，并返回不以 '#' 开头的行
func ReadMU8(fileUrl string) (string, string, []string, error) {
	// 打开 mu8 文件
	dos := NewHttpDos(fileUrl, nil, nil, nil)
	get, err := dos.Get()
	if err != nil {
		return "", "", nil, err
	}
	var lines []string
	// 逐行读取文件内容
	scanner := bufio.NewScanner(bytes.NewReader(get))
	var uri string
	var iv string
	split := strings.Split(fileUrl, "/")

	path := strings.Join(split[:len(split)-1], "/")

	for scanner.Scan() {
		line := scanner.Text()
		// 如果不以 '#' 开头，则将其添加到结果中
		if !strings.HasPrefix(line, "#") {
			if !strings.HasPrefix(line, "http") {
				line = path + "/" + line
			}
			lines = append(lines, line)
		}
		if strings.Contains(line, "URI") {
			split = strings.Split(line, `"`)
			uri = split[1]
		}

		if strings.Contains(line, "IV") {
			split := strings.Split(line, `IV=`)
			iv = split[1]
		}
	}
	uri = path + "/" + uri
	// 检查扫描过程是否出错
	if err = scanner.Err(); err != nil {
		return "", "", nil, err
	}

	return uri, iv, lines, nil
}

// ReadKey 函数用于读取 enc.key 文件
func ReadKey(filename string) string {
	file, err := os.Open(filename)
	if err != nil {
		return ""
	}
	defer file.Close()

	all, err := io.ReadAll(file)
	if err != nil {
		return ""
	}
	return string(all)
}
