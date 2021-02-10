package entity

import (
	"errors"
	"strconv"

	"github.com/CezarGarrido/crypto-index/api/helper"
)

var ErrInvalidEmail = errors.New("Invalid email")
var ErrInvalidPassword = errors.New("Invalid password")


// User: Structure that represents a user
type User struct {
	Email    string // name: represents the user name - required
	Password string// password: represents the user name - required
}

// Validate: Function responsible for validating user data.
// If the email or password is in the wrong format, the function returns an error.
func (u *User) Validate() error {
	if len(u.Password) < 6 {
		return ErrInvalidPassword
	}
	_, err := strconv.Atoi(u.Password)
	if err != nil {
		return ErrInvalidPassword
	}

	if !helper.IsEmailValid(u.Email) {
		return ErrInvalidEmail
	}
	return nil
}
