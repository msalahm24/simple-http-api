package handler

import (
    "encoding/json"
    "net/http"
    "strings"
)

// Response structures
type SuccessResponse struct {
    Message string `json:"message"`
}

type ErrorResponse struct {
    Error string `json:"error"`
}

// HelloWorld handles the /hello-world endpoint
func HelloWorld(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodGet {
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusMethodNotAllowed)
        json.NewEncoder(w).Encode(ErrorResponse{Error: "Method not allowed"})
        return
    }
    
    name := r.URL.Query().Get("name")

    if name == "" {
        sendErrorResponse(w, http.StatusBadRequest, "Invalid Input")
        return
    }

    runes := []rune(name)
    firstRune := runes[0]
    firstLetter := strings.ToUpper(string(firstRune))
    firstChar := []rune(firstLetter)[0]

    if firstChar < 'A' || firstChar > 'Z' {
        sendErrorResponse(w, http.StatusBadRequest, "Invalid Input")
        return
    }

    if firstChar >= 'A' && firstChar <= 'M' {
        sendSuccessResponse(w, "Hello " + name)
    } else {
        sendErrorResponse(w, http.StatusBadRequest, "Invalid Input")
    }
}


func sendSuccessResponse(w http.ResponseWriter, message string) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(SuccessResponse{Message: message})
}


func sendErrorResponse(w http.ResponseWriter, statusCode int, errorMsg string) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(statusCode)
    json.NewEncoder(w).Encode(ErrorResponse{Error: errorMsg})
}