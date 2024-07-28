package utils

import (
	user "cmn-express/domain/user/entity"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type UserType int

const (
	Administrator UserType = iota
	Driver
	Provider
	Unknown
)

// HashPassword hashes the given password using bcrypt.
func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

// VerifyPassword compares a hashed password with its plaintext version.
func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

// CheckUserType determines the type of a user based on their role.
func CheckUserType(u user.User) UserType {
	switch *u.Role {
	case "administrator":
		return Administrator
	case "driver":
		return Driver
	case "provider":
		return Provider
	default:
		return Unknown
	}
}

// MatchUserToID checks if the given user ID matches the ID of the provided user.
func MatchUserToID(u user.User, id string) (bool, error) {
	if u.ID == id {
		return true, nil
	}
	return false, errors.New("user ID does not match")
}
