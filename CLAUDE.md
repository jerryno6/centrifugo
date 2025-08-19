# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Overview
Centrifugo is an open-source scalable real-time messaging server that delivers messages to application users WebSocket, 
TTP-streaming, Server-Sent Events, GRPC, and WebTransport. It's a language-agnostic PUB/SUB server with channel scriptions.

## Architecture
- **Core Engine**: Uses centrifuge library for real-time messaging
- **Transport Layer**: Supports multiple protocols (WebSocket, SSE, HTTP-streaming, GRPC, WebTransport)
- **Storage Backend**: Memory (single-node) or Redis (distributed/clustering)
- **Broker**: Memory, Redis, NATS, or Redis+NATS hybrid
- **Presence Manager**: Memory or Redis for presence tracking
- **API Layer**: HTTP and gRPC APIs for backend communication
- **Consumers**: PostgreSQL, Kafka, AWS SQS, Google Pub/Sub, NATS JetStream, Azure Service Bus

## Key Components

### Core Packages
- `internal/app/` - Main application setup and engine configuration
- `internal/config/` - Configuration loading and validation using Viper
- `internal/api/` - HTTP/gRPC API handlers and protocol definitions
- `internal/proxy/` - Proxy handlers for authentication and authorization
- `internal/consuming/` - Message queue consumers

### Transport Protocols
- `internal/websocket/` - WebSocket transport implementation
- `internal/unisse/` - Unidirectional SSE transport
- `internal/unihttpstream/` - Unidirectional HTTP streaming
- `internal/uniws/` - Unidirectional WebSocket
- `internal/unigrpc/` - Unidirectional gRPC
- `internal/wt/` - WebTransport (experimental)

### Storage & Broker
- `internal/redisnatsbroker/` - Redis+NATS hybrid broker
- `internal/natsbroker/` - NATS broker implementation
- `internal/redisqueue/` - Redis queue management
- `internal/redisshard/` - Redis sharding utilities

## Development Commands

### Build & Test
```bash
# Build the binary
go build

# Run all tests
go test -count=1 -v ./... -cover -race

# Run integration tests
go test -count=1 -v ./... -cover -race --tags=integration

# Generate code (protobuf, handlers, etc.)
make generate

# Update web UI assets
make web

# Update swagger documentation
make swagger-web
```

### Development Setup
```bash
# Install dependencies
go mod tidy

# Run with default config
./centrifugo

# Run with custom config
./centrifugo --config config.json

# Development server for static files
./centrifugo serve --dir ./ --port 3000
```

### Configuration
- Uses JSON/TOML/YAML config files
- Environment variable support via `envconfig` tags
- Default config: `config.json`
- Configuration structure defined in `internal/config/config.go`

### Code Generation
- Protocol buffer definitions in `internal/apiproto/` and `internal/proxyproto/`
- Generated handlers and encoders in `internal/gen/`
- Run `make generate` to regenerate after proto changes

### Testing Patterns
- Unit tests in `*_test.go` files
- Integration tests use `--tags=integration`
- Test data in `testdata/` directories
- Mock objects and test helpers in `internal/tools/`