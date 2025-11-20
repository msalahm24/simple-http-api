package server

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

// RecoveryMiddleware recovers from panics and returns a 500 error
func RecoveryMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("PANIC: %v", err)
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(map[string]string{
					"error": "Internal server error",
				})
			}
		}()
		next(w, r)
	}
}

// LoggingMiddleware logs details about each HTTP request
func LoggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		log.Printf("[%s] %s %s", r.Method, r.URL.Path, r.URL.RawQuery)

		next(w, r)
		
		duration := time.Since(start)
		log.Printf("[%s] %s - completed in %v", r.Method, r.URL.Path, duration)
	}
}
