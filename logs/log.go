package logs

import (
	"log"
	"os"
)

var (
	instanceName = ""
)

// InitLogStyle 初始化日志格式
func InitLogStyle() {
	log.SetFlags(log.Ldate | log.Lshortfile)
	log.SetFlags(log.Ldate | log.Ltime)
	instanceName, _ = os.Hostname()
}

func Println(a ...any) {
	log.Println(a)
}
