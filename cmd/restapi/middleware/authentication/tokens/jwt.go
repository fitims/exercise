package tokens

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

const (
	JwtSecret = "界 - 出る釘は打たれる -=# this is a JWT Secret !!! #=- 界"
	Day       = time.Hour * 24
)

var (
	InvalidTokenError = errors.New("invalid token")
	TokenExpiredError = errors.New("token has expired")
	ClaimsError       = errors.New("could not decode claims")
	UnknownTokenError = errors.New("unknown error")
)

type BearerToken string
type AccessToken string

// TokenPair holds a pair of tokens ie. the bearer and renew tokens.
// The pair is returned when the user logs in.
//
// The bearer token is usually used for accessing the resources, and it is valid for 10 minutes
// whereas the renew token is used to renew an expired bearer token, and it is valid for 2 hours
type TokenPair struct {
	RenewToken  AccessToken `json:"renew_token"`
	BearerToken BearerToken `json:"bearer_token"`
}

// GenerateAuthTokens generates bearer and renew tokens
func GenerateAuthTokens(claims jwt.MapClaims) (TokenPair, error) {
	rt, err := CreateRenewToken(claims)
	if err != nil {
		return TokenPair{}, err
	}

	bt, err := CreateBearerToken(claims)
	if err != nil {
		return TokenPair{}, err
	}

	return TokenPair{
		RenewToken:  rt,
		BearerToken: bt,
	}, nil
}

// CreateBearerToken generates bearer token that is valid for 10 mins
func CreateBearerToken(claims jwt.MapClaims) (BearerToken, error) {

	claims["iat"] = time.Now().Unix()
	claims["exp"] = time.Now().Add(time.Minute * 10).Unix() // expires in 10 mins

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(JwtSecret))
	return BearerToken(tokenString), err
}

// CreateRenewToken generates renew token that is valid for 2 hours
func CreateRenewToken(claims jwt.MapClaims) (AccessToken, error) {

	claims["iat"] = time.Now().Unix()
	claims["exp"] = time.Now().Add(time.Hour * 2).Unix() // expires in 2 hours

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(JwtSecret))
	return AccessToken(tokenString), err
}

// DecodeToken will try to decode the provided token.
// if the decoding is successful, firstly the token is checked if it is valid. if the token is valid
// then claims are returned, otherwise if the claims cannot be extracted then ClaimsError is returned.
// If the token is malformed, then InvalidTokenError is returned
// If the token has expired then TokenExpiredError is returned
//
// If decoding fails then UnknownTokenError is returned
func DecodeToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(JwtSecret), nil
	})

	if err == nil && token.Valid {
		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			return claims, nil
		}
		return nil, ClaimsError
	}

	if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			return nil, InvalidTokenError
		} else if ve.Errors&(jwt.ValidationErrorNotValidYet|jwt.ValidationErrorExpired) != 0 {
			return nil, TokenExpiredError
		}
	}

	return nil, UnknownTokenError
}
