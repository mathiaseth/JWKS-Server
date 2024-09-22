package main

import (
	"encoding/base64"
	"math/big"
	"time"
)

// JWKSKey represents a JSON Web Key
type JWKSKey struct {
	Kty string `json:"kty"`
	N   string `json:"n"`
	E   string `json:"e"`
	Kid string `json:"kid"`
}

// getValidKeys retrieves valid keys
func getValidKeys() []JWKSKey {
	validKeys := []JWKSKey{}
	for _, key := range keys {
		if key.Expiry.After(time.Now()) {
			validKeys = append(validKeys, JWKSKey{
				Kty: "RSA",
				N:   encodeBase64URL(key.PublicKey.N.Bytes()),
				E:   encodeBase64URL(big.NewInt(int64(key.PublicKey.E)).Bytes()),
				Kid: key.Kid,
			})
		}
	}
	return validKeys
}

// encodeBase64URL encodes bytes to Base64 URL format
func encodeBase64URL(input []byte) string {
	// Use base64.RawURLEncoding to avoid padding
	return base64.RawURLEncoding.EncodeToString(input)
}
