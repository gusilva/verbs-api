package pkg

import (
	"encoding/json"
	"errors"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestWriteJSON(t *testing.T) {
	app := &Config{}
	w := httptest.NewRecorder()

	data := jsonResponse{
		Error:   false,
		Message: "test message",
		Data:    "test data",
	}

	errorWriteJSON := app.writeJSON(w, http.StatusOK, data)
	assert.NoError(t, errorWriteJSON)

	result := w.Result()
	assert.Equal(t, http.StatusOK, result.StatusCode)
	assert.Equal(t, "application/json", result.Header.Get("Content-Type"))

	var response jsonResponse
	errorResponse := json.NewDecoder(result.Body).Decode(&response)
	assert.NoError(t, errorResponse)
	assert.Equal(t, data.Error, response.Error)
	assert.Equal(t, data.Message, response.Message)
	assert.Equal(t, data.Data, response.Data)
}

func TestErrorJSON(t *testing.T) {
	app := &Config{}
	w := httptest.NewRecorder()

	errorJSON := app.errorJSON(w, errors.New("json error response"))
	assert.NoError(t, errorJSON)

	result := w.Result()
	assert.Equal(t, http.StatusBadRequest, result.StatusCode)
	assert.Equal(t, "application/json", result.Header.Get("Content-Type"))

	var response jsonResponse
	errorResponse := json.NewDecoder(result.Body).Decode(&response)

	assert.NoError(t, errorResponse)
	assert.Truef(t, response.Error, "Expected error to be true")
	assert.Equal(t, "json error response", response.Message)
}
