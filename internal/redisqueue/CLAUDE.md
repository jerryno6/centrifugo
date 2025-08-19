# RedisQueue Package

Redis-based queue management for Centrifugo.

## Purpose
Provides Redis stream-based queue functionality for message processing and distribution.

## Files
- `consumer.go` - Redis stream consumer implementation
- `producer.go` - Redis stream producer implementation
- `message.go` - Message structure and serialization
- `stream.go` - Redis stream management utilities
- `signals.go` - Signal handling for graceful shutdown
- `x*.go` - Redis stream commands (XADD, XREAD, XACK, etc.)

## Features
- Redis Streams for reliable message delivery
- Consumer groups for load balancing
- Automatic message acknowledgment
- Message retry with backoff
- Stream trimming and cleanup