package requests

// LoginRequest encapsulates the request for login
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
