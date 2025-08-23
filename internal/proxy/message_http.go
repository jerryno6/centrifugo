package proxy

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/centrifugal/centrifugo/v6/internal/proxyproto"
)

// MessageRequestHTTP ...
type MessageRequestHTTP struct {
	baseRequestHTTP

	UserID  string `json:"user"`
	Channel string `json:"channel"`

	Data json.RawMessage `json:"data,omitempty"`
	// Base64Data to proxy binary data.
	Base64Data string `json:"b64data,omitempty"`
}

// HTTPMessageProxy ...
type HTTPMessageProxy struct {
	config     Config
	httpCaller HTTPCaller
}

var _ MessageProxy = (*HTTPMessageProxy)(nil)

// NewHTTPMessageProxy ...
func NewHTTPMessageProxy(p Config) (*HTTPMessageProxy, error) {
	httpClient, err := proxyHTTPClient(p, "message_proxy")
	if err != nil {
		return nil, fmt.Errorf("error creating HTTP client: %w", err)
	}
	return &HTTPMessageProxy{
		httpCaller: NewHTTPCaller(httpClient),
		config:     p,
	}, nil
}

// ProxyMessage proxies Message to application backend.
func (p *HTTPMessageProxy) ProxyMessage(ctx context.Context, req *proxyproto.MessageRequest) (*proxyproto.MessageResponse, error) {
	data, err := httpEncoder.EncodeMessageRequest(req)
	if err != nil {
		return nil, err
	}
	respData, err := p.httpCaller.CallHTTP(ctx, p.config.Endpoint, httpRequestHeaders(ctx, p.config), data)
	if err != nil {
		return transformMessageResponse(err, p.config.HTTP.StatusToCodeTransforms)
	}
	return httpDecoder.DecodeMessageResponse(respData)
}

// Protocol ...
func (p *HTTPMessageProxy) Protocol() string {
	return "http"
}

// UseBase64 ...
func (p *HTTPMessageProxy) UseBase64() bool {
	return p.config.BinaryEncoding
}

// IncludeMeta ...
func (p *HTTPMessageProxy) IncludeMeta() bool {
	return p.config.IncludeConnectionMeta
}
