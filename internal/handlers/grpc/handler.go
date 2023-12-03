package grpchandler

import (
	"URL_Shortener/internal/grpc/api"
	"URL_Shortener/internal/service"
)

type Handler struct {
	api.UnimplementedURLShortenerServer
	service *service.Service
}

func New(service *service.Service) *Handler {
	return &Handler{service: service}
}
