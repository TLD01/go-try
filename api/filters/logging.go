package filters

import (
	"net/http"
	"time"

	"aerowatch.com/api/config/logging"
)

// LoggingMiddleware logs method, path, status code, and elapsed time for every request.
func LoggingMiddleware(next http.Handler) http.Handler {
	logger := logging.GetLogger("http")
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		rw := newResponseWriter(w)

		next.ServeHTTP(rw, r)

		logger.Info("request",
			"method", r.Method,
			"path", r.URL.Path,
			"status", rw.statusCode,
			"duration_ms", time.Since(start).Milliseconds(),
			"correlation_id", GetCorrelationID(r.Context()),
		)

	})
}
