# ConfigTypes Package

Configuration type definitions and utilities for Centrifugo.

## Purpose
Defines all configuration structures and custom types used throughout Centrifugo.

## Files
- `types.go` - Core configuration type definitions
- `duration.go` - Custom duration type with JSON support
- `engine.go` - Engine configuration types
- `namespace.go` - Channel namespace configuration
- `rpc_namespace.go` - RPC namespace configuration
- `pem.go` - PEM certificate handling
- `redis.go` - Redis configuration types
- `tls.go` - TLS configuration types
- `stringmap.go` - String map utilities

## Features
- Type-safe configuration structures
- JSON/TOML/YAML serialization support
- Environment variable mapping
- Configuration validation
- Custom type parsers (duration, PEM, TLS)