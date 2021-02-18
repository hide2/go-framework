package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	. "server/internal"
	. "server/pkg/db"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

type UserHandler struct {
	BaseHandler
}

// ListUsers action
func (h *UserHandler) ListUsers(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	size, _ := strconv.Atoi(c.DefaultQuery("size", "5"))
	data := User.Paginate(page, size)
	c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "ok", "data": data})
}

// CreateUser action
func (h *UserHandler) CreateUser(c *gin.Context) {
	var params CreateUserParam
	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": 400, "msg": "Bad Request"})
		return
	}
	data := User.Create(&params)
	c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "ok", "data": data})
}

// GetUser action
func (h *UserHandler) GetUser(c *gin.Context) {
	id := c.Param("id")

	// Redis
	v, err := Redis.Get(Ctx, "uid:"+id).Result()
	if err == redis.Nil {
		// uid does not exist
	} else if err != nil {
		// err
	} else {
		var data map[string]interface{}
		json.Unmarshal([]byte(v), &data)
		c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "ok", "data": data})
		return
	}

	// DB
	uid, _ := strconv.Atoi(id)
	data := User.Find(uid)
	j, _ := json.Marshal(data)
	Redis.Set(Ctx, "uid:"+id, j, 0)
	c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "ok", "data": data})
}
