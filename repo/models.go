package repo

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
)

const (
	usernameClaim = "username"
)

var (
	UsernameErr = errors.New("invalid username")
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// NewUser creates a new User from the claims. It needs username claim to create a valid user.
// If the username claim is missing then UsernameErr is returned
func NewUser(claims jwt.MapClaims) (*User, error) {
	username, success := claims[usernameClaim].(string)
	if !success {
		return nil, UsernameErr
	}

	usr := &User{
		Username: username,
	}
	return usr, nil
}

// GetUsername returns username for the user. It implements Identity interface
func (u User) GetUsername() string {
	return u.Username
}

// ToClaims converts the existing user to Claims
func (u User) ToClaims() jwt.MapClaims {
	return jwt.MapClaims{
		usernameClaim: u.Username,
	}
}
