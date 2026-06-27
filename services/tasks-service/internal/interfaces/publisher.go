package interfaces

import "context"

type Publisher interface {
	PublishMsg(ctx context.Context, userId int, path, filetype, queueName string) error
}
