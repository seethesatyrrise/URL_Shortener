package repo

import (
	"math/rand"
	"time"
)

const (
	symbols     = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_"
	tokenLength = 10
)

func generateToken() string {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	token := make([]byte, tokenLength)
	for i := range token {
		token[i] = symbols[rnd.Intn(len(symbols))]
	}
	return string(token)
}
