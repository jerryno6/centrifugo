# Gen Package

Code generation utilities for Centrifugo.

## Purpose
Provides automated code generation for API handlers, protocol buffers, and other boilerplate code.

## Files
- `types.go` - Generation type definitions
- `utils.go` - Generation utilities
- `api/` - API handler generation
  - `gen_handlers_*.go` - Generated handlers for different protocols
  - `gen_request_decoder.go` - Request decoder generation
  - `gen_response_encoder.go` - Response encoder generation
  - `gen_result_encoder.go` - Result encoder generation
  - `generate.sh` - Generation script
  - `main.go` - Main generator

## Usage
Run `./generate.sh` to regenerate all API handlers and protocol code after making changes to proto files or API definitions.