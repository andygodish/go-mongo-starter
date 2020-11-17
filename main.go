package main

import "time"

const (
	defaultPort int = 8080

	// default HTTP server limits - based on bigbird project/P1 standards
	defaultHTTPReadTimeout       time.Duration = 90 * time.Second
	defaultHTTPReadHeaderTimeout time.Duration = 90 * time.Second
	defaultHTTPWriteTimeout      time.Duration = 120 * time.Second
	defaultHTTPIdleTImeout       time.Duration = 120 * time.Second // Keep-Alive timeout
	defaultMaxHeaderBytes        int           = 1024 * 1024
	defaultMaxBodyBytes          int64         = 5 * 1024 * 1024

	// Should be used to connect to mongodb container running locally in your development environment
	devMongoDBURI string = "mongodb://flashcard:mysecretpassword@127.0.0.1:27017/my_database?authSource=admin"

	pathAPI = "/api"
)

func main() {

}
