package jwtauth

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// Define your secret key
var jwtKey = []byte("your_secret_key")

// GenerateAllToken generates the access and refresh tokens
func GenerateAllToken(userID, email string) (accessToken, refreshToken string, err error) {
	// Define the token expiration times
	accessTokenExpiresAt := time.Now().Add(15 * time.Minute)
	refreshTokenExpiresAt := time.Now().Add(24 * time.Hour)

	// Create the claims
	claims := CustomClaims{
		UserID: userID,
		Email:  email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(accessTokenExpiresAt),
		},
	}

	// Create the access token
	accessToken, err = createToken(claims)
	if err != nil {
		return "", "", err
	}

	// Update the claims for refresh token
	claims.RegisteredClaims.ExpiresAt = jwt.NewNumericDate(refreshTokenExpiresAt)

	// Create the refresh token
	refreshToken, err = createToken(claims)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

// UpdateAllToken updates the access and refresh tokens
func UpdateAllToken(refreshToken string) (newAccessToken, newRefreshToken string, err error) {
	// Validate the refresh token
	claims, err := ValidateToken(refreshToken)
	if err != nil {
		return "", "", err
	}

	// Generate new tokens
	return GenerateAllToken(claims.UserID, claims.Email)
}

// ValidateToken validates the given token
func ValidateToken(tokenString string) (*CustomClaims, error) {
	claims := &CustomClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, errors.New("invalid token")
	}
	return claims, nil
}

// Helper function to create a token
func createToken(claims CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}
