# Clean Go Framework

## Features
- Router
- Handler & Model
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
go run cmd/app/server.go

or

# build
./build.sh
./server -c config/config.json

# test
curl -v http://localhost:8080/health
curl -X POST -H "Content-Type: application/json" -d '{"name": "Andy"}'  -v http://localhost:8080/api/v1/users
curl -X POST -H "Content-Type: application/json" -d '{"name": "Calvin"}'  -v http://localhost:8080/api/v1/users
curl -v http://localhost:8080/api/v1/users?page=1
curl -v http://localhost:8080/api/v1/users/1

# benchmark
cd test
./wrk.sh

# profiling
go tool pprof http://localhost:8080/debug/pprof/goroutine\?second\=20
top10

brew install graphviz
go tool pprof -http=:8081 xxx/pprof.goroutine.001.pb.gz
```

## Router
cmd/app/router/router.go
```
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
```

## Config
config/config.json
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
cmd/app/router/handler   // Handler
cmd/app/router/router.go // Router
config/config.json       // Config
internal                 // Biz & Data
pkg                      // Libs
sql                      // Database init script
build.sh                 // Build script
```