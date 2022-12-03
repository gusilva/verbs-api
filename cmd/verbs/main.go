package main

import (
	"fmt"
	"github.com/gusilva/verbs-api/pkg"
	"log"
	"net/http"
)

const webPort = "8088"

func main() {
	app := pkg.Config{}

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.Routes(),
	}

	log.Printf("Start server on port: %s\n", webPort)
	if errorListenServer := srv.ListenAndServe(); errorListenServer != nil {
		log.Panic(errorListenServer)
	}
}
