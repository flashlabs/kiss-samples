package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func FuzzCreateUserHandler(f *testing.F) {
	// Seed corpus with a valid JSON input
	f.Add(`{"username":"testuser","age":30}`)

	f.Fuzz(func(t *testing.T, body string) {
		defer func() {
			if r := recover(); r != nil {
				t.Fatalf("Handler panicked with input %q: %v", body, r)
			}
		}()

		req := httptest.NewRequest(http.MethodPost, "/users", strings.NewReader(body))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()

		CreateUserHandler(w, req)

		resp := w.Result()
		defer func(Body io.ReadCloser) {
			err := Body.Close()
			if err != nil {
				t.Fatalf("Failed to close response body: %v", err)
			}
		}(resp.Body)

		if resp.StatusCode == http.StatusInternalServerError {
			t.Errorf("Unexpected 500 Internal Server Error for input: %q", body)
		}
	})
}
