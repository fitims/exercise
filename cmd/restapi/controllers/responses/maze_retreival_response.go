package responses

import "github.com/fitims/exercise/maze"

type MazeRetrievalResponse struct {
	IsSuccess bool           `json:"is_success"`
	Message   string         `json:"message"`
	Mazes     []MazeResponse `json:"mazes,omitempty"`
}

func MazeRetrievalFailed(msg string) MazeRetrievalResponse {
	return MazeRetrievalResponse{
		IsSuccess: false,
		Message:   msg,
	}
}

func MazeRetrievalSuccessful(mazes []maze.Maze) MazeRetrievalResponse {
	mr := make([]MazeResponse, 0)
	for _, v := range mazes {
		mr = append(mr, ConvertToResponse(v))
	}

	return MazeRetrievalResponse{
		IsSuccess: true,
		Message:   "",
		Mazes:     mr,
	}
}
