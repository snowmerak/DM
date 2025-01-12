package receiver

import (
	"context"
	"fmt"
	"time"

	"github.com/snowmerak/DM/lib/auth"
	"github.com/snowmerak/DM/lib/auth/encrypt"
	"github.com/snowmerak/DM/lib/broker"
)

type Receiver struct {
	keyStore auth.KeyStore
	verifier auth.Verifier
	broker   broker.Broker
}

func New(verifier auth.Verifier, broker broker.Broker) *Receiver {
	return &Receiver{
		verifier: verifier,
		broker:   broker,
	}
}

func (r *Receiver) CheckToken(ctx context.Context, keyId string, message []byte, signature []byte) error {
	algorithm, parsedPublicKey, err := r.keyStore.Get(ctx, keyId, auth.KeyTypePublic)
	if err != nil {
		return fmt.Errorf("failed to get public key: %w", err)
	}

	publicKey, err := encrypt.UnmarshalPublicKey(algorithm, parsedPublicKey)
	if err != nil {
		return fmt.Errorf("failed to unmarshal public key: %w", err)
	}

	if err := r.verifier.Verify(publicKey, signature, message); err != nil {
		return fmt.Errorf("failed to verify signature: %w", err)
	}

	return nil
}

func (r *Receiver) Receive(ctx context.Context, namespace string, groupId string, sender string, message []byte, publishedAt time.Time) error {
	if err := r.broker.Publish(ctx, namespace, groupId, sender, message, publishedAt); err != nil {
		return fmt.Errorf("failed to publish message: %w", err)
	}

	return nil

}
