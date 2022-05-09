module github.com/869413421/micro-service/common

go 1.16

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0

require (
	github.com/869413421/micro-service/user v0.0.0-20220505144205-50dcb29fea0e
	github.com/dustin/go-humanize v1.0.0 // indirect
	github.com/gin-gonic/gin v1.7.7
	github.com/json-iterator/go v1.1.9
	github.com/micro/go-micro/v2 v2.9.1
	golang.org/x/crypto v0.0.0-20220427172511-eb4f295cb31f
	gorm.io/driver/mysql v1.3.3
	gorm.io/gorm v1.23.5
)
