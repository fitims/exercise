package maze

type CellType int

const (
	Wall     CellType = 0
	Free     CellType = 1
	Visited  CellType = 2
	Solution CellType = 9
)

// Matrix represents the actual Maze as a matrix of numbers.
// Number 0 on the cell represents a Wall
// Number 1 represents Space (Free)
// Number 2 means that the cell has been visited already
// Number 9 represents a Solution. The Solution can only be in the last row
// of the matrix
type Matrix [][]CellType

// NewMatrix initialises a new Matrix of a given Size and sets the Walls on the matrix
func NewMatrix(size Size, walls []Cell) Matrix {
	m := make(Matrix, size.Rows)
	for r := 0; r < size.Rows-1; r++ {
		cols := make([]CellType, size.Columns)
		for c := 0; c < size.Columns; c++ {
			cols[c] = Free
		}
		m[r] = cols
	}

	// add the last row
	cols := make([]CellType, size.Columns)
	for c := 0; c < size.Columns; c++ {
		cols[c] = Solution
	}
	m[size.Rows-1] = cols

	// set the walls
	for _, w := range walls {
		m[w.Row][w.Col] = Wall
	}

	return m
}

func (m Matrix) Visit(c Cell) {
	m[c.Row][c.Col] = Visited
}

func (m Matrix) IsSolution(c Cell) bool {
	return m[c.Row][c.Col] == Solution
}
