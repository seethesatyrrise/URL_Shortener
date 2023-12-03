package grpchandler

import (
	"URL_Shortener/internal/grpc/api"
	"URL_Shortener/internal/utils"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h Handler) GetLinkByToken(ctx context.Context, req *api.GetLinkRequest) (*api.GetLinkResponse, error) {
	link, err := h.service.GetFullLink(ctx, req.Token)
	if err != nil {
		if err == utils.ErrNotFound {
			return nil, status.Error(codes.NotFound, err.Error())
		}

		h.service.Logger.Error(err.Error())
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &api.GetLinkResponse{Link: link}, nil
}

func (h Handler) GetToken(ctx context.Context, req *api.GetTokenRequest) (*api.GetTokenResponse, error) {
	token, err := h.service.GetToken(ctx, req.Link)
	if err != nil {
		h.service.Logger.Error(err.Error())
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &api.GetTokenResponse{Token: token}, nil
}
