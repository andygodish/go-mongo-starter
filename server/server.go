package server

import "github.com/andygodish/go-mongo-starter/stores"

// Server comment
type Server struct {
	appPath      string
	Database     stores.Storer
	maxBodyBytes int64
}

// Constructor comment
func Constructor(appPath string, store stores.Storer, maxBodyBytes int64) *Server {
	s := &Server{
		appPath:      appPath,
		Database:     store,
		maxBodyBytes: maxBodyBytes,
	}

	return s
}
