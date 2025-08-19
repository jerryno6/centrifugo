# Config Package

Configuration management for Centrifugo.

## Purpose
Handles loading, validation, and management of configuration from files and environment variables.

## Files
- `config.go` - Main configuration structure
- `container.go` - Configuration container with validation
- `cache.go` - Configuration caching utilities
- `validate.go` - Configuration validation logic
- `migrations.go` - Configuration version migrations
- `testdata/` - Example configuration files

## Configuration Sources
- JSON, TOML, YAML files
- Environment variables
- CLI flags
- Configuration validation with defaults

## Usage
```go
// Load configuration
cfg, err := config.LoadFromFile("config.json")

// Validate configuration
container := config.NewContainer(cfg)
err = container.Validate()
```