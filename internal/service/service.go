package service

import (
	"URL_Shortener/internal/repo"
	s "URL_Shortener/internal/storage"
	"context"
	"go.uber.org/zap"
)

type Token interface {
	GetToken(context.Context, string) (string, error)
	GetFullLink(context.Context, string) (string, error)
}

type Service struct {
	Token
	repo    *repo.Repo
	storage s.Storage
	Logger  *zap.Logger
}

func New(repo *repo.Repo, storage s.Storage, logger *zap.Logger) *Service {
	return &Service{
		Token:   NewTokenService(repo),
		repo:    repo,
		storage: storage,
		Logger:  logger,
	}
}
