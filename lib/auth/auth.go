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

type KeyType string

const (
	KeyTypePrivate KeyType = "private"
	KeyTypePublic  KeyType = "public"
)

type KeyStore interface {
	Set(ctx context.Context, keyId string, keyType KeyType, algorithm string, key []byte, expiredAt time.Time) error
	Get(ctx context.Context, keyId string, keyType KeyType) (algorithm string, key []byte, err error)
}
