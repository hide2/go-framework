package config

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
)

// 配置文件
type MySQLConfig struct {
	Write string
	Read  string
}
type Config struct {
	Env       string
	Listen    string
	Mysql     MySQLConfig
	Redis     string
	Accesslog string
	ErrorLog  string
}

var GlobalConfig *Config

func init() {
	// 解析命令行参数
	var configFile, accessLog, errorLog string
	flag.StringVar(&configFile, "c", "config.json", "set configuration `file`")
	flag.StringVar(&accessLog, "a", "/tmp/access.log", "set access log `file`")
	flag.StringVar(&errorLog, "e", "/tmp/error.log", "set error log `file`")
	flag.Parse()

	// 加载配置文件config.json
	data, err := ioutil.ReadFile(configFile)
	if err != nil {
		fmt.Println(err)
		return
	}
	config := Config{}
	err = json.Unmarshal(data, &config)
	if err != nil {
		fmt.Println(err)
		return
	}
	config.Accesslog = accessLog
	config.ErrorLog = errorLog
	GlobalConfig = &config
}
