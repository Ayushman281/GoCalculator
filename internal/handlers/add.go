package handlers

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/Ayushman281/calculator-api/internal/models"
	"github.com/Ayushman281/calculator-api/internal/respond"
)

func AddIntHandler(logger *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			respond.WriteError(logger, w, http.StatusMethodNotAllowed, "only POST allowed")
			return
		}

		/**
		🧩 What NewDecoder does

			json.NewDecoder(r.Body)
			→ creates a JSON reader that reads from the HTTP request body.
			It doesn’t yet read anything — it just prepares a decoder connected to the request’s body stream.

		🧩 What Decode does

			Decode(&req)
			→ actually reads the JSON data from the body, parses it, and fills your req struct.
		**/

		var req models.AddRequestInt
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			respond.WriteError(logger, w, http.StatusBadRequest, "invalid JSON body")
			return
		}

		if req.A == nil || req.B == nil {
			respond.WriteError(logger, w, http.StatusBadRequest, "a and b are required and must be integers")
			return
		}

		result := *req.A + *req.B
		if err := respond.WriteJSON(w, http.StatusOK, map[string]int{"result": result}); err != nil {
			respond.WriteError(logger, w, http.StatusInternalServerError, "failed to write response")
		}
	}
}

func AddFloatHandler(logger *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger.Info("request", "path", r.URL.Path, "remote", r.RemoteAddr, "method", r.Method)

		if r.Method != "POST" {
			respond.WriteError(logger, w, http.StatusMethodNotAllowed, "only POST allowed")
			return
		}

		var req models.AddRequestFloat
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			respond.WriteError(logger, w, http.StatusBadRequest, "invalid JSON body")
			return
		}
		if req.A == nil || req.B == nil {
			respond.WriteError(logger, w, http.StatusBadRequest, "a and b are required and must be numbers")
			return
		}

		result := *req.A + *req.B
		if err := respond.WriteJSON(w, http.StatusOK, map[string]float64{"result": result}); err != nil {
			respond.WriteError(logger, w, http.StatusInternalServerError, "failed to write response")
		}
	}
}
