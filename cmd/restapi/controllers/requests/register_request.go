package requests

// RegisterRequest contains all the details necessary for registering the user
type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
