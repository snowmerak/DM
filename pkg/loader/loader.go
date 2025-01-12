package loader

import (
	"context"
	"fmt"
	"time"

	"github.com/snowmerak/DM/lib/broker"
	"github.com/snowmerak/DM/lib/message"
)

// Loader is a service that receives messages from a broker and stores them in a storage.
type Loader struct {
	broker  broker.Broker
	storage message.Storage
}

func New(b broker.Broker, s message.Storage) *Loader {
	return &Loader{
		broker:  b,
		storage: s,
	}
}

func (l *Loader) SubscribeAndSave(ctx context.Context, namespace string) error {
	if err := l.broker.Subscribe(ctx, namespace, broker.AllGroups, func(ctx context.Context, namespace string, groupId string, sender string, message []byte, publishedAt time.Time) error {
		if err := l.storage.Save(ctx, namespace, groupId, sender, message, publishedAt); err != nil {
			return fmt.Errorf("failed to store message: %w", err)
		}

		return nil
	}); err != nil {
		return fmt.Errorf("failed to subscribe to broker: %w", err)
	}

	return nil
}
