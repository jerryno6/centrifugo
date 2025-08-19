Centrifugo is an open-source scalable real-time messaging server. Centrifugo can instantly deliver messages to application online users connected over supported transports (WebSocket, HTTP-streaming, Server-Sent Events (aka EventSource), GRPC, WebTransport). Centrifugo has the concept of channel subscriptions – so it's a user-facing PUB/SUB server.

Centrifugo is language-agnostic and can be used to build chat apps, live comments, multiplayer games, real-time data visualizations, collaborative tools, etc. in combination with any backend. It is well suited for modern architectures and allows decoupling the business logic from the real-time transport layer.

Several official client SDKs for browser and mobile development wrap the bidirectional protocol. In addition, Centrifugo supports a unidirectional approach for simple use cases with no SDK dependency.

## Documentation

* [Centrifugo official documentation site](https://centrifugal.dev)
* [Installation instructions](https://centrifugal.dev/docs/getting-started/installation)
* [Getting started tutorial](https://centrifugal.dev/docs/getting-started/quickstart)
* [Design overview and idiomatic usage](https://centrifugal.dev/docs/getting-started/design)
* [Build a WebSocket chat/messenger app with Centrifugo](https://centrifugal.dev/docs/tutorial/intro) tutorial
* [Centrifugal blog](https://centrifugal.dev/blog)
* [FAQ](https://centrifugal.dev/docs/faq)

## Join community

* [Telegram](https://t.me/joinchat/ABFVWBE0AhkyyhREoaboXQ)
* [Discord](https://discord.gg/tYgADKx)
* [Twitter](https://twitter.com/centrifugalabs)

## Why Centrifugo

The core idea of Centrifugo is simple – it's a PUB/SUB server on top of modern real-time transports:

<img src="https://centrifugal.dev/img/protocol_pub_sub.png?v=2" />

The hard part is to make this concept production-ready, efficient, flexible and available from different application environments. Centrifugo is a mature solution that already helped many projects with adding real-time features and scale towards many concurrent connections. Centrifugo provides a set of features not available in other open-source solutions in the area:

* Efficient real-time transports: WebSocket, HTTP-streaming, Server-Sent Events, GRPC, WebTransport
* Built-in scalability with Redis (or Redis Cluster, or Redis-compatible storage – ex. AWS Elasticache, Valkey, KeyDB, DragonflyDB, etc), or Nats.
* Simple HTTP and GRPC server API to communicate with Centrifugo from the app backend
* Asynchronous PostgreSQL and Kafka consumers to support transactional outbox and CDC patterns
* Flexible connection authentication mechanisms: JWT and proxy-like (via request from Centrifugo to the backend)
* Channel subscription multiplexing over a single connection
* Different types of subscriptions: client-side and server-side
* Various channel permission strategies, channel namespace concept
* Hot message history in channels, with automatic message recovery upon reconnect, cache recovery mode (deliver latest publication immediately upon subscription)
* Delta compression in channels based on Fossil algorithm
* Online channel presence information, with join/leave notifications
* A way to send RPC calls to the backend over the real-time connection
* Strict and effective client protocol wrapped by several official SDKs
* JSON and binary Protobuf message transfer, with optimized serialization and built-in batching
* Beautiful embedded admin web UI
* Great observability with lots of Prometheus metrics exposed and official Grafana dashboard
* And much more, visit [Centrifugo documentation site](https://centrifugal.dev)

## Backing

This repository is hosted by [packagecloud.io](https://packagecloud.io/).

<a href="https://packagecloud.io/"><img height="46" width="158" alt="Private NPM registry and Maven, RPM, DEB, PyPi and RubyGem Repository · packagecloud" src="https://packagecloud.io/images/packagecloud-badge.png" /></a>

Also thanks to [JetBrains](https://www.jetbrains.com/) for supporting OSS (most of the code here written in Goland):

<a href="https://www.jetbrains.com/"><img height="140" src="https://resources.jetbrains.com/storage/products/company/brand/logos/jb_beam.png" alt="JetBrains logo"></a>

## Local Development

### Prerequisites

- **Go 1.21+** - [Download Go](https://golang.org/dl/)
- **Git** - for version control
- **VSCode** - recommended IDE with Go extension

### Quick Start

1. **Clone the repository**
   ```bash
   git clone https://github.com/centrifugal/centrifugo.git
   cd centrifugo
   ```

2. **Install dependencies**
   ```bash
   go mod tidy
   ```

3. **Build the project**
   ```bash
   go build
   ```

4. **Run with basic configuration**
   ```bash
   ./centrifugo --health.enabled=true --log_level=debug
   ```

### VSCode Setup

#### 1. Install Extensions
- **Go** - Official Go extension from Google
- **GitLens** - Enhanced Git capabilities

#### 2. Create Debug Configuration

Create `.vscode/launch.json` in the project root:

```json
{
    "version": "0.2.0",
    "configurations": [
        {
            "name": "Launch Centrifugo",
            "type": "go",
            "request": "launch",
            "mode": "auto",
            "program": "${workspaceFolder}",
            "args": [
                "--health.enabled=true",
                "--log_level=debug"
            ],
            "env": {},
            "console": "integratedTerminal"
        },
        {
            "name": "Launch Centrifugo with Config",
            "type": "go",
            "request": "launch",
            "mode": "auto",
            "program": "${workspaceFolder}",
            "args": [
                "--config=config.json"
            ],
            "env": {},
            "console": "integratedTerminal"
        },
        {
            "name": "Launch with Custom Port",
            "type": "go",
            "request": "launch",
            "mode": "auto",
            "program": "${workspaceFolder}",
            "args": [
                "--health.enabled=true",
                "--port=8100",
                "--log_level=debug"
            ],
            "env": {},
            "console": "integratedTerminal"
        }
    ]
}
```

#### 3. Create Workspace Settings

Create `.vscode/settings.json`:

```json
{
    "go.useLanguageServer": true,
    "go.formatTool": "gofmt",
    "go.lintTool": "golangci-lint",
    "go.testFlags": ["-v", "-race"],
    "go.buildFlags": ["-race"],
    "files.eol": "\n",
    "files.trimTrailingWhitespace": true,
    "files.insertFinalNewline": true
}
```

### Development Workflow

#### 1. Build and Test

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

#### 2. Run with Different Configurations

**Basic development setup:**
```bash
./centrifugo --health.enabled=true --log_level=debug
```

**With custom config:**
```bash
./centrifugo --config=config.json
```

**With Redis backend:**
```bash
./centrifugo --config=config.json --engine=redis --redis_address=localhost:6379
```

#### 3. Testing Endpoints

**Test health endpoints:**
```bash
curl http://localhost:8000/health
curl http://localhost:8000/health-full-check
```

**Test API endpoints:**
```bash
# Get server info
curl -X POST http://localhost:8000/api/info

# Publish a message (requires API key if not in insecure mode)
curl -X POST http://localhost:8000/api/publish \
  -H "Content-Type: application/json" \
  -d '{"channel": "test", "data": {"message": "hello"}}'
```

### Configuration Examples

#### Basic Development Config

Create `config.json`:
```json
{
  "health": {
    "enabled": true
  },
  "log_level": "debug",
  "debug": {
    "enabled": true
  }
}
```

#### Advanced Development Config

Create `dev-config.yaml`:
```yaml
health:
  enabled: true

log_level: debug

debug:
  enabled: true

prometheus:
  enabled: true

admin:
  enabled: true

http_api:
  insecure: true
```

### Debugging Tips

#### 1. Using VSCode Debugger
- Set breakpoints in Go files
- Use F5 to start debugging
- Use Debug Console for interactive debugging
- Check Variables panel for current state

#### 2. Logging
- Use `--log_level=debug` for verbose logging
- Check console output for connection details
- Use `--log_format=json` for structured logging

#### 3. Monitoring
- Visit `http://localhost:8000/debug/pprof/` for profiling
- Check `http://localhost:8000/metrics` for Prometheus metrics
- Use embedded admin UI at `http://localhost:8000/admin` (when enabled)

#### 4. Common Issues

**Port already in use:**
```bash
lsof -i :8000
./centrifugo --port=8080
```

**Health endpoint 404:**
- Ensure `--health.enabled=true` is used
- Check logs for enabled endpoints

**Connection issues:**
- Check firewall settings
- Verify network connectivity
- Use `--log_level=debug` for detailed logs

### Contributing

1. **Fork the repository** on GitHub
2. **Create a feature branch**:
   ```bash
   git checkout -b feature/your-feature
   ```
3. **Make your changes** and test them
4. **Run tests**:
   ```bash
   go test ./...
   make generate
   ```
5. **Commit and push** your changes
6. **Create a pull request**

### Development Commands Reference

| Command | Description |
|---------|-------------|
| `go build` | Build the binary |
| `go test ./...` | Run all tests |
| `go test -race ./...` | Run tests with race detector |
| `go run . --help` | Show all available flags |
| `make generate` | Generate protobuf and handler code |
| `make web` | Update web UI assets |
| `make swagger-web` | Update swagger documentation |

### VSCode Keyboard Shortcuts

- **F5**: Start debugging
- **Ctrl+F5**: Run without debugging
- **Shift+F5**: Stop debugging
- **Ctrl+Shift+D**: Open debug panel
- **Ctrl+Shift+P**: Command palette (search "Go:" commands)
