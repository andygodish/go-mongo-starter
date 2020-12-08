package server

import (
	"github.com/go-chi/chi"
	"github.com/andygodish/go-mongo-starter/stores"
)

// Server comment
type Server struct {
	appPath      string
	Database     stores.Storer
	maxBodyBytes int64
	Router       *chi.Mux
}

// Constructor comment
func Constructor(appPath string, store stores.Storer, maxBodyBytes int64) *Server {
	s := &Server{
		appPath:      appPath,
		Database:     store,
		maxBodyBytes: maxBodyBytes,
		Router:       chi.NewRouter(),
	}

	return s
}
