package handler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	tests := []struct {
		name           string
		method         string
		queryParam     string
		expectedCode   int
		expectedMsg    string
		expectedErr    string
		isSuccess      bool
	}{
		// Valid names (A-M)
		{"valid name starting with A", http.MethodGet, "Alice", http.StatusOK, "Hello Alice", "", true},
		{"valid name starting with B", http.MethodGet, "Bob", http.StatusOK, "Hello Bob", "", true},
		{"valid name starting with M", http.MethodGet, "Mike", http.StatusOK, "Hello Mike", "", true},
		{"valid lowercase name", http.MethodGet, "alice", http.StatusOK, "Hello alice", "", true},
		{"valid name exactly 100 chars", http.MethodGet, strings.Repeat("A", 100), http.StatusOK, "Hello " + strings.Repeat("A", 100), "", true},

		// Invalid names (N-Z)
		{"invalid name starting with N", http.MethodGet, "Nancy", http.StatusBadRequest, "", "Invalid Input", false},
		{"invalid name starting with Z", http.MethodGet, "Zack", http.StatusBadRequest, "", "Invalid Input", false},
		{"invalid name starting with O", http.MethodGet, "Oscar", http.StatusBadRequest, "", "Invalid Input", false},

		// Empty and invalid input
		{"empty name", http.MethodGet, "", http.StatusBadRequest, "", "Invalid Input", false},
		{"name starting with number", http.MethodGet, "123John", http.StatusBadRequest, "", "Invalid Input", false},
		{"name starting with special char", http.MethodGet, "!John", http.StatusBadRequest, "", "Invalid Input", false},
		{"name starting with space", http.MethodGet, "%20John", http.StatusBadRequest, "", "Invalid Input", false},
		{"name with unicode accent Á", http.MethodGet, "Álvaro", http.StatusBadRequest, "", "Invalid Input", false},
		{"name with unicode accent É", http.MethodGet, "Émilie", http.StatusBadRequest, "", "Invalid Input", false},

		// HTTP method validation
		{"POST method not allowed", http.MethodPost, "Alice", http.StatusMethodNotAllowed, "", "Method not allowed", false},
		{"PUT method not allowed", http.MethodPut, "Alice", http.StatusMethodNotAllowed, "", "Method not allowed", false},
		{"DELETE method not allowed", http.MethodDelete, "Alice", http.StatusMethodNotAllowed, "", "Method not allowed", false},
		{"PATCH method not allowed", http.MethodPatch, "Alice", http.StatusMethodNotAllowed, "", "Method not allowed", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(tt.method, "/hello-world?name="+tt.queryParam, nil)
			w := httptest.NewRecorder()

			HelloWorld(w, req)

			if w.Code != tt.expectedCode {
				t.Errorf("expected status %d, got %d", tt.expectedCode, w.Code)
			}

			if tt.isSuccess {
				var response SuccessResponse
				if err := json.NewDecoder(w.Body).Decode(&response); err != nil {
					t.Fatalf("failed to decode success response: %v", err)
				}

				if response.Message != tt.expectedMsg {
					t.Errorf("expected message %q, got %q", tt.expectedMsg, response.Message)
				}
			} else {
				var response ErrorResponse
				if err := json.NewDecoder(w.Body).Decode(&response); err != nil {
					t.Fatalf("failed to decode error response: %v", err)
				}

				if response.Error != tt.expectedErr {
					t.Errorf("expected error %q, got %q", tt.expectedErr, response.Error)
				}
			}
		})
	}
}
