package seeker

import "github.com/snowmerak/DM/lib/message"

type Seeker struct {
	stroage message.Storage
	indexer message.Indexer
}
