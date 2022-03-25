package maze

// Size represents the grid size of the maze in Rows and Columns
type Size struct {
	Rows    int `json:"rows"`
	Columns int `json:"columns"`
}

// IsValid check if the size is a valid grid size. A minimal valid grid Size
// can be a grid with 1 row and 1 column
func (s Size) IsValid() bool {
	return s.Rows > 0 && s.Columns > 0
}
