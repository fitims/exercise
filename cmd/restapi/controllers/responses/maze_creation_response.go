package responses

type MazeCreationResponse struct {
	IsSuccess bool   `json:"is_success"`
	Message   string `json:"message"`
}

// MazeCreationFailed builds a failed response
func MazeCreationFailed(msg string) MazeCreationResponse {
	return MazeCreationResponse{
		IsSuccess: false,
		Message:   msg,
	}
}

// MazeCreationSuccessful builds a successful response
func MazeCreationSuccessful() MazeCreationResponse {
	return MazeCreationResponse{
		IsSuccess: true,
		Message:   "Maze saved successfully",
	}
}
