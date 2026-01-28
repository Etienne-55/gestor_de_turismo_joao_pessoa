package controllers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"projeto_turismo_jp/models"
	"projeto_turismo_jp/server"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)


func TestSignup(t *testing.T) {
	server := server.SetupServer()
	server.POST("/signup", Signup)

	w := httptest.NewRecorder()

	exampleTourist := models.Tourist{
		Email: "etienne@gmail.com",
		Password: "testpassword123",
	}

	userJSON, _ := json.Marshal(exampleTourist)
	req, _ := http.NewRequest("POST", "/signup", strings.NewReader(string(userJSON)))

	req.Header.Set("Content-Type", "application/json")

	server.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, string(userJSON), w.Body.String())
}

