package server

import (
	"URL_Shortener/internal/config"
	"context"
	"go.uber.org/zap"
	"net/http"
)

type HTTPServer struct {
	server *http.Server
	logger *zap.Logger
}

func NewHTTPServer(cfg *config.HTTP, router http.Handler, logger *zap.Logger) (*HTTPServer, error) {
	s := &HTTPServer{
		server: &http.Server{
			Addr:    ":" + cfg.Port,
			Handler: router,
		},
		logger: logger,
	}
	return s, nil
}

func (s *HTTPServer) Start() error {
	go func() {
		if err := s.server.ListenAndServe(); err != nil {
			s.logger.Error(err.Error())
		}
	}()

	return nil
}

func (s *HTTPServer) Shutdown(ctx context.Context) error {
	if s.server != nil {
		return s.server.Shutdown(ctx)
	}

	return nil
}
