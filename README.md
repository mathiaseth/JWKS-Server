# RESTful JWKS Server

This project implements a RESTful JWKS server in Go, designed to provide public keys for verifying JSON Web Tokens (JWTs). It features RSA key generation, authentication endpoints, and key expiry for enhanced security. The server adheres to best practices with organized code, comprehensive testing, and thorough documentation.

## Features

- **JWKS Endpoint**: Serves public keys in JWKS format with unique identifiers (kid).
- **Authentication Endpoint**: Issues signed JWTs supporting expired keys based on query parameters.
- **Key Expiry**: Ensures that only valid keys are served.
- **Unit Tests**: Includes a test suite with over 80% code coverage.

## Getting Started

### Prerequisites

- Go (version 1.16 or higher)

### Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/mathiaseth/jwks-server.git
   cd jwks-server

2. Run the server:
   ```bash
   go run main.go

4. Access the JWKS endpoint at [`http://localhost:8080/.well-known/jwks.json`](#http://localhost:8080/.well-known/jwks.json) and the authentication endpoint at [`http://localhost:8080/auth`](#http://localhost:8080/auth)

## Testing

- To run the tests and check coverage:
  ```bash
  go test -coverprofile=coverage.out
  go tool cover -html=coverage.out -o coverage.html
- Open [`coverage.html`](#coverage.html) in your browser to view the coverage report.

## Screenshots

[Test Client Screenshot](screenshots/Test_Client.png)  
[Test Suite Screenshot](screenshots/Test_Suite.png)
