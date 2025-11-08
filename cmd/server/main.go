package main

import (
	"log"
	"net/http"

	"github.com/Ayushman281/calculator-api/internal/handlers"
	"github.com/Ayushman281/calculator-api/internal/logger"
	"github.com/Ayushman281/calculator-api/internal/middleware"
)

func main() {
	// create app logger (text handler)
	appLogger := logger.NewLoggerText()

	mux := http.NewServeMux()
	logMw := middleware.Logging(appLogger)

	mux.Handle("/health", logMw(http.HandlerFunc(handlers.HealthHandler(appLogger))))
	mux.Handle("/add/int", logMw(http.HandlerFunc(handlers.AddIntHandler(appLogger))))
	mux.Handle("/add/float", logMw(http.HandlerFunc(handlers.AddFloatHandler(appLogger))))
	mux.Handle("/sub/int", logMw(http.HandlerFunc(handlers.SubIntHandler(appLogger))))
	mux.Handle("/sub/float", logMw(http.HandlerFunc(handlers.SubFloatHandler(appLogger))))
	mux.Handle("/multiply/int", logMw(http.HandlerFunc(handlers.MultiplyIntHandler(appLogger))))
	mux.Handle("/multiply/float", logMw(http.HandlerFunc(handlers.MultiplyFloatHandler(appLogger))))
	mux.Handle("/divide/int", logMw(http.HandlerFunc(handlers.DivideIntHandler(appLogger))))
	mux.Handle("/divide/float", logMw(http.HandlerFunc(handlers.DivideFloatHandler(appLogger))))

	srv := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	appLogger.Info("starting server", "addr", srv.Addr)

	// Run server (no graceful shutdown)
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		// fallback to standard log to avoid depending on logger's behavior during fatal exit
		appLogger.Error("server failed", "err", err)
		log.Fatalf("server failed: %v", err)
	}
}
