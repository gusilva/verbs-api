package main

import (
	"fmt"
	"github.com/gusilva/verbs-api/pkg"
	"go.uber.org/zap"
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

	database, errorDatabase := pkg.ConnectToMongo()
	if errorDatabase != nil {
		logger.Error("database error", zap.Error(errorDatabase))
	}

	verbRepository := pkg.CreateVerbRepository(database.Collection("conjugations"))

	app := pkg.Config{
		Log:        logger,
		Repository: verbRepository,
	}

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.Routes(),
	}

	logger.Info("Start server", zap.String("port", webPort))
	if errorListenServer := srv.ListenAndServe(); errorListenServer != nil {
		logger.Panic("listen server error", zap.Error(errorListenServer))
	}
}
