# JWTVerify Package

JWT token verification and validation for Centrifugo.

## Purpose
Handles JWT token decoding, verification, and validation for client authentication and authorization.

## Files
- `token_verifier.go` - Token verification interface
- `token_verifier_jwt.go` - JWT token verifier implementation
- `token_decoder.go` - Token decoding utilities
- `token_verifier_jwt_test.go` - JWT verification tests

## Features
- JWT token validation with configurable algorithms
- JWKS (JSON Web Key Set) support for key rotation
- Token claims extraction and validation
- Custom claim verification
- Token expiration handling