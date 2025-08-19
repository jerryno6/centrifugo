# App Package

Main application setup and engine configuration.

## Purpose
Orchestrates the entire Centrifugo server including:
- Engine initialization (Redis/Memory)
- Broker setup (Redis/NATS/Memory)
- Presence manager configuration
- TLS setup
- Logging configuration

## Files
- `cmd.go` - CLI command setup
- `engine.go` - Engine configuration and initialization
- `run.go` - Main server startup logic
- `node.go` - Centrifuge node management
- `tls.go` - TLS certificate handling
- `log.go` - Logging configuration
- `mux.go` - HTTP route multiplexer
- `origin.go` - Origin server configuration
- `proxy.go` - Proxy configuration
- `grpc.go` - gRPC server setup
- `graphite.go` - Graphite metrics integration