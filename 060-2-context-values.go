package main

import (
	"context"
	"net/http"
)

/*
 * Context values are mostly used as a storage for lifecycle-related data (like an HTTP request)
 */

// the reason for having a custom, unexported type for key is collision prevention
// even if other package uses "key" string, it won't collide with the custom typed key
type contextKey string

const correlationIdKey contextKey = "correlationId"

func addCorrelationIdMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), correlationIdKey, "uniqueCorrelationId")
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
