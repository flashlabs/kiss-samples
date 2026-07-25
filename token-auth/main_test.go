package main_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/flashlabs/kiss-samples/token-auth"
)

func TestAuthenticateServiceToken(t *testing.T) {
	cache := main.NewTokenCache()
	cache.Register("valid-token-xyz", "payment-service")

	handler := main.AuthenticateServiceToken(cache, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("ok"))
	}))

	t.Run("Valid Token", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/api/v1/internal/data", nil)
		req.Header.Set("X-Service-Token", "valid-token-xyz")
		rec := httptest.NewRecorder()

		handler.ServeHTTP(rec, req)

		if rec.Code != http.StatusOK {
			t.Errorf("expected status 200, got %d", rec.Code)
		}
	})

	t.Run("Invalid Token", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/api/v1/internal/data", nil)
		req.Header.Set("X-Service-Token", "wrong-token")
		rec := httptest.NewRecorder()

		handler.ServeHTTP(rec, req)

		if rec.Code != http.StatusUnauthorized {
			t.Errorf("expected status 401, got %d", rec.Code)
		}
	})

	t.Run("Missing Token Header", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/api/v1/internal/data", nil)
		rec := httptest.NewRecorder()

		handler.ServeHTTP(rec, req)

		if rec.Code != http.StatusUnauthorized {
			t.Errorf("expected status 401, got %d", rec.Code)
		}
	})
}

func BenchmarkAuthenticateServiceToken(b *testing.B) {
	cache := main.NewTokenCache()
	cache.Register("valid-token-xyz", "payment-service")

	handler := main.AuthenticateServiceToken(cache, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	req := httptest.NewRequest(http.MethodGet, "/api/v1/internal/data", nil)
	req.Header.Set("X-Service-Token", "valid-token-xyz")

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		rec := httptest.NewRecorder()
		handler.ServeHTTP(rec, req)
	}
}
