package filters

import (
	"context"
	"net/http"

	"github.com/google/uuid"
)

const correlationIDHeader = "X-Correlation-ID"

type contextKey string

const correlationIDContextKey contextKey = correlationIDHeader

func getCorrelationID(r *http.Request) string {
	if id := r.Header.Get(correlationIDHeader); id != "" {
		if parsed, err := uuid.Parse(id); err == nil {
			return parsed.String()
		}
	}
	return uuid.New().String()
}

func CorrelationIDMiddleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		correlationID := getCorrelationID(r)
		// Set the correlation ID in the response header
		w.Header().Set(correlationIDHeader, correlationID)

		// Optionally, you can also set it in the request context for downstream handlers
		ctx := r.Context()
		ctx = context.WithValue(ctx, correlationIDContextKey, correlationID)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}

func GetCorrelationID(ctx context.Context) string {
	id, _ := ctx.Value(correlationIDContextKey).(string)
	return id
}
