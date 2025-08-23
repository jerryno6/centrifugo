package proxy

import (
	"context"

	"github.com/centrifugal/centrifugo/v6/internal/proxyproto"
)

// MessageProxy allows to send Message requests.
type MessageProxy interface {
	ProxyMessage(context.Context, *proxyproto.MessageRequest) (*proxyproto.MessageResponse, error)
	// Protocol for metrics and logging.
	Protocol() string
	// UseBase64 for bytes in requests from Centrifugo to application backend.
	UseBase64() bool
	// IncludeMeta ...
	IncludeMeta() bool
}
