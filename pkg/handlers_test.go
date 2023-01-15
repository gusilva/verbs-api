package pkg

import (
	"context"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

type mockLogger struct{}

func (m *mockLogger) Info(v ...interface{})  {}
func (m *mockLogger) Debug(v ...interface{}) {}
func (m *mockLogger) Error(v ...interface{}) {}

type mockVerbRepository struct{}

func (m *mockVerbRepository) FindVerb(ctx context.Context, name string) (*VerbConjugation, error) {

	return &VerbConjugation{
		Verb: "hacer",
		Imperative: ImperativeTime{
			Negative: ConjugationImperative{
				I:         "-",
				It:        "no se haga",
				They:      "no te hagas",
				We:        "no se hagan",
				YouPlural: "no nos hagamos",
			},
		},
	}, nil
}

func (m *mockVerbRepository) SearchVerb(ctx context.Context, name string) ([]string, error) {
	return []string{
		"universalizar",
		"sobresalir",
		"salivar",
		"salir",
		"nasalizar",
		"desalinizar",
	}, nil
}

func TestFindVerb(t *testing.T) {
	app := &Config{
		Log:        &mockLogger{},
		Repository: &mockVerbRepository{},
	}
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/find/verb/hacer", nil)
	chi.URLParam(r, "VERB_NAME")

	app.findVerb(w, r)

	assert.Equal(t, http.StatusOK, w.Code)

	var response jsonResponse
	errorFindResponse := json.NewDecoder(w.Body).Decode(&response)

	assert.NoError(t, errorFindResponse)
	assert.False(t, response.Error)
	assert.Equal(t, "verb found", response.Message)

	dataByte, errorEncodeData := json.Marshal(response.Data)
	assert.NoError(t, errorEncodeData)

	var verbConjugation VerbConjugation
	errorDecodeVerbConjugation := json.Unmarshal(dataByte, &verbConjugation)
	assert.NoError(t, errorDecodeVerbConjugation)

	assert.Equal(t, "hacer", verbConjugation.Verb)
	assert.Equal(t, "-", verbConjugation.Imperative.Negative.I)
	assert.Equal(t, "no se haga", verbConjugation.Imperative.Negative.It)
	assert.Equal(t, "no te hagas", verbConjugation.Imperative.Negative.They)
	assert.Equal(t, "no se hagan", verbConjugation.Imperative.Negative.We)
	assert.Equal(t, "", verbConjugation.Imperative.Negative.You)
	assert.Equal(t, "no nos hagamos", verbConjugation.Imperative.Negative.YouPlural)
}

func TestSearchVerb(t *testing.T) {
	app := &Config{
		Log:        &mockLogger{},
		Repository: &mockVerbRepository{},
	}
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/search/verb/sali", nil)
	chi.URLParam(r, "QUERY_WORD")

	app.searchVerb(w, r)

	assert.Equal(t, http.StatusOK, w.Code)

	var response jsonResponse
	errorSearchResponse := json.NewDecoder(w.Body).Decode(&response)

	assert.NoError(t, errorSearchResponse)
	assert.False(t, response.Error)
	assert.Equal(t, "verbs found", response.Message)

	dataResponse, hasDataResponse := response.Data.([]interface{})
	assert.True(t, hasDataResponse)

	expected := []string{"universalizar", "sobresalir", "salivar", "salir", "nasalizar", "desalinizar"}

	assert.Equal(t, len(expected), len(dataResponse))
	for idx, verb := range dataResponse {
		assert.EqualValues(t, expected[idx], verb)
	}
}
