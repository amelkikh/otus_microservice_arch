package main

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHealthHandler(t *testing.T) {
	req, err := http.NewRequestWithContext(context.TODO(), http.MethodGet, "/health", nil)
	require.NoError(t, err)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(HealthHandler)
	handler.ServeHTTP(rr, req)
	require.Equal(t, rr.Code, http.StatusOK)
	expected := `{"status":"OK"}`
	require.Equal(t, expected, rr.Body.String())
}
