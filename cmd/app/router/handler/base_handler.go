package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type BaseHandler struct{}

func (bh *BaseHandler) responseJSON(c *gin.Context, code int, msg string, data interface{}) {
	if code == 200 {
		c.JSON(http.StatusOK, gin.H{"code": code, "msg": "OK", "data": data})
	} else {
		c.JSON(code, gin.H{"code": code, "msg": msg, "data": data})
	}
}
