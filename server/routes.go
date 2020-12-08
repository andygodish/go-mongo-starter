package server

import (
	"time"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

const (
	pathOrganization string = "/organization"
)

func (s *Server) InitRoutePaths() {
	s.Router.Use(middleware.RequestID)
	s.Router.Use(middleware.RequestID)
	s.Router.Use(middleware.RequestID)
	s.Router.Use(middleware.RequestID)
	s.Router.Use(middleware.Timeout(60 * time.Second))
	s.Router.Use(middleware.NoCache)

	s.Router.Group(func(r chi.Router) {
		r.Get(pathOrganization, s.handleGETOrganization)
	})	
}  

