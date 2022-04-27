package subscriber

import (
	"context"
	log "github.com/micro/go-micro/v2/logger"

	proto "github.com/869413421/micro-service/user/proto/user"
)

type User struct{}

func (e *User) Handle(ctx context.Context, msg *proto.Message) error {
	log.Info("Handler Received message: ", msg.Say)
	return nil
}

func Handler(ctx context.Context, msg *proto.Message) error {
	log.Info("Function Received message: ", msg.Say)
	return nil
}
