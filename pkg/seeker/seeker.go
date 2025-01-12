package seeker

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/snowmerak/DM/lib/message"
)

type Seeker struct {
	storage message.Storage
	indexer message.Indexer
}

func New(storage message.Storage, indexer message.Indexer) *Seeker {
	return &Seeker{
		storage: storage,
		indexer: indexer,
	}
}

const (
	MaxTriedCount = 5
)

func (s *Seeker) search(ctx context.Context, namespace string, groupId string, sender string, query string, startTime time.Time, endTime time.Time) ([]*message.Message, error) {
	temporaryNumber := rand.Int63()
	triedCount := 0
	err := s.indexer.Create(ctx, namespace, groupId, temporaryNumber)
	for ; err != nil || triedCount < MaxTriedCount; err, triedCount, temporaryNumber = s.indexer.Create(ctx, namespace, groupId, temporaryNumber), triedCount+1, rand.Int63() {
	}
	if triedCount == MaxTriedCount {
		return nil, fmt.Errorf("failed to create temporary index: %w", err)
	}
	defer func() {
		if err := s.indexer.Drop(ctx, namespace, groupId, temporaryNumber); err != nil {
			fmt.Printf("failed to drop temporary index: %v\n", err)
		}
	}()

	rawMessages := make([]*message.Message, 0)
	switch len(sender) {
	case 0:
		rawMessages, err = s.storage.Search(ctx, namespace, groupId, startTime, endTime)
		if err != nil {
			return nil, fmt.Errorf("failed to search messages: %w", err)
		}
	default:
		rawMessages, err = s.storage.SearchViaSender(ctx, namespace, groupId, sender, startTime, endTime)
		if err != nil {
			return nil, fmt.Errorf("failed to search messages: %w", err)
		}
	}

	for _, rawMessage := range rawMessages {
		if err := s.indexer.Insert(ctx, namespace, groupId, temporaryNumber, rawMessage.Sender, rawMessage.Message, rawMessage.PublishedAt); err != nil {
			return nil, fmt.Errorf("failed to insert message into index: %w", err)
		}
	}

	messages, err := s.indexer.Search(ctx, namespace, groupId, temporaryNumber, query)
	if err != nil {
		return nil, fmt.Errorf("failed to search messages: %w", err)
	}

	return messages, nil
}

func (s *Seeker) Search(ctx context.Context, namespace string, groupId string, query string, startTime time.Time, endTime time.Time) ([]*message.Message, error) {
	return s.search(ctx, namespace, groupId, "", query, startTime, endTime)
}

func (s *Seeker) SearchViaSender(ctx context.Context, namespace string, groupId string, sender string, query string, startTime time.Time, endTime time.Time) ([]*message.Message, error) {
	return s.search(ctx, namespace, groupId, sender, query, startTime, endTime)
}
