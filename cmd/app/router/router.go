package router

import (
	"net/http"
	"server/cmd/app/router/handler"

	"github.com/gin-gonic/gin"
)

func InitRoutes(engine *gin.Engine) {

	engine.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "ok"})
	})

	engine.GET("/error", func(c *gin.Context) {
		panic("Error")
	})

	api := engine.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			v1.GET("/hello", new(handler.HelloHandler).Hello)

			v1.GET("/users", new(handler.UserHandler).ListUsers)
			v1.POST("/users", new(handler.UserHandler).CreateUser)
			v1.GET("/users/:id", new(handler.UserHandler).GetUser)
		}
	}

}
