package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/andygodish/go-mongo-starter/server"
	mongostore "github.com/andygodish/go-mongo-starter/stores/mongodb"
	"github.com/andygodish/go-mongo-starter/utils/helpers"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

const (
	envMongoDBURI string = "MONGODB_URI"

	defaultPort int = 8080

	// default HTTP server limits - based on bigbird project/P1 standards
	defaultHTTPReadTimeout       time.Duration = 90 * time.Second
	defaultHTTPReadHeaderTimeout time.Duration = 90 * time.Second
	defaultHTTPWriteTimeout      time.Duration = 120 * time.Second
	defaultHTTPIdleTImeout       time.Duration = 120 * time.Second // Keep-Alive timeout
	defaultMaxHeaderBytes        int           = 1024 * 1024
	defaultMaxBodyBytes          int64         = 5 * 1024 * 1024

	// Should be used to connect to mongodb container running locally in your development environment
	devMongoDBURI string = "mongodb://app:mysecretpassword@127.0.0.1:27017/my_database?authSource=admin"

	pathAPI = "/api"
)

var (
	defaultAppPath string = pathAPI
)

func main() {
	mongoDBURI := helpers.GetEnv("MONGODB_URI", devMongoDBURI)
	log.Printf("MongoDB URI: %s", mongoDBURI)

	dataStore, err := mongostore.Constructor(mongoDBURI)
	if err != nil {
		log.Printf("Unable to create mongoDB datastore: %v", err)
	}

	server := server.Constructor(defaultAppPath, dataStore, 1000000)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome"))
	})

	httpServer := &http.Server{
		Addr:    fmt.Sprintf(":%d", defaultPort),
		Handler: r,
	}

	if err := httpServer.ListenAndServe(); err != nil {
		log.Fatal(err)
	}

	fmt.Print(server)
}
