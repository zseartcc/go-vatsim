package govatsim

import (
	"time"
)

type general struct {
	version          int
	reload           int
	update           time.Time
	connectedClients int
	uniqueUsers      int
}
