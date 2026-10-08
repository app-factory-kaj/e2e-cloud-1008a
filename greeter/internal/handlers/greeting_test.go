package handlers_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"greeter/internal/gen"
	"greeter/internal/handlers"
)

func newTestHandler() http.Handler {
	srv := handlers.NewServer()
	return gen.Handler(gen.NewStrictHandler(srv, nil))
}

func TestGetGreetingWithName(t *testing.T) {
	h := newTestHandler()
	req := httptest.NewRequest(http.MethodGet, "/hello?name=Alice", nil)
	rec := httptest.NewRecorder()

	h.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", rec.Code)
	}

	var body map[string]any
	if err := json.Unmarshal(rec.Body.Bytes(), &body); err != nil {
		t.Fatalf("invalid JSON body: %v", err)
	}
	if len(body) != 1 {
		t.Fatalf("expected exactly one field, got %v", body)
	}
	if body["message"] != "Hello, Alice!" {
		t.Fatalf("expected 'Hello, Alice!', got %v", body["message"])
	}
}

func TestGetGreetingWithoutName(t *testing.T) {
	h := newTestHandler()
	req := httptest.NewRequest(http.MethodGet, "/hello", nil)
	rec := httptest.NewRecorder()

	h.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", rec.Code)
	}

	var body map[string]any
	if err := json.Unmarshal(rec.Body.Bytes(), &body); err != nil {
		t.Fatalf("invalid JSON body: %v", err)
	}
	if len(body) != 1 {
		t.Fatalf("expected exactly one field, got %v", body)
	}
	if body["message"] != "Hello, World!" {
		t.Fatalf("expected 'Hello, World!', got %v", body["message"])
	}
}
