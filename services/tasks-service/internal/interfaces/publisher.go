package interfaces

import "context"

type Publisher interface {
	PublishMsg(ctx context.Context, msg, queueName string) error
}
