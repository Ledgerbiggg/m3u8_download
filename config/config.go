package config

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"strings"
)

// Configs 使用全局的配置变量
var Configs *DWConfig

// LoadConfig viper读取yaml
func LoadConfig() error {
	// yaml
	vconfig := viper.New()
	//表示 先预加载匹配的环境变量
	vconfig.AutomaticEnv()
	//设置环境变量分割符，将点号和横杠替换为下划线
	vconfig.SetEnvKeyReplacer(strings.NewReplacer("-", "_", ".", "_"))
	// 设置读取的配置文件
	vconfig.SetConfigName("config")
	// 添加读取的配置文件路径
	vconfig.AddConfigPath(".")
	// 读取文件类型
	vconfig.SetConfigType("yaml")

	err := vconfig.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %s", err))
	}
	var cng DWConfig
	if err := vconfig.Unmarshal(&cng); err != nil {
		log.Panicln("unmarshal cng file fail " + err.Error())
	}
	// 赋值全局变量
	Configs = &cng
	return err
}

func init() {
	err := LoadConfig()
	if err != nil {
		log.Println("load config fail " + err.Error())
	}
}

// DWConfig 系统整体配置
type DWConfig struct {
	Mu8 mu8 `yaml:"mu8"`
}

type mu8 struct {
	Url    string `yaml:"url"`
	Thread int    `yaml:"thread"`
}
