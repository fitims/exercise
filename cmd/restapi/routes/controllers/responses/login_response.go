package responses

import "github.com/fitims/exercise/cmd/restapi/middleware/authentication/tokens"

type LoginResponse struct {
	IsSuccess bool              `json:"is_success"`
	Message   string            `json:"message"`
	Tokens    *tokens.TokenPair `json:"tokens, omitempty"`
}

func LoginFailed(msg string) LoginResponse {
	return LoginResponse{
		IsSuccess: false,
		Message:   msg,
	}
}

func LoginSuccessful(tokens *tokens.TokenPair) LoginResponse {
	return LoginResponse{
		IsSuccess: true,
		Message:   "login successful",
		Tokens:    tokens,
	}
}
