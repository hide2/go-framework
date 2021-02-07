package auth

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// before
		if c.Request.URL.Path == "/auth" {
			fmt.Println("Auth:", c.Request.Header)
			c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "ok"})
		}

		c.Next()

		// after
		// nothing
	}
}
