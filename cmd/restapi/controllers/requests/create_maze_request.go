package requests

type CreateMazeRequest struct {
	Entrance string   `json:"entrance"`
	GridSize string   `json:"gridSize"`
	Walls    []string `json:"walls"`
}
