package handler

import (
	"net/http"
)

// Health handles the /health endpoint
func Health(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		sendErrorResponse(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	sendSuccessResponse(w, "OK")
}
