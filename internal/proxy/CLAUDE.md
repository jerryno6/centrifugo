# Proxy Package

Proxy handlers for authentication and authorization.

## Purpose
Implements proxy-based authentication and authorization by delegating decisions to external HTTP/gRPC services.

## Supported Operations
- **Connect** - Client connection authentication
- **Subscribe** - Channel subscription authorization
- **Publish** - Message publishing authorization
- **RPC** - Remote procedure call handling
- **Refresh** - Token refresh operations
- **Sub Refresh** - Subscription refresh operations

## Files
- `connect_*.go` - Connection proxy handlers
- `subscribe_*.go` - Subscription proxy handlers
- `publish_*.go` - Publishing proxy handlers
- `rpc_*.go` - RPC proxy handlers
- `refresh_*.go` - Token refresh proxy handlers
- `client.go` - HTTP/gRPC proxy client
- `proxy.go` - Core proxy configuration
- `transform.go` - Request/response transformation