package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestParseRequestDetails(t *testing.T) {
	req := httptest.NewRequest(
		http.MethodGet,
		"http://localhost:8080/",
		nil,
	)

	date := time.Now().Format(time.RFC1123)

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Date", date)
	req.RemoteAddr = "127.0.0.1:12345"

	got := parseRequestDetails(req)
	want := map[string]string{
		"URL":       "http://localhost:8080/",
		"Method":    "GET",
		"Headers":   "application/json",
		"Timestamp": time.Now().Format(time.RFC3339),
		"IPAddress": "127.0.0.1:12345",
	}

	for k, wantVal := range want {
		gotVal, ok := got[k]
		if !ok {
			t.Errorf("Missing key %q in result", k)
			continue
		}
		if gotVal != wantVal {
			t.Errorf("For key %q, got %q, want %q", k, gotVal, wantVal)
		}
	}
}
