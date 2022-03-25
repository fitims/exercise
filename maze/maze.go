package maze

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
)

var (
	InvalidGridSizeErr = errors.New("invalid grid size")
)

// Maze encapsulates all the details about the maze
type Maze struct {
	Id        uint   `json:"id"`
	Entrance  Cell   `json:"entrance"`
	Matrix    Matrix `json:"matrix"`
	GridSize  Size   `json:"gridSize"`
	Walls     []Cell `json:"walls"`
	IsSolved  bool   `json:"isSolved"`
	Solutions []Path `json:"solutions,omitempty"`
}

// NewMaze initialises a new Maze by creating the actual maze as a matrix of the size provided,
// and storing the entrance of the maze and the walls
func NewMaze(id uint, entrance, gridSize string, wall []string) (*Maze, error) {
	e, err := ParseCell(entrance)
	if err != nil {
		log.Println("Could not parse entrance. Error: ", err)
		return nil, err
	}

	s, err := ParseGridSize(gridSize)
	if err != nil {
		log.Println("Could not parse grid Size. Error: ", err)
		return nil, err
	}

	// if the grid size is not valid return an error
	if !s.IsValid() {
		log.Println("The grid size is too small")
		return nil, InvalidGridSizeErr
	}

	w, err := ParseWall(wall)
	if err != nil {
		log.Println("Could not parse the wall. Error: ", err)
		return nil, err
	}

	maze := &Maze{
		Id:        id,
		Entrance:  e,
		GridSize:  s,
		Matrix:    NewMatrix(s, w),
		IsSolved:  false,
		Walls:     w,
		Solutions: make([]Path, 0),
	}
	return maze, nil
}

// Solve will try to solve the maze by walking through the maze and finding all possible solutions.
func (m *Maze) Solve() ([]Path, bool) {
	path := make(Path, 0)
	path = append(path, m.Entrance)

	// walk through the maze and find all the possible solutions
	m.walkTheMaze(m.Entrance, path)

	// mark the maze as solved
	m.IsSolved = true
	return m.Solutions, len(m.Solutions) > 0
}

// walkTheMaze is a recursive functions that walks through the maze finding all the possible solutions.
// For each possible cell movement (Up, Down, Left, Right) it calls itself passing the adjacent cell and
// the path travelled so far
func (m *Maze) walkTheMaze(currentCell Cell, currentPath Path) {
	p := currentPath
	m.Matrix.Visit(currentCell)

	// move up
	if c, ok := currentCell.Up(m.Matrix); ok {
		p = append(p, c)
		if m.Matrix.IsSolution(c) {
			m.Solutions = append(m.Solutions, p)
			return
		}
		m.walkTheMaze(c, p)
	}

	// move down
	if c, ok := currentCell.Down(m.Matrix); ok {
		p = append(p, c)
		if m.Matrix.IsSolution(c) {
			m.Solutions = append(m.Solutions, p)
			return
		}
		m.walkTheMaze(c, p)
	}

	// move left
	if c, ok := currentCell.Left(m.Matrix); ok {
		p = append(p, c)
		if m.Matrix.IsSolution(c) {
			m.Solutions = append(m.Solutions, p)
			return
		}
		m.walkTheMaze(c, p)
	}

	// move right
	if c, ok := currentCell.Right(m.Matrix); ok {
		p = append(p, c)
		if m.Matrix.IsSolution(c) {
			m.Solutions = append(m.Solutions, p)
			return
		}
		m.walkTheMaze(c, p)
	}
}

// ParseWall parses an array of string into an array of valid cells
// If any of the provided strings in not a valid Cell, then an error is returned
func ParseWall(wall []string) ([]Cell, error) {
	wallCells := make([]Cell, 0)
	for _, v := range wall {
		cell, err := ParseCell(v)
		if err != nil {
			log.Println(fmt.Sprintf("Error parsing the wall cell [%s]. Error: %s", v, err.Error()))
			return nil, err
		}
		wallCells = append(wallCells, cell)
	}
	return wallCells, nil
}

// ParseGridSize parses the provided string into rows and columns. The valid format for
// gridSize string si in the format of [cols]x[rows] ie. 5x5, 8x8, 12x10, etc.
// If the string is not in valid format then InvalidGridSizeErr is returned.
func ParseGridSize(gridSize string) (Size, error) {
	parts := strings.Split(gridSize, "x")
	if len(parts) != 2 {
		return Size{}, InvalidGridSizeErr
	}

	rows, err := strconv.Atoi(parts[0])
	if err != nil {
		log.Println("Error converting string to integer. Error: ", err)
		return Size{}, InvalidGridSizeErr
	}

	cols, err := strconv.Atoi(parts[1])
	if err != nil {
		log.Println("Error converting string to integer. Error: ", err)
		return Size{}, InvalidGridSizeErr
	}

	return Size{
		Rows:    rows,
		Columns: cols,
	}, nil
}
