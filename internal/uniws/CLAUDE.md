# UniWS Package

Unidirectional WebSocket transport for Centrifugo.

## Purpose
Provides unidirectional WebSocket transport for simple use cases without requiring full bidirectional protocol support.

## Files
- `handler.go` - HTTP handler for unidirectional WebSocket
- `transport.go` - Transport interface implementation
- `config.go` - Configuration structure
- `cancelctx.go` - Context cancellation utilities

## Usage
Unidirectional WebSocket is useful for:
- Simple pub/sub scenarios
- Server-to-client streaming
- Broadcasting without client-to-server messaging
- Legacy client compatibility