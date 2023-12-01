package service

import (
	"URL_Shortener/internal/repo"
	"context"
)

type TokenService struct {
	repo *repo.Repo
}

func NewTokenService(repo *repo.Repo) *TokenService {
	return &TokenService{repo: repo}
}

func (t *TokenService) GetToken(ctx context.Context, link string) (string, error) {
	return t.repo.GetToken(ctx, link)
}

func (t *TokenService) GetFullLink(ctx context.Context, token string) (string, error) {
	return t.repo.GetFullLink(ctx, token)
}
