package entity

import (
	"errors"
	"time"
)

type User struct {
	ID            string  `gorm:"primaryKey"`
	Username      *string `gorm:"unique"`
	Phone         *string `gorm:"unique"`
	Password      *string
	Role          *string
	First_name    *string
	Last_name     *string
	Token         *string
	Refresh_token *string
	Created_at    time.Time
	Updated_at    time.Time
}

func (u User) isPhoneValid() bool {
	return len(*u.Phone) == 10
}

func (u User) isFieldLengthValid(field *string, minLength int, maxLength int) bool {
	if field == nil {
		return false
	}
	length := len(*field)
	return length >= minLength && length <= maxLength
}

func (u User) IsValidUser() error {
	if !u.isPhoneValid() {
		return errors.New("phone invalid")
	}
	if !u.isFieldLengthValid(u.Password, 8, 64) {
		return errors.New("password length invalid")
	}
	return nil
}
