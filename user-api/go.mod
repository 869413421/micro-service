module github.com/869413421/micro-service/user-api

go 1.13

replace github.com/869413421/micro-service/common => ../common

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0

require (
	github.com/869413421/micro-service/common v0.0.0-20220428152058-528eea77a565
	github.com/869413421/micro-service/user v0.0.0-20220505144205-50dcb29fea0e
	github.com/gin-gonic/gin v1.7.7
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/kr/pretty v0.3.0 // indirect
	github.com/micro/go-micro/v2 v2.9.1
	github.com/rogpeppe/go-internal v1.8.0 // indirect
)
