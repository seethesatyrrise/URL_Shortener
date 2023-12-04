package repo

import (
	"URL_Shortener/internal/storage"
	"sync"
)

type Repo struct {
	storage *storage.Storage
	mtx     sync.Mutex
}

func New(storage *storage.Storage) *Repo {
	return &Repo{
		storage: storage,
		mtx:     sync.Mutex{},
	}
}
