package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"server/libs/auth"
	"server/libs/db"
	"server/libs/logger"
	"server/routes"

	"github.com/gin-gonic/gin"
)

// 配置文件
type MySQLConfig struct {
	Write string
	Read  string
}
type Config struct {
	Env    string
	Listen string
	Mysql  MySQLConfig
	Redis  string
}

func main() {
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

	// Env
	if config.Env == "local" || config.Env == "dev" || config.Env == "test" {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	// ErrorLog
	ef, _ := os.OpenFile(errorLog, os.O_CREATE|os.O_RDWR|os.O_APPEND, os.ModeAppend|os.ModePerm)
	gin.DefaultErrorWriter = ef

	// New Engine
	engine := gin.New()

	// Middleware: Auth
	engine.Use(auth.Auth())

	// 自定义Logger
	if err := logger.InitLogger("info", accessLog, 500, 10, 10); err != nil {
		fmt.Printf("init logger failed, err:%v\n", err)
		return
	}
	engine.Use(logger.GinLogger(logger.Logger), logger.GinRecovery(logger.Logger, true))

	// 异常崩溃处理
	engine.Use(gin.Recovery())

	// 路由配置
	routes.InitRoutes(engine)

	// MySQL
	db.InitDB(config.Env, config.Mysql.Write, config.Mysql.Read)

	// Redis
	db.InitRedis(config.Redis)

	// 启动服务器
	logger.Infof("Server Started with env: %s, listen: %s", config.Env, config.Listen)
	engine.Run(config.Listen)
}
