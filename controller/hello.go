package controller

import (
	"net/http"
	"server/libs/logger"

	"github.com/gin-gonic/gin"
)

// Hello action
func Hello(c *gin.Context) {
	logger.Debug("Hello")
	logger.Infof("Hello: %d %s", 123, "name")
	logger.Warn("Hello", 123, []int{1, 2, 3})
	logger.Error("Hello")

	c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "Hello"})
}
