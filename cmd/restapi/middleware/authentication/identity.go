package authentication

import (
	"github.com/dgrijalva/jwt-go"
)

// Identity holds the details of currently logged user
type Identity interface {
	GetUsername() string
	ToClaims() jwt.MapClaims
}
