package maze

type CellType int

const (
	Wall    CellType = 0
	Free    CellType = 1
	Visited CellType = 2
	Exit    CellType = 9
)

// Matrix represents the actual Maze as a matrix of numbers.
// Number 0 on the cell represents a Wall
// Number 1 represents Space (Free)
// Number 2 means that the cell has been visited already
// Number 9 represents an Exit. The Exit can only be in the last row
// of the matrix
type Matrix [][]CellType

// NewMatrix initialises a new Matrix of a given Size and sets the Walls on the matrix
func NewMatrix(size Size, walls []Cell) Matrix {
	m := make(Matrix, size.Rows)

	// set the rest of the matrix
	for r := 0; r < size.Rows-1; r++ {
		cols := make([]CellType, size.Columns)
		// set the rest of the columns to free
		for c := 0; c < size.Columns; c++ {
			cols[c] = Free
		}
		m[r] = cols
	}

	// set the exits for the maze (they can only be at in the bottom)
	cols := make([]CellType, size.Columns)
	// set the rest of the columns to free
	for c := 0; c < size.Columns; c++ {
		cols[c] = Exit
	}
	m[size.Rows-1] = cols

	// set the walls
	for _, w := range walls {
		m[w.Row][w.Col] = Wall
	}

	return m
}

// IsInside checks if the cell provided is inside the matrix or not
func (m Matrix) IsInside(c Cell) bool {
	return c.Row >= 0 && c.Row < len(m) && c.Col >= 0 && c.Col < len(m[0])
}

// Visit marks the cell as visited
func (m Matrix) Visit(c Cell) {
	if m[c.Row][c.Col] != Exit {
		m[c.Row][c.Col] = Visited
	}
}

// IsSolution checks if the cell provided is a solution to the maze. If the cell is on the edge of the
// matrix, and it is not Wall then the cell is the solution
func (m Matrix) IsSolution(c Cell) bool {
	return m[c.Row][c.Col] == Exit
}
