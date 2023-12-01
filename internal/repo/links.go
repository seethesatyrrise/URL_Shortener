package repo

import (
	"context"
)

func (r *Repo) GetFullLink(ctx context.Context, token string) (string, error) {
	link, err := (*r.storage).GetLinkByToken(token)
	if err != nil {
		return "", err
	}

	return link, nil
}

func (r *Repo) GetToken(ctx context.Context, link string) (string, error) {
	token, err := (*r.storage).TryGetTokenByLink(link)
	if err != nil {
		return "", err
	}
	if token != "" {
		return token, nil
	}

	token = generateToken()
	err = (*r.storage).SaveData(link, token)
	if err != nil {
		return "", err
	}

	return token, nil
}
