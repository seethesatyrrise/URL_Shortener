package storage

import "context"

//go:generate mockgen -source=interface.go -destination=mocks/mock.go

type Storage interface {
	SaveData(context.Context, string, string) error
	GetLinkByToken(context.Context, string) (string, error)
	TryGetTokenByLink(context.Context, string) (string, error)
}
