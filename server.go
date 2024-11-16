package doodocsrestapi

import (
	"context"
	"net/http"
	"time"
)

type Server struct {
	httpSerever *http.Server
}

func (s *Server) Run(port string, handler http.Handler) error {
	s.httpSerever = &http.Server{
		Addr:           ":" + port,
		Handler:        handler,
		MaxHeaderBytes: 1 << 20,
		ReadTimeout:    1 * time.Second,
		WriteTimeout:   1 * time.Second,
	}

	return s.httpSerever.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpSerever.Shutdown(ctx)
}
