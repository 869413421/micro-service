
GOPATH:=$(shell go env GOPATH)

.PHONY: build
build:
	 CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-w' -i -o micro-user-api ./main.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	docker build . -t micro-user-api:latest
