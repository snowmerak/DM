package auth

import (
	"context"
	"crypto"
	"time"
)

type Signer interface {
	Sign(key crypto.PrivateKey, payload []byte, expiredAt time.Time) ([]byte, error)
}

type Verifier interface {
	Verify(key crypto.PublicKey, signature []byte, payload []byte) error
}

type KeyStore interface {
	Set(ctx context.Context, keyId string, key []byte, expiredAt time.Time) error
	Get(ctx context.Context, keyId string) ([]byte, error)
}
