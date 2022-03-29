package maze

type Path []Cell

// GetExitCell returns the last cell in the path
func (p Path) GetExitCell() Cell {
	return p[len(p)-1]
}
