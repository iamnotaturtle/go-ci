package server

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestServerGetsCharacter(t *testing.T) {
	r := CreateRouter()

	request, _ := http.NewRequest("GET", "/character/baby-yoda", nil)
	response := httptest.NewRecorder()

	r.ServeHTTP(response, request)

	assert.Equal(t, 200, response.Code)
	assert.Equal(t, "Character: baby-yoda\n", response.Body.String())
}
