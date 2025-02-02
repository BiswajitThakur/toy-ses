package ses

import (
	"sync"

	"github.com/BiswajitThakur/toy-ses/session"
)

type Client struct {
	options session.Options
	// TODO: add more field
	TotalSuccessCount uint64
	TotalFaildCount   uint64
	mu                sync.Mutex
}
