package server

import (
	"net/http"

	"github.com/msalahm24/simple-http-api/handler"
)

// setupRoutes configures all the routes
func (s *Server) setupRoutes() {
	http.HandleFunc("/hello-world", RecoveryMiddleware(LoggingMiddleware(handler.HelloWorld)))
	http.HandleFunc("/health", RecoveryMiddleware(LoggingMiddleware(handler.Health)))
}
