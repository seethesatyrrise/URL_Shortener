package storage

type Storage interface {
	SaveData(string, string) error
	GetLinkByToken(string) (string, error)
	TryGetTokenByLink(string) (string, error)
}
