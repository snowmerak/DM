package pusher

import (
	"github.com/snowmerak/DM/lib/broker"
	"github.com/snowmerak/DM/lib/push"
)

type Pusher struct {
	broker       broker.Broker
	tokenStorage push.TokenStorage
	checker      push.Checker
	pusher       push.Push
}
