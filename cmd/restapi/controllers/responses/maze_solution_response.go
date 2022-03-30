package responses

import "github.com/fitims/exercise/maze"

type MazeSolutionResponse struct {
	IsSuccess bool   `json:"is_success"`
	Message   string `json:"message"`
}

type MazeSolutionWithPathResponse struct {
	Path []string `json:"path"`
}

func MazeSolutionFailed(msg string) MazeSolutionResponse {
	return MazeSolutionResponse{
		IsSuccess: false,
		Message:   msg,
	}
}

func MazeSolutionSuccessful(p maze.Path) MazeSolutionWithPathResponse {
	return MazeSolutionWithPathResponse{
		Path: p.ToString(),
	}
}
