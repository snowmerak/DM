package broker

import (
	"context"
	"time"
)

const (
	AllGroups = "*"
)

type Broker interface {
	Publish(ctx context.Context, namespace string, groupId string, sender string, message []byte, publishedAt time.Time) error
	Subscribe(ctx context.Context, namespace string, groupId string, callback func(ctx context.Context, namespace string, groupId string, sender string, message []byte, publishedAt time.Time) error) error
}
