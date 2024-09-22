package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

func main() {
	// Generate an initial key for demonstration purposes
	_, err := generateRSAKey("key1", 24*time.Hour)
	if err != nil {
		log.Fatalf("Error generating key: %v", err)
	}

	http.HandleFunc("/.well-known/jwks.json", jwksHandler)
	http.HandleFunc("/auth", authHandler)

	log.Println("Server running on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

// jwksHandler serves the JWKS endpoint
func jwksHandler(w http.ResponseWriter, r *http.Request) {
	validKeys := getValidKeys()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string][]JWKSKey{"keys": validKeys})
}

// authHandler issues a signed JWT (implementation needed)
func authHandler(w http.ResponseWriter, r *http.Request) {
	// Check if the "expire" query parameter is present
	expired := r.URL.Query().Get("expire") == "true"

	// Create JWT using the appropriate key
	token, err := createJWT("key1", expired)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set the response header and write the token
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}
