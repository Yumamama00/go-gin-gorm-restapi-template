package idutil

import (
	"github.com/google/uuid"
	xid "github.com/rs/xid"
)

// XID XIDの生成
func XID() string {
	return xid.New().String()
}

// UUID UUIDの生成
func UUID() string {
	return uuid.New().String()
}
