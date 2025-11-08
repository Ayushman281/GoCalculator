package middleware

import (
	"log/slog"
	"net/http"
	"time"
)

// Logging returns middleware that logs incoming requests and responses.
func Logging(logger *slog.Logger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()

			// Basic start log
			logger.Info("request started",
				"method", r.Method,
				"path", r.URL.Path,
				"remote", r.RemoteAddr,
			)

			rw := newResponseWriter(w)
			// Call next handler
			next.ServeHTTP(rw, r)

			// If handler never wrote a status, treat as 200
			if rw.status == 0 {
				rw.status = http.StatusOK
			}

			logger.Info("request finished",
				"method", r.Method,
				"path", r.URL.Path,
				"status", rw.status,
				"bytes", rw.size,
				"duration_ms", time.Since(start).Milliseconds(),
			)
		})
	}
}
