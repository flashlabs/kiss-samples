# Token Optimization in Authorization: OAuth vs. API Tokens

Demonstrates sub-millisecond inter-service authentication in Go using pre-hashed API tokens and local memory cache lookup (`Lookup`) as a lightweight alternative to full OAuth 2.0 introspection round-trips.

## Features

- **Sub-millisecond verification**: Local lookup avoiding network RPC calls to Identity Providers.
- **Timing-attack safe**: Tokens are hashed with SHA-256 before cache lookup.
- **Black-box tests & Benchmarks**: Included in `main_test.go` using `package main_test`.

## Running the Server

Start the HTTP server:

```bash
go run main.go
```

The server listens on `:8080` with endpoint `GET /api/v1/internal/data`.

## Testing with `curl`

```bash
# Valid token (returns 200 OK):
curl -i -H "X-Service-Token: secret-internal-token-123" http://localhost:8080/api/v1/internal/data

# Invalid token (returns 401 Unauthorized):
curl -i -H "X-Service-Token: invalid-token" http://localhost:8080/api/v1/internal/data

# Missing token (returns 401 Unauthorized):
curl -i http://localhost:8080/api/v1/internal/data
```

## Running Unit Tests & Benchmarks

```bash
# Run unit tests:
go test -v ./...

# Run benchmarks:
go test -bench=. ./...
```

For details see: https://blog.skopow.ski/stop-using-full-oauth-handshakes-for-internal-microservice-auth
