module github.com/869413421/micro-service/user

go 1.13

// This can be removed once etcd becomes go gettable, version 3.4 and 3.5 is not,
// see https://github.com/etcd-io/etcd/issues/11154 and https://github.com/etcd-io/etcd/issues/11931.
replace google.golang.org/grpc => google.golang.org/grpc v1.26.0
replace github.com/869413421/micro-service/common => ../common

require (
	github.com/869413421/micro-service/common v0.0.0-20220428152058-528eea77a565 // indirect
	github.com/golang/protobuf v1.5.2
	github.com/micro/go-micro/v2 v2.9.1
	google.golang.org/protobuf v1.28.0
)
