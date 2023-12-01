package storage

import (
	"URL_Shortener/internal/utils"
	"context"
)

type MemoryStorage struct {
	linksByTokens map[string]string
	tokensByLinks map[string]string
}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		linksByTokens: make(map[string]string),
		tokensByLinks: make(map[string]string),
	}
}

func (ms *MemoryStorage) SaveData(ctx context.Context, link, token string) error {
	ms.linksByTokens[token] = link
	ms.tokensByLinks[link] = token
	return nil
}

func (ms *MemoryStorage) GetLinkByToken(ctx context.Context, token string) (string, error) {
	link, ok := ms.linksByTokens[token]
	if !ok {
		return "", utils.ErrNotFound
	}
	return link, nil
}

func (ms *MemoryStorage) TryGetTokenByLink(ctx context.Context, link string) (string, error) {
	token, _ := ms.tokensByLinks[link]
	return token, nil
}
