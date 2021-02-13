# Clean Go Framework

## Features
- Router
- Controller & Model
- RESTful JSON API
- MySQL/Redis
- Read/Write Splitting
- ConnectionPool
- ORM
- Logger
- Config
- Build
- Middleware: Auth

## Quicktart
```
# init db
mysql -uroot -proot < sql/test.sql

# local
go run server.go

or

# build
./build.sh
./server -c config.json

# test
curl -v http://localhost:8080/health
curl -X POST -H "Content-Type: application/json" -d '{"name": "Andy"}'  -v http://localhost:8080/api/v1/users
curl -X POST -H "Content-Type: application/json" -d '{"name": "Calvin"}'  -v http://localhost:8080/api/v1/users
curl -v http://localhost:8080/api/v1/users?page=1
curl -v http://localhost:8080/api/v1/users/1
```

## Routes
routes/config.go
```
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
```

## Config
config.json
```
{
    "env": "local",
    "listen": ":8080",
    "mysql": {
        "write": "root:root@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local",
        "read": "root:root@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
    },
    "redis": "localhost:6379",
    "accesslog": "/tmp/access.log",
    "errorLog": "/tmp/error.log",
    "errorDingTalk": "https://oapi.dingtalk.com/robot/send?access_token=xxx"
}
```

## Structure
```
controller  // Controller
model       // Model
routes      // Routes
libs        // Libraries
sql         // Database init script
build.sh    // Build script
config.json // Configuration
server.go   // Main file
```