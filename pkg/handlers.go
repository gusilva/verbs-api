package pkg

import (
	"context"
	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
	"net/http"
	"time"
)

type Config struct {
	Log        *zap.Logger
	Repository VerbRepository
}

func (app *Config) findVerb(w http.ResponseWriter, r *http.Request) {
	verbName := chi.URLParam(r, "VERB_NAME")
	app.Log.Debug("find handler", zap.String("verbName", verbName))

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	verbConjugation, errorFindVerb := app.Repository.FindVerb(ctx, verbName)
	if errorFindVerb != nil {
		app.Log.Warn("error find verb", zap.String("msg", errorFindVerb.Error()))
		_ = app.errorJSON(w, errorFindVerb, http.StatusBadRequest)

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

func (app *Config) searchVerb(w http.ResponseWriter, r *http.Request) {
	queryWord := chi.URLParam(r, "QUERY_WORD")
	app.Log.Debug("search handler", zap.String("queryWord", queryWord))

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	verbs, errorSearchVerb := app.Repository.SearchVerb(ctx, queryWord)
	if errorSearchVerb != nil {
		app.Log.Warn("error search verb", zap.String("msg", errorSearchVerb.Error()))
		_ = app.errorJSON(w, errorSearchVerb, http.StatusBadRequest)

		return
	}

	payload := jsonResponse{
		Error:   false,
		Message: "verbs found",
		Data:    verbs,
	}

	if errorWriteJson := app.writeJSON(w, http.StatusOK, payload); errorWriteJson != nil {
		app.Log.Error("error writing json response", zap.Error(errorWriteJson))
	}
}
