package uuid

import (
	"github.com/rs/xid"
)

const IDLen = 20

// New Generates a new unique id
func New() string {
	id := xid.New()
	return id.String()
}
