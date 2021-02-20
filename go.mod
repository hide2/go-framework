module server

go 1.15

require (
	github.com/gin-contrib/pprof v1.3.0
	github.com/gin-gonic/gin v1.6.3-0.20210208152422-1bdf86b72202
	github.com/go-redis/redis v6.15.9+incompatible
	github.com/go-redis/redis/v8 v8.5.0
	github.com/natefinch/lumberjack v2.0.0+incompatible
	github.com/pkg/errors v0.9.1 // indirect
	go.uber.org/zap v1.16.0
	gopkg.in/natefinch/lumberjack.v2 v2.0.0 // indirect
	gorm.io/driver/mysql v1.0.4
	gorm.io/gorm v1.20.12
	gorm.io/plugin/dbresolver v1.1.0
)
