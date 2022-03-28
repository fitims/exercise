package authentication

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/fitims/exercise/cmd/restapi/middleware/authentication/tokens"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
)

const (
	CookieName = "x-access-token"
	AuthHeader = "Authorization"
	AuthUser   = "user"
)

// JSON is a map of string representation of JSON
type JSON map[string]interface{}

// ClaimsValidator is a function that receives a set of claims (username and id), verifies if the claims are valid
// then if the claims are valid then a user is returned, otherwise an error is returned
type ClaimsValidator func(claims jwt.MapClaims) (Identity, error)

// Authenticator is the interface that defines the behaviour of the authentication middleware
type Authenticator interface {
	AuthenticateUser(next echo.HandlerFunc) echo.HandlerFunc
}

// defaultAuthenticator is the implementation of the authentication middleware
type defaultAuthenticator struct {
	validator ClaimsValidator
}

// NewUserAuthenticator creates a validation middleware
func NewUserAuthenticator(v ClaimsValidator) Authenticator {
	return defaultAuthenticator{
		validator: v,
	}
}

// AuthenticateUser will try and receive the token form authorization header (Authorization), if it is not found
// then it will try to get if from the authorization cookie (x-access-token). If the authorization cookie is
// missing as well, then status 401 (Unauthorized) is returned with the message "Invalid Token"
//
// If the token is found then it will check if the token is Bearer token or Renew token and will try to decode it.
// If the decoding fails then 401 (Unauthorized) is returned with the message "Invalid Token"
//
// If the token is decoded, then the validator is called passing the claims (username and id), and if the validator
// validates the user, and if validation succeeds then the middleware will pass the call back, otherwise
// 401 (Unauthorized) is returned with the message "Could not Validate user"
func (a defaultAuthenticator) AuthenticateUser(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var token string
		header := c.Request().Header[AuthHeader]
		if len(header) > 0 {
			token = extractToken(header[0])
		} else {
			cookie, err := c.Cookie(CookieName)
			if err != nil {
				return c.JSON(http.StatusUnauthorized, JSON{"message": "User is not authenticated"})
			}
			token = extractToken(cookie.Value)
		}

		claims, tokenErr := tokens.DecodeToken(token)
		if tokenErr == nil {
			usr, err := a.validator(claims)
			if err != nil {
				return c.JSON(http.StatusUnauthorized, JSON{"message": "Could not Validate user"})
			}

			c.Set(AuthUser, usr)
			return next(c)
		}

		return c.JSON(http.StatusUnauthorized, JSON{"message": "Invalid token"})
	}
}

// extractToken will remove the words "bearer" or "access" from the token string
func extractToken(header string) string {
	lower := strings.ToLower(header)
	if strings.HasPrefix(lower, "bearer") {
		return header[7:]
	}

	if strings.HasPrefix(lower, "access") {
		return header[7:]
	}

	return header
}
