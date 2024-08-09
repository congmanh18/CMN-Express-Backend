package model

import "github.com/golang-jwt/jwt/v4"

type Token struct {
	AccessToken  *string `json:"access_token,omitempty"`
	RefreshToken *string `json:"refresh_token,omitempty"`
}

type JwtCustomClaims struct {
	ID          string `json:"id"`
	AccountType string `json:"account_type"`
	jwt.RegisteredClaims
}
