# Middleware Package

HTTP middleware components for Centrifugo.

## Purpose
Provides reusable HTTP middleware for authentication, logging, CORS, rate limiting, and instrumentation.

## Files
- `auth.go` - Authentication middleware
- `cors.go` - CORS (Cross-Origin Resource Sharing) middleware
- `headers.go` - HTTP header manipulation middleware
- `log.go` - Request/response logging middleware
- `connlimit.go` - Connection limiting middleware
- `method.go` - HTTP method restriction middleware
- `otel.go` - OpenTelemetry instrumentation middleware
- `http_instrumentation.go` - HTTP metrics collection
- `user_header_auth.go` - User header-based authentication

## Features
- Pluggable middleware chain
- HTTP metrics collection
- Security headers
- Rate limiting
- OpenTelemetry tracing integration