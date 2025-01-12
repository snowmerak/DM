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
	Save(ctx context.Context, groupId string, sender string, message []byte, publishedAt time.Time) error
	Search(ctx context.Context, groupId string, startPublishedAt, endPublishedAt time.Time) ([]*Message, error)
	SearchViaSender(ctx context.Context, groupdId string, sender string, startPublishedAt, endPublishedAt time.Time) ([]*Message, error)
}

type Indexer interface {
	Create(ctx context.Context, groupId string) error
	Insert(ctx context.Context, groupdId string, sender string, message []byte, publishedAt time.Time) error
	DeleteOne(ctx context.Context, groupdId string, sender string, pblishedAt time.Time) error
	Drop(ctx context.Context, groupdId string) error
	DeleteRange(ctx context.Context, groupdId string, startTimestamp time.Time, endTimestamp time.Time) error
	Search(ctx context.Context, groupId string, sender string, query string) ([]*Message, error)
}
