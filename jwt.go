package main

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// Declare ErrKeyNotFound globally
var ErrKeyNotFound = errors.New("key not found")

func createJWT(kid string, expired bool) (string, error) {
	var key *Key
	for _, k := range keys {
		if k.Kid == kid {
			key = &k
			break
		}
	}

	if key == nil {
		return "", ErrKeyNotFound
	}

	expirationTime := time.Now().Add(5 * time.Minute)
	if expired {
		expirationTime = time.Now().Add(-5 * time.Minute)
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"exp": expirationTime.Unix(),
		"iat": time.Now().Unix(),
	})

	token.Header["kid"] = kid
	return token.SignedString(key.PrivateKey)
}
