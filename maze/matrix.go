package maze

type CellType int

const (
	Wall     CellType = 0
	Free     CellType = 1
	Visited  CellType = 2
	Solution CellType = 9
)

type Matrix [][]CellType

func (m Matrix) Visit(c Cell) {
	m[c.Row][c.Col] = Visited
}

func (m Matrix) IsSolution(c Cell) bool {
	return m[c.Row][c.Col] == Solution
}
