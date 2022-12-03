package pkg

import (
	"encoding/json"
	"net/http"
)

type Config struct{}

func (app *Config) ping(w http.ResponseWriter, _ *http.Request) {
	status := map[string]string{
		"status": "ok",
	}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(status)
}
