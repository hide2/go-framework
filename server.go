package main

import (
	"os"

	"server/libs/auth"
	. "server/libs/config"
	"server/libs/logger"
	"server/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	// Env
	if GlobalConfig.Env == "local" || GlobalConfig.Env == "dev" || GlobalConfig.Env == "test" {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	// ErrorLog
	ef, _ := os.OpenFile(GlobalConfig.ErrorLog, os.O_CREATE|os.O_RDWR|os.O_APPEND, os.ModeAppend|os.ModePerm)
	gin.DefaultErrorWriter = ef

	// New Engine
	engine := gin.New()

	// Middleware: Auth
	engine.Use(auth.Auth())

	// 自定义Logger
	engine.Use(logger.GinLogger(logger.Logger), logger.GinRecovery(logger.Logger, true))

	// 异常崩溃处理
	engine.Use(gin.Recovery())

	// 路由配置
	routes.InitRoutes(engine)

	// 启动服务器
	logger.Infof("Server Started with env: %s, listen: %s", GlobalConfig.Env, GlobalConfig.Listen)
	engine.Run(GlobalConfig.Listen)
}
