package push

import "context"

type Push interface {
	ToDevice(ctx context.Context, message any, target ...string) error
	ToGroup(ctx context.Context, message any, group ...string) error
}

type Checker interface {
	Commit(ctx context.Context, messageId string) error
	Check(ctx context.Context, messageId string) (bool, error)
}

type TokenInfo struct {
	Device string
	Token  string
}

type TokenStorage interface {
	Set(ctx context.Context, user string, device string, token string) error
	Get(ctx context.Context, user string, device string) (string, error)
	GetList(ctx context.Context, user string) ([]*TokenInfo, error)
}
