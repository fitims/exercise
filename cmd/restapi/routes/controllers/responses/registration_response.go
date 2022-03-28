package responses

type RegistrationResponse struct {
	IsSuccess bool   `json:"is_success"`
	Message   string `json:"message"`
}

func RegistrationFailed(msg string) RegistrationResponse {
	return RegistrationResponse{
		IsSuccess: false,
		Message:   msg,
	}
}

func RegistrationSuccessful() RegistrationResponse {
	return RegistrationResponse{
		IsSuccess: true,
		Message:   "User registered successfully",
	}
}
