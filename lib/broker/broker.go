package broker

import "context"

type Broker interface {
	Publish(ctx context.Context, namespace string, groupId string, message []byte) error
	Subscribe(ctx context.Context, namespace string, groupId string, callback func(ctx context.Context, namespace string, groupId string, message []byte) error) error
}
