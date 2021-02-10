package entity

import (
	"errors"
	"strconv"

	"github.com/CezarGarrido/crypto-index/api/helper"
)
// ErrInvalidEmail: represents an invalid user.Email error
var ErrInvalidEmail = errors.New("Invalid email")
// ErrInvalidEmail: represents an invalid user.Password error
var ErrInvalidPassword = errors.New("Invalid password")

// User: Structure that represents a user
// - Name: required
// - Password: required
// Example:
// user := entity.User{Email: "test@mail", Password: "123456"}
// or
// var user User
// user.Email = "test@mail"
// user.Password = "123456"

type User struct {
	Email    string // name: represents the user name - required
	Password string // password: represents the user name - required
}

// Validate: Function responsible for validating user data.
// If the email or password is in the wrong format, the function returns an error.
// Example:
// user := entity.User{"test@mail", "123456"}
// if err := user.Validate(); err != nil { panic(err) } // its work

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
