package pusher

import (
	"context"
	"fmt"
	"time"

	"github.com/snowmerak/DM/lib/broker"
	"github.com/snowmerak/DM/lib/manager"
	"github.com/snowmerak/DM/lib/push"
)

type Pusher struct {
	broker       broker.Broker
	tokenStorage push.TokenStorage
	checker      push.Checker
	pusher       push.Push
	manager      manager.Manager
}

func New(b broker.Broker, ts push.TokenStorage, c push.Checker, p push.Push) *Pusher {
	return &Pusher{
		broker:       b,
		tokenStorage: ts,
		checker:      c,
		pusher:       p,
	}
}

func (p *Pusher) Push(ctx context.Context, namespace string, groupId string, sender string, message []byte, publishedAt time.Time) error {
	ok, err := p.checker.Check(ctx, publishedAt.Format(time.RFC3339Nano))
	if err != nil {
		return fmt.Errorf("failed to check token: %w", err)
	}
	if !ok {
		return push.InvalidTokenErr
	}

	users, err := p.manager.GetGroupMembers(namespace, groupId)
	if err != nil {
		return fmt.Errorf("failed to get group members: %w", err)
	}

	for _, user := range users {
		tokens, err := p.tokenStorage.GetList(ctx, user)
		if err != nil {
			return fmt.Errorf("failed to get tokens: %w", err)
		}

		pushTokens := make([]string, 0, len(tokens))
		for _, token := range tokens {
			pushTokens = append(pushTokens, token.Token)
		}

		if err := p.pusher.ToDevice(ctx, map[string]string{
			"Sender":  sender,
			"Message": string(message),
		}, pushTokens...); err != nil {
			return fmt.Errorf("failed to push message: %w", err)
		}
	}

	if err := p.checker.Commit(ctx, publishedAt.Format(time.RFC3339Nano)); err != nil {
		return fmt.Errorf("failed to commit token: %w", err)
	}

	return nil
}

func (p *Pusher) SubscribeAndPush(ctx context.Context, namespace string) error {
	if err := p.broker.Subscribe(ctx, namespace, broker.AllGroups, func(ctx context.Context, namespace string, groupId string, sender string, message []byte, publishedAt time.Time) error {
		if err := p.Push(ctx, namespace, groupId, sender, message, publishedAt); err != nil {
			return fmt.Errorf("failed to push message: %w", err)
		}

		return nil
	}); err != nil {
		return fmt.Errorf("failed to subscribe to broker: %w", err)
	}

	return nil
}
