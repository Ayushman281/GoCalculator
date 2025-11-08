package handlers

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/Ayushman281/calculator-api/internal/models"
	"github.com/Ayushman281/calculator-api/internal/respond"
)

func SubIntHandler(logger *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			respond.WriteError(logger, w, http.StatusMethodNotAllowed, "only POST allowed")
			return
		}

		var req models.SubtractRequestInt
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			respond.WriteError(logger, w, http.StatusBadRequest, "invalid request payload")
			return
		}
		if req.A == nil || req.B == nil {
			respond.WriteError(logger, w, http.StatusBadRequest, "a and b are required and must be integers")
			return
		}

		result := *req.A - *req.B
		if err := respond.WriteJSON(w, http.StatusOK, map[string]int{"result": result}); err != nil {
			respond.WriteError(logger, w, http.StatusInternalServerError, "failed to write response")
		}
	}
}

func SubFloatHandler(logger *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			respond.WriteError(logger, w, http.StatusMethodNotAllowed, "only POST allowed")
			return
		}
		var req models.SubtractRequestFloat
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			respond.WriteError(logger, w, http.StatusBadRequest, "invalid request payload")
			return
		}
		if req.A == nil || req.B == nil {
			respond.WriteError(logger, w, http.StatusBadRequest, "a and b are required and must be numbers")
			return
		}
		result := *req.A - *req.B
		if err := respond.WriteJSON(w, http.StatusOK, map[string]float64{"result": result}); err != nil {
			respond.WriteError(logger, w, http.StatusInternalServerError, "failed to write response")
		}
	}
}
