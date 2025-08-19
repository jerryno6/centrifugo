# WebSocket Package

WebSocket transport implementation for Centrifugo.

## Purpose
Provides WebSocket protocol support for bidirectional real-time communication.

## Files
- `server.go` - WebSocket server implementation
- `client.go` - WebSocket client utilities
- `conn.go` - WebSocket connection management
- `compression.go` - Per-message compression support
- `prepared.go` - Prepared message utilities
- `util.go` - WebSocket helper functions
- `convert_*.go` - Message format conversion utilities
- `mask.go` - Frame masking utilities
- `tls_handshake.go` - TLS connection handling

## Features
- RFC 6455 compliant WebSocket implementation
- Per-message compression (RFC 7692)
- Automatic ping/pong handling
- Connection multiplexing
- Message batching
- Binary and text message support