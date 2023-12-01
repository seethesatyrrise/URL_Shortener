package storage

import "context"

type Storage interface {
	SaveData(context.Context, string, string) error
	GetLinkByToken(context.Context, string) (string, error)
	TryGetTokenByLink(context.Context, string) (string, error)
}
