package main

import (
	"crypto/sha256"
	"context"
	"fmt"
	"log"
	"net/http"
	"sync"
)

type contextKey string

const serviceIDKey contextKey = "serviceID"

// TokenCache provides in-memory lookup of hashed API tokens.
type TokenCache struct {
	mu     sync.RWMutex
	tokens map[[32]byte]string
}

func NewTokenCache() *TokenCache {
	return &TokenCache{
		tokens: make(map[[32]byte]string),
	}
}

// Register stores a raw token string hashed into the lookup map.
func (c *TokenCache) Register(token, serviceID string) {
	hash := sha256.Sum256([]byte(token))
	c.mu.Lock()
	defer c.mu.Unlock()
	c.tokens[hash] = serviceID
}

// Lookup retrieves the service ID for a given token hash (idiomatic Go getter: no Get prefix).
func (c *TokenCache) Lookup(hash [32]byte) (string, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	serviceID, ok := c.tokens[hash]
	return serviceID, ok
}

// AuthenticateServiceToken is a middleware that verifies internal API tokens without expensive external RPCs.
func AuthenticateServiceToken(cache *TokenCache, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("X-Service-Token")
		if token == "" {
			http.Error(w, "missing service token", http.StatusUnauthorized)
			return
		}

		// Hash token to prevent timing attacks and lookup in local memory/cache
		tokenHash := sha256.Sum256([]byte(token))
		serviceID, ok := cache.Lookup(tokenHash)
		if !ok {
			http.Error(w, "invalid service token", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), serviceIDKey, serviceID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func main() {
	cache := NewTokenCache()
	cache.Register("secret-internal-token-123", "inventory-service")

	protected := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		serviceID, _ := r.Context().Value(serviceIDKey).(string)
		_, _ = fmt.Fprintf(w, "Authenticated request from service: %s\n", serviceID)
	})

	mux := http.NewServeMux()
	mux.Handle("/api/v1/internal/data", AuthenticateServiceToken(cache, protected))

	log.Println("Server running on :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatalf("http.ListenAndServe: %v", err)
	}
}
