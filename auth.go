package main

import (
	"net/http"
)

func AuthHandler(w http.ResponseWriter, r *http.Request) {
	expired := r.URL.Query().Get("expired") != ""
	kid := "key1" // Hardcoded for simplicity
	token, err := createJWT(kid, expired)
	if err != nil {
		http.Error(w, "Error generating JWT", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/jwt")
	w.Write([]byte(token))
}
