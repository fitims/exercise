package responses

type MazeRetrievalResponse struct {
	IsSuccess bool           `json:"is_success"`
	Message   string         `json:"message"`
	Mazes     []MazeResponse `json:"mazes,omitempty"`
}
