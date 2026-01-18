# HTTP Server in Go vs Java: the stuff that actually hurts in production

A production-ready HTTP server in Go demonstrating critical patterns for reliability:

## Features

- **Graceful shutdown** - Handles SIGTERM/SIGINT with configurable timeout
- **Request timeouts** - Read (5s), write (10s), and idle (60s) timeouts prevent resource exhaustion
- **Backpressure** - Semaphore-based concurrency limiting (100 max in-flight requests)
- **Context cancellation** - Respects client disconnections to avoid wasted work
- **Request tracing** - Request ID generation and propagation via headers
- **Structured logging** - Method, path, duration, and request ID tracking
- **Middleware chain** - Composable request processing pipeline

## Running

```bash
go run main.go
```

Server listens on `:8080` with endpoints:
- `GET /health` - Health check
- `GET /api/v1/items/{id}` - Item retrieval with simulated latency

## Testing Backpressure

```bash
# Flood with requests to trigger 429 responses
for i in {1..200}; do curl -s http://localhost:8080/api/v1/items/$i & done
```

For details see: https://blog.skopow.ski/http-server-in-go-vs-java-the-stuff-that-actually-hurts-in-production
