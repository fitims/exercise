package controllers

import (
	"errors"
	"github.com/fitims/exercise/cmd/restapi/middleware/authentication/tokens"
	"github.com/fitims/exercise/cmd/restapi/routes/controllers/requests"
	"github.com/fitims/exercise/cmd/restapi/routes/controllers/responses"
	"github.com/fitims/exercise/repo"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"unicode"
)

type UserController interface {
	Register(c echo.Context) error
	Login(c echo.Context) error
}

type defaultUserController struct {
	repository repo.UserRepository
}

// NewUserController creates new user controller that handles user routes
func NewUserController(r repo.UserRepository) UserController {
	return &defaultUserController{
		repository: r,
	}
}

// Register handles registration
func (ctrl defaultUserController) Register(c echo.Context) error {
	var req requests.RegisterRequest
	err := c.Bind(&req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.RegistrationFailed("The request is not valid"))
	}

	hash, ok := validatePassword(req.Password)
	if !ok {
		return c.JSON(http.StatusBadRequest, responses.RegistrationFailed("password is too weak"))
	}

	err = ctrl.repository.RegisterUser(req.Username, string(hash))
	if err != nil {
		if errors.Is(err, repo.UserExistsErr) {
			return c.JSON(http.StatusBadRequest, responses.RegistrationFailed("User is already registered"))
		}
		return c.JSON(http.StatusInternalServerError, responses.RegistrationFailed("Could not get user details. please try later on."))
	}

	return c.JSON(http.StatusOK, responses.RegistrationSuccessful())
}

// Login handles login
func (ctrl defaultUserController) Login(c echo.Context) error {
	var req requests.LoginRequest
	err := c.Bind(&req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.LoginFailed("The request is not valid"))
	}

	// get user from the repository
	usr, err := ctrl.repository.GetUser(req.Username)
	if err != nil {
		if errors.Is(err, repo.UserDoesNotExistErr) {
			return c.JSON(http.StatusBadRequest, responses.LoginFailed("User does not exist"))
		}
		return c.JSON(http.StatusInternalServerError, responses.LoginFailed("Could not get user details. please try later on."))
	}

	// compare passwords
	err = bcrypt.CompareHashAndPassword([]byte(usr.Password), []byte(req.Password))
	if err != nil {
		return c.JSON(http.StatusUnauthorized, responses.LoginFailed("Invalid credentials"))
	}

	// generate login tokens
	authTokens, err := tokens.GenerateAuthTokens(usr.ToClaims())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.LoginFailed("could not generate login tokens"))
	}
	return c.JSON(http.StatusOK, responses.LoginSuccessful(&authTokens))
}

// validatePassword check if the passwords meets the criteria. The criteria are that the password must:
// - be at least 7 characters long
// - has at least one Upper Case character
// - has at least one Lower Case character
// - has at least one Number
// - has at least one Special character
// If the password meets the criteria then password is hashed and is returned with ok
func validatePassword(pwd string) ([]byte, bool) {
	var (
		hasMinLen  = false
		hasUpper   = false
		hasLower   = false
		hasNumber  = false
		hasSpecial = false
	)
	if len(pwd) >= 7 {
		hasMinLen = true
	}
	for _, char := range pwd {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsNumber(char):
			hasNumber = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		}
	}

	isValid := hasMinLen && hasUpper && hasLower && hasNumber && hasSpecial
	if isValid {
		hashed, _ := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
		return hashed, true
	}
	return nil, false
}
