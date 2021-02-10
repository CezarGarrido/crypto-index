package entity_test

import (
	"testing"

	. "github.com/CezarGarrido/crypto-index/api/entity"
)

// TestNewUser: Creating test for new users
func TestNewUser(t *testing.T) {

	// test: Setup user and expected
	type test struct {
		user     User
		expected error
	}

	usersToTests := []test{
		{User{
			Email:    "invalid_email",
			Password: "123456",
		}, ErrInvalidEmail},
		{User{
			Email:    "test@coin.com",
			Password: "123456b",
		}, ErrInvalidPassword},
		{User{
			Email:    "",
			Password: "123456",
		}, ErrInvalidEmail},
		{User{
			Email:    "test@coin.com",
			Password: "",
		}, ErrInvalidPassword},

		{User{
			Email:    "test@coin.com",
			Password: "123456",
		}, nil},
		{User{
			Email:    "test@coin.com",
			Password: "123456789",
		}, nil},
	}

	for _, tt := range usersToTests {
		err := tt.user.Validate()
		if err != tt.expected {
			t.Error(tt.user.Email, tt.user.Password, err.Error())
		}
	}
}
