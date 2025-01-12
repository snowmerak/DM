package encrypt

import (
	"crypto"
	"errors"
	"fmt"
)

var ErrUnsupportedAlgorithm = errors.New("unsupported algorithm")

var privateKeyMarshaller = map[string]func(key crypto.PrivateKey) ([]byte, error){}
var publicKeyMarshaller = map[string]func(key crypto.PublicKey) ([]byte, error){}

func RegisterPrivateKeyMarshaller(algorithm string, marshaller func(key crypto.PrivateKey) ([]byte, error)) {
	privateKeyMarshaller[algorithm] = marshaller
}

func RegisterPublicKeyMarshaller(algorithm string, marshaller func(key crypto.PublicKey) ([]byte, error)) {
	publicKeyMarshaller[algorithm] = marshaller
}

var privateKeyUnMarshaller = map[string]func(data []byte) (crypto.PrivateKey, error){}
var publicKeyUnMarshaller = map[string]func(data []byte) (crypto.PublicKey, error){}

func RegisterPrivateKeyUnMarshaller(algorithm string, unmarshaller func(data []byte) (crypto.PrivateKey, error)) {
	privateKeyUnMarshaller[algorithm] = unmarshaller
}

func RegisterPublicKeyUnMarshaller(algorithm string, unmarshaller func(data []byte) (crypto.PublicKey, error)) {
	publicKeyUnMarshaller[algorithm] = unmarshaller
}

func MarshalPrivateKey(algorithm string, key crypto.PrivateKey) ([]byte, error) {
	marshaller, ok := privateKeyMarshaller[algorithm]
	if !ok {
		return nil, ErrUnsupportedAlgorithm
	}

	data, err := marshaller(key)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal private key: %w", err)
	}

	return data, nil
}

func MarshalPublicKey(algorithm string, key crypto.PublicKey) ([]byte, error) {
	marshaller, ok := publicKeyMarshaller[algorithm]
	if !ok {
		return nil, ErrUnsupportedAlgorithm
	}

	data, err := marshaller(key)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal public key: %w", err)
	}

	return data, nil
}

func UnmarshalPrivateKey(algorithm string, data []byte) (crypto.PrivateKey, error) {
	unmarshaller, ok := privateKeyUnMarshaller[algorithm]
	if !ok {
		return nil, ErrUnsupportedAlgorithm
	}

	key, err := unmarshaller(data)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal private key: %w", err)
	}

	return key, nil
}

func UnmarshalPublicKey(algorithm string, data []byte) (crypto.PublicKey, error) {
	unmarshaller, ok := publicKeyUnMarshaller[algorithm]
	if !ok {
		return nil, ErrUnsupportedAlgorithm
	}

	key, err := unmarshaller(data)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal public key: %w", err)
	}

	return key, nil
}
