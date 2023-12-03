package server

import (
	"URL_Shortener/internal/config"
	"URL_Shortener/internal/grpc/api"
	grpchandler "URL_Shortener/internal/handlers/grpc"
	"context"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net"
)

type GRPCServer struct {
	server *grpc.Server
	cfg    *config.GRPC
	logger *zap.Logger
}

func NewGRPCServer(cfg *config.GRPC, handler *grpchandler.Handler, logger *zap.Logger) (*GRPCServer, error) {
	s := &GRPCServer{
		server: grpc.NewServer(),
		cfg:    cfg,
		logger: logger,
	}
	api.RegisterURLShortenerServer(s.server, handler)
	return s, nil
}

func (g *GRPCServer) Start() error {
	grpcListener, err := net.Listen("tcp", ":"+g.cfg.Port)
	if err != nil {
		g.logger.Error(err.Error())
		return err
	}

	go func() {
		if err = g.server.Serve(grpcListener); err != nil {
			g.logger.Error(err.Error())
		}
	}()

	return nil
}

func (g *GRPCServer) Shutdown(ctx context.Context) error {
	if g.server != nil {
		g.server.GracefulStop()
	}

	return nil
}
