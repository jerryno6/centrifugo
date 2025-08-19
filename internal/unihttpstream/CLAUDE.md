# UniHTTPStream Package

Unidirectional HTTP streaming transport.

## Purpose
Provides unidirectional server-to-client streaming over HTTP long polling.

## Files
- `handler.go` - HTTP handler for streaming endpoints
- `transport.go` - Transport interface implementation
- `config.go` - Configuration structure

## Features
- HTTP long polling fallback
- Works with restrictive proxies/firewalls
- Chunked transfer encoding
- Automatic reconnection handling