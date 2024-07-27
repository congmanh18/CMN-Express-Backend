package utils

import (
	user "cmn-express/domain/user/entity"
	"errors"
)

type UserType int

const (
	Administrator UserType = iota
	Driver
	Provider
	Unknown
)

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
