package main

import (
	"fmt"
	"github.com/gusilva/verbs-api/pkg"
	"go.uber.org/zap"
	"log"
	"net/http"
)

const webPort = "8088"

func main() {
	logger, errorZapLogger := zap.NewProduction()
	if errorZapLogger != nil {
		panic(errorZapLogger)
	}

	defer func() {
		if errorLoggerSync := logger.Sync(); errorLoggerSync != nil {
			fmt.Printf("error sync zap logger: %s", errorLoggerSync.Error())
		}
	}()

	app := pkg.Config{
		Log: logger,
	}

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.Routes(),
	}

	logger.Info("Start server", zap.String("port", webPort))
	if errorListenServer := srv.ListenAndServe(); errorListenServer != nil {
		log.Panic(errorListenServer)
	}
}
