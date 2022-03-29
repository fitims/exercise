package authentication

import (
	"github.com/dgrijalva/jwt-go"
)

// Identity holds the details of currently logged user. Any struct that
// implements this interface can be used as an Identity
type Identity interface {
	GetUsername() string
	ToClaims() jwt.MapClaims
}
