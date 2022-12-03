package pkg

import (
	"encoding/json"
	"go.uber.org/zap"
	"net/http"
)

type Config struct {
	Log *zap.Logger
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
