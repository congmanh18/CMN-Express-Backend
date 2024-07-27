package user

import (
	"errors"
	"regexp"
	"time"
)

type User struct {
	ID            string  `gorm:"primaryKey"`
	Username      *string `gorm:"unique"`
	Email         *string `gorm:"unique"`
	Password      *string
	Role          *string
	First_name    *string
	Last_name     *string
	Token         *string
	Refresh_token *string
	Created_at    time.Time
	Updated_at    time.Time
}

func (u User) isEmailValid() bool {
	if u.Email == nil {
		return false
	}
	emailRegex := `^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,}$`
	re := regexp.MustCompile(emailRegex)
	return re.MatchString(*u.Email)
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
	if !u.isEmailValid() {
		return errors.New("email invalid")
	}
	if !u.areRequiredFieldsPresent() {
		return errors.New("required fields are missing")
	}
	if !u.isFieldLengthValid(u.Username, 3, 20) {
		return errors.New("username length invalid")
	}
	if !u.isFieldLengthValid(u.Password, 8, 64) {
		return errors.New("password length invalid")
	}
	return nil
}
