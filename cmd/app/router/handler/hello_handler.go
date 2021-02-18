package handler

import (
	"net/http"
	"server/pkg/logger"

	"github.com/gin-gonic/gin"
)

type HelloHandler struct {
	BaseHandler
}

// Hello action
func (h *HelloHandler) Hello(c *gin.Context) {
	logger.Debug("Hello")
	logger.Infof("Hello: %d %s", 123, "name")
	logger.Warn("Hello", 123, []int{1, 2, 3})
	logger.Error("Hello")

	c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "Hello"})
}
