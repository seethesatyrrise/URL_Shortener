package repo

import "URL_Shortener/internal/storage"

type Repo struct {
	storage *storage.Storage
}

func New(storage *storage.Storage) *Repo {
	return &Repo{storage: storage}
}
