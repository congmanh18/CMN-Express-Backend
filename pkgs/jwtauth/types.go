package jwtauth

import "github.com/golang-jwt/jwt/v4"

// CustomClaims defines the structure of the JWT claims
type CustomClaims struct {
	UserID string `json:"user_id"`
	Email  string `json:"email"`
	jwt.RegisteredClaims
}
