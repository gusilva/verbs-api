package main

import (
	"fmt"
	"github.com/gusilva/verbs-api/pkg"
	"log"
	"net/http"
	"os"
)

const (
	webPort         = "8088"
	mongoDefaultDSN = "mongodb://localhost:27017"
)

func main() {
	flags := log.LstdFlags | log.Lshortfile
	l := &pkg.Logger{
		InfoLogger:  log.New(os.Stdout, "[INFO] ", flags),
		ErrorLogger: log.New(os.Stdout, "[ERROR] ", flags),
		DebugLogger: log.New(os.Stdout, "[DEBUG] ", flags),
		IsDevMode:   true,
	}

	mongoDB := &pkg.MongoDatabase{
		DSN: mongoDefaultDSN,
		Log: l,
	}

	if mongoEnvUrl := os.Getenv("MONGO_DSN"); mongoEnvUrl != "" {
		mongoDB.DSN = mongoEnvUrl
	} else {
		l.Info("MONGO_DSN env is not set")
	}

	if errorConnectDB := mongoDB.ConnectToMongo(); errorConnectDB != nil {
		l.Error("error on connect to mongo DB:", errorConnectDB.Error())
		panic(errorConnectDB)
	}

	mongoDB.PingDB()

	collection, errorCollection := mongoDB.GetCollection("conjugations")
	if errorCollection != nil {
		l.Error(errorCollection.Error())
	}

	verbRepository := pkg.CreateVerbRepository(collection)

	app := pkg.Config{
		Log:        l,
		Repository: verbRepository,
	}

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.Routes(),
	}

	l.Info("Start server port -", webPort)
	if errorListenServer := srv.ListenAndServe(); errorListenServer != nil {
		log.Panic("listen server error:", errorListenServer.Error())
	}
}
