# CLI Package

Command-line interface utilities for Centrifugo.

## Purpose
Provides helper commands for configuration, token generation, and development.

## Files
- `serve.go` - Static file server for development
- `gentoken.go` - Generate JWT tokens for testing
- `gensubtoken.go` - Generate subscription tokens
- `checktoken.go` - Validate JWT tokens
- `checksubtoken.go` - Validate subscription tokens
- `checkconfig.go` - Validate configuration files
- `genconfig.go` - Generate configuration templates
- `defaultconfig.go` - Show default configuration
- `defaultenv.go` - Show default environment variables
- `configdoc.go` - Generate configuration documentation
- `version.go` - Display version information

## Usage
```bash
./centrifugo gentoken -u user123
./centrifugo checkconfig -c config.json
./centrifugo serve -d ./web -p 8080
```