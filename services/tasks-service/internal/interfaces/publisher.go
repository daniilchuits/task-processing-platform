package interfaces

import "context"

type Publisher interface {
	PublishMsg(ctx context.Context, path, filetype, queueName string) error
}
