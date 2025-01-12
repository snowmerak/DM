package message

import (
	"context"
	"time"
)

type Message struct {
	GroupId     string
	Sender      string
	Message     []byte
	PublishedAt time.Time
}

type Storage interface {
	Save(ctx context.Context, namespace string, groupId string, sender string, message []byte, publishedAt time.Time) error
	Search(ctx context.Context, namespace string, groupId string, startPublishedAt, endPublishedAt time.Time) ([]*Message, error)
	SearchViaSender(ctx context.Context, namespace string, groupId string, sender string, startPublishedAt, endPublishedAt time.Time) ([]*Message, error)
}

type Indexer interface {
	Create(ctx context.Context, namespace string, groupId string, temporaryNumber int64) error
	Insert(ctx context.Context, namespace string, groupId string, temporaryNumber int64, sender string, message []byte, publishedAt time.Time) error
	DeleteOne(ctx context.Context, namespace string, groupId string, temporaryNumber int64, sender string, publishedAt time.Time) error
	Drop(ctx context.Context, namespace string, groupId string, temporaryNumber int64) error
	DeleteRange(ctx context.Context, namespace string, groupId string, temporaryNumber int64, startTimestamp time.Time, endTimestamp time.Time) error
	Search(ctx context.Context, namespace string, groupId string, temporaryNumber int64, query string) ([]*Message, error)
}
