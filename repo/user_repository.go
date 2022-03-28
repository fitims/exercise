package repo

import (
	"errors"
	"github.com/fitims/exercise/cmd/restapi/middleware/authentication"
)

type UserRepository interface {
	GetUser(username string) (*User, error)
	SaveUser(username, password string) error
	GetUserValidator() authentication.ClaimsValidator
}

var (
	UserExistsErr       = errors.New("user already exists")
	UserDoesNotExistErr = errors.New("user does not exist")
)
