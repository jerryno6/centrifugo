# Publishing Package

Message queue publisher for Centrifugo.

## Purpose
Implements publishers for various message queue systems to integrate external data sources.

## Supported Backends
- **Kafka** - Apache Kafka topics

Will support later
- **PostgreSQL** - Transactional outbox pattern
- **AWS SQS** - Amazon Simple Queue Service
- **Google Pub/Sub** - Google Cloud Pub/Sub
- **NATS JetStream** - NATS streaming
- **Azure Service Bus** - Microsoft Azure messaging
- **Redis Stream** - Redis streams

## Files
- `kafka.go` - Kafka publisher implementation

## Usage
Configure publishers in the main configuration file to enable message publishing to external systems.
