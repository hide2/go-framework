package routes

import (
	"net/http"
	"server/controller"

	"github.com/gin-gonic/gin"
)

func InitRoutes(router *gin.Engine) {

	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "ok"})
	})
	router.GET("/error", func(c *gin.Context) {
		panic("Error")
	})

	api := router.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			v1.GET("/hello", controller.Hello)

			v1.GET("/users", controller.ListUsers)
			v1.POST("/users", controller.CreateUser)
			v1.GET("/users/:id", controller.GetUser)
		}

	}
}
