package responses

import "github.com/fitims/exercise/maze"

type MazeSolutionResponse struct {
	IsSuccess bool      `json:"is_success"`
	Message   string    `json:"message"`
	Solution  maze.Path `json:"solution,omitempty"`
}

func MazeSolutionFailed(msg string) MazeSolutionResponse {
	return MazeSolutionResponse{
		IsSuccess: false,
		Message:   msg,
	}
}

func MazeSolutionSuccessful(p maze.Path) MazeSolutionResponse {
	return MazeSolutionResponse{
		IsSuccess: true,
		Message:   "maze retrieved successfully",
		Solution:  p,
	}
}
