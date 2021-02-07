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