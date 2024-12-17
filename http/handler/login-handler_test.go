package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestLoginHandler_Success(t *testing.T) {
	reqBody, _ := json.Marshal(map[string]string{
		"username": "eduardo",
		"password": "123456",
	})

	req, err := http.NewRequest("POST", "/api/auth", bytes.NewBuffer(reqBody))
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(LoginHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, status)
	}

	var response map[string]interface{}
	if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
		t.Errorf("Failed to parse response body: %v", err)
	}

	if response["token"] == "" {
		t.Error("Expected token in response, got empty")
	}
}

func TestLoginHandler_InvalidCredentials(t *testing.T) {
	reqBody, _ := json.Marshal(map[string]string{
		"username": "invalid",
		"password": "wrongpassword",
	})

	req, err := http.NewRequest("POST", "/api/auth", bytes.NewBuffer(reqBody))
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(LoginHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("Expected status code %d, got %d", http.StatusBadRequest, status)
	}
}