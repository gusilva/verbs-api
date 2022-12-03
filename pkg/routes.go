package pkg

import "net/http"

func (app *Config) Routes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/ping", app.ping)

	return mux
}
