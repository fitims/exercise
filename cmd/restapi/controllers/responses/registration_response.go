package responses

// RegistrationResponse is the response sent back when user tries to register.
// The registration can be successful or it can fail.
type RegistrationResponse struct {
	IsSuccess bool   `json:"is_success"`
	Message   string `json:"message"`
}

// RegistrationFailed build a failed response
func RegistrationFailed(msg string) RegistrationResponse {
	return RegistrationResponse{
		IsSuccess: false,
		Message:   msg,
	}
}

// RegistrationSuccessful builds a successful response
func RegistrationSuccessful() RegistrationResponse {
	return RegistrationResponse{
		IsSuccess: true,
		Message:   "User registered successfully",
	}
}
