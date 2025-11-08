package respond

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

// WriteJSON encodes v as JSON with the given status code and writes it to the response.
// It returns an error if encoding or writing fails.
func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

// WriteError sends a structured JSON error message and logs it.
func WriteError(logger *slog.Logger, w http.ResponseWriter, status int, msg string) {
	logger.Error("client error", "status", status, "error", msg)
	_ = WriteJSON(w, status, map[string]string{"error": msg})
}
