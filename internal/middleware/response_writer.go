/**
In Go’s standard net/http package, the http.ResponseWriter interface lets you send a response to the client —
but it does not expose information like:

The HTTP status code you sent (200, 400, 500, etc.)
The number of bytes written in the response body

That’s a problem for things like:

Logging middleware (you want to log status, bytes, duration)
Metrics collection (for Prometheus or similar)
Rate limiting / auditing

So, the responseWriter struct acts as a wrapper around the standard http.ResponseWriter to record those details as they happen.
**/

package middleware

import "net/http"

// responseWriter wraps http.ResponseWriter to capture status and size.
type responseWriter struct {
	http.ResponseWriter
	status int
	size   int
}

func (rw *responseWriter) WriteHeader(status int) {
	rw.status = status
	rw.ResponseWriter.WriteHeader(status)
}

//Overriding the Write() method:
func (rw *responseWriter) Write(b []byte) (int, error) {
	// If WriteHeader not called explicitly, default to 200
	if rw.status == 0 {
		rw.status = http.StatusOK
	}
	n, err := rw.ResponseWriter.Write(b)
	rw.size += n
	return n, err
}

// newResponseWriter convenience
func newResponseWriter(w http.ResponseWriter) *responseWriter {
	return &responseWriter{ResponseWriter: w}
}
