package responses

import "github.com/fitims/exercise/maze"

type MazeResponse struct {
	Id       uint64   `json:"id"`
	GridSize string   `json:"gridSize"`
	Entrance string   `json:"entrance"`
	Wall     []string `json:"wall"`
}

func ConvertToResponse(m maze.Maze) MazeResponse {
	wall := make([]string, 0)
	for _, v := range m.Walls {
		wall = append(wall, v.String())
	}

	return MazeResponse{
		Id:       m.Id,
		GridSize: m.GridSize.String(),
		Entrance: m.Entrance.String(),
		Wall:     wall,
	}
}
