# API Proto Package

Protocol buffer definitions and generated code for the Centrifugo API.

## Purpose
Defines the gRPC and HTTP API contracts for backend communication.

## Files
- `api.proto` - Protocol buffer definitions
- `api.pb.go` - Generated Go code from proto
- `api_grpc.pb.go` - Generated gRPC service code
- `encode.go` / `decode.go` - Request/response encoding
- `generate.sh` - Code generation script
- `swagger/` - OpenAPI/Swagger documentation

## Regeneration
Run `./generate.sh` to regenerate Go code from proto definitions.