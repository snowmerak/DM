package receiver

import (
	"github.com/snowmerak/DM/lib/auth"
	"github.com/snowmerak/DM/lib/broker"
)

type Recevier struct {
	verifier auth.Verifier
	broker   broker.Broker
}
