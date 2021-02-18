package main

import (
	"os"
	"server/cmd/app/router"
	"server/pkg/auth"
	. "server/pkg/config"
	"server/pkg/logger"

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
	engine.Use(logger.GinLogger())

	// 自定义Recovery
	engine.Use(logger.GinRecovery())

	// 路由配置
	router.InitRoutes(engine)

	// 启动服务器
	logger.Infof("Server Started with env: %s, listen: %s", GlobalConfig.Env, GlobalConfig.Listen)
	engine.Run(GlobalConfig.Listen)
}
