package handler

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestPostLogin(t *testing.T) {
	// Create a new Gin router and set the handler function
	router := gin.Default()
	handler := &EmployeeHandler{}
	router.POST("/login", handler.PostLogin)

	// Test case 1: Successful login
	t.Run("Test successful login", func(t *testing.T) {
		// Create a new HTTP request with valid credentials
		jsonData := `{"user_name": "username", "pass_word": "password"}`
		req, _ := http.NewRequest("POST", "/login", strings.NewReader(jsonData))
		req.Header.Set("Content-Type", "application/json")

		// Create a response recorder to capture the response
		respRecorder := httptest.NewRecorder()

		// Serve the request
		router.ServeHTTP(respRecorder, req)

		// Verify the response
		assert.Equal(t, http.StatusOK, respRecorder.Code)
		// TODO: Add more assertions for the response body or cookies if needed
	})

	// Test case 2: Missing credentials
	t.Run("Test missing credentials", func(t *testing.T) {
		// Create a new HTTP request with missing credentials
		jsonData := `{"user_name": "", "pass_word": ""}`
		req, _ := http.NewRequest("POST", "/login", strings.NewReader(jsonData))
		req.Header.Set("Content-Type", "application/json")

		// Create a response recorder to capture the response
		respRecorder := httptest.NewRecorder()

		// Serve the request
		router.ServeHTTP(respRecorder, req)

		// Verify the response
		assert.Equal(t, http.StatusBadRequest, respRecorder.Code)
		// TODO: Add more assertions for the response body if needed
	})
}
