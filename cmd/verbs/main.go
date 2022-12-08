package main

import (
	"fmt"
	"github.com/gusilva/verbs-api/pkg"
	"go.uber.org/zap"
	"net/http"
	"os"
)

const (
	webPort         = "8088"
	mongoDefaultDSN = "mongodb://localhost:27017"
)

func main() {
	logger, errorZapLogger := zap.NewDevelopment()
	if errorZapLogger != nil {
		panic(errorZapLogger)
	}

	defer func() {
		if errorLoggerSync := logger.Sync(); errorLoggerSync != nil {
			fmt.Printf("error sync zap logger: %s", errorLoggerSync.Error())
		}
	}()

	mongoDB := &pkg.MongoDatabase{
		DSN:    mongoDefaultDSN,
		Logger: logger,
	}

	if mongoEnvUrl := os.Getenv("MONGO_DSN"); mongoEnvUrl != "" {
		mongoDB.DSN = mongoEnvUrl
	} else {
		logger.Warn("MONGO_DSN env is not set")
	}

	if errorConnectDB := mongoDB.ConnectToMongo(); errorConnectDB != nil {
		logger.Error("error on connect to mongo DB", zap.Error(errorConnectDB))
		panic(errorConnectDB)
	}

	mongoDB.PingDB()

	collection, errorCollection := mongoDB.GetCollection("conjugations")
	if errorCollection != nil {
		logger.Error(errorCollection.Error())
		panic(errorCollection)
	}

	verbRepository := pkg.CreateVerbRepository(collection)

	app := pkg.Config{
		Log:        logger,
		Repository: verbRepository,
	}

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.Routes(),
	}

	logger.Debug("Start server", zap.String("port", webPort))
	if errorListenServer := srv.ListenAndServe(); errorListenServer != nil {
		logger.Panic("listen server error", zap.Error(errorListenServer))
	}
}
