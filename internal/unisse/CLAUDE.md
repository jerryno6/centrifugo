# UniSSE Package

Unidirectional Server-Sent Events (SSE) transport.

## Purpose
Provides unidirectional server-to-client streaming using Server-Sent Events protocol.

## Files
- `handler.go` - HTTP handler for SSE endpoints
- `transport.go` - Transport interface implementation
- `config.go` - Configuration structure

## Features
- Standard SSE protocol (EventSource)
- Automatic reconnection
- HTTP/1.1 and HTTP/2 support
- Simple integration with browsers
- No client SDK required