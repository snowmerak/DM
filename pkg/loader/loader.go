package loader

import (
	"github.com/snowmerak/DM/lib/broker"
	"github.com/snowmerak/DM/lib/message"
)

type Loader struct {
	broker  broker.Broker
	storage message.Storage
}
