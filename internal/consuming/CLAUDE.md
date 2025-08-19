# Consuming Package

Message queue consumers for Centrifugo.

## Purpose
Implements consumers for various message queue systems to integrate external data sources.

## Supported Backends
- **PostgreSQL** - Transactional outbox pattern
- **Kafka** - Apache Kafka topics
- **AWS SQS** - Amazon Simple Queue Service
- **Google Pub/Sub** - Google Cloud Pub/Sub
- **NATS JetStream** - NATS streaming
- **Azure Service Bus** - Microsoft Azure messaging
- **Redis Stream** - Redis streams

## Files
- `consuming.go` - Core consumer interface and management
- `kafka.go` - Kafka consumer implementation
- `postgresql.go` - PostgreSQL outbox consumer
- `aws_sqs.go` - AWS SQS consumer
- `google_pub_sub.go` - Google Pub/Sub consumer
- `nats_jetstream.go` - NATS JetStream consumer
- `azure_service_bus.go` - Azure Service Bus consumer
- `redis_stream.go` - Redis stream consumer
- `metrics.go` - Consumer metrics collection
- `backoff.go` - Retry/backoff strategies

## Usage
Configure consumers in the main configuration file to enable message ingestion from external systems.