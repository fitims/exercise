package maze

type Path []Cell

// GetExitCell returns the last cell in the path
func (p Path) GetExitCell() Cell {
	return p[len(p)-1]
}

// ToString returns string representation fo the cells in the path
func (p Path) ToString() []string {
	cells := make([]string, 0)
	for _, v := range p {
		cells = append(cells, v.String())
	}

	return cells
}
