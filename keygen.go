package main

import (
	"crypto/rand"
	"crypto/rsa"
	"time"
)

type Key struct {
	PrivateKey *rsa.PrivateKey
	PublicKey  rsa.PublicKey
	Kid        string
	Expiry     time.Time
}

var keys []Key

func generateRSAKey(kid string, expiry time.Duration) (Key, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return Key{}, err
	}

	key := Key{
		PrivateKey: privateKey,
		PublicKey:  privateKey.PublicKey,
		Kid:        kid,
		Expiry:     time.Now().Add(expiry),
	}
	keys = append(keys, key)
	return key, nil
}
