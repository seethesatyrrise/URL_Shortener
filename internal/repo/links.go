package repo

import (
	"context"
)

func (r *Repo) GetFullLink(ctx context.Context, token string) (string, error) {
	link, err := r.storage.GetLinkByToken(ctx, token)
	if err != nil {
		return "", err
	}

	return link, nil
}

func (r *Repo) GetToken(ctx context.Context, link string) (string, error) {
	r.mtx.Lock()
	defer r.mtx.Unlock()

	token, err := r.storage.TryGetTokenByLink(ctx, link)
	if err != nil {
		return "", err
	}
	if token != "" {
		return token, nil
	}

	token = generateToken()
	err = r.storage.SaveData(ctx, link, token)
	if err != nil {
		return "", err
	}

	return token, nil
}
