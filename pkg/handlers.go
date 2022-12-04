package pkg

import (
	"context"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
	"net/http"
)

type Config struct {
	Log        *zap.Logger
	Repository VerbRepository
}

func (app *Config) ping(w http.ResponseWriter, _ *http.Request) {
	status := map[string]string{
		"status": "ok",
	}

	w.WriteHeader(http.StatusOK)
	if errorEncodeStatus := json.NewEncoder(w).Encode(status); errorEncodeStatus != nil {
		app.Log.Error("error decoding server status response: ", zap.String("error", errorEncodeStatus.Error()))
	}
}

func (app *Config) findVerb(w http.ResponseWriter, r *http.Request) {
	verbName := chi.URLParam(r, "VERB_NAME")

	verbConjugation, errorFindVerb := app.Repository.FindVerb(context.TODO(), verbName)
	if errorFindVerb != nil {
		errorWriteJson := app.errorJSON(w, errorFindVerb, http.StatusBadRequest)
		app.Log.Error("error writing error response", zap.Error(errorWriteJson))

		return
	}

	payload := jsonResponse{
		Error:   false,
		Message: "verb found",
		Data:    verbConjugation,
	}

	if errorWriteJson := app.writeJSON(w, http.StatusOK, payload); errorWriteJson != nil {
		app.Log.Error("error writing json response", zap.Error(errorWriteJson))
	}

}
