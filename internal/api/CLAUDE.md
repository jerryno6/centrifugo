# API Package

Core API handlers and protocol implementations for Centrifugo.

## Purpose
Implements the HTTP and gRPC server APIs for backend integration, including:
- Publishing messages to channels
- Managing channels and subscriptions
- Server-side subscriptions
- Message history management
- Consumer integration

## Files
- `api.go` - Main API interface and methods
- `handler.go` - HTTP request handlers
- `grpc.go` - gRPC service implementation
- `consuming.go` - Message consumer integration
- `metrics.go` - API metrics collection
- `*_gen.go` - Generated code for protocol handling
- `*_test.go` - Comprehensive test suite