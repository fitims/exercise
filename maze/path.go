package maze

type Path []Cell

func (p Path) GetExitCell() Cell {
	return p[len(p)-1]
}
