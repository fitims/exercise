package responses

import "github.com/fitims/exercise/cmd/restapi/middleware/authentication/tokens"

// LoginResponse is the response given when user tries to login. The login can
// be successful or it can fail
type LoginResponse struct {
	IsSuccess bool              `json:"is_success"`
	Message   string            `json:"message"`
	Tokens    *tokens.TokenPair `json:"tokens,omitempty"`
}

// LoginFailed build a login failed response
func LoginFailed(msg string) LoginResponse {
	return LoginResponse{
		IsSuccess: false,
		Message:   msg,
	}
}

// LoginSuccessful builds a successful response and contains auth tokens
func LoginSuccessful(tokens *tokens.TokenPair) LoginResponse {
	return LoginResponse{
		IsSuccess: true,
		Message:   "login successful",
		Tokens:    tokens,
	}
}
