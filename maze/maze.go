package maze

import (
	"errors"
	"fmt"
	"github.com/fitims/exercise/log"
)

type MazeState int
type PathPredicate func(Path, Path) bool

const (
	NotSolved        MazeState = 0
	NoSolutions      MazeState = 1
	TooManySolutions MazeState = 2
	Solved           MazeState = 3
)

// Maze encapsulates all the details about the maze
type Maze struct {
	Id        uint64    `json:"id"`
	Entrance  Cell      `json:"entrance"`
	Matrix    Matrix    `json:"matrix"`
	GridSize  Size      `json:"gridSize"`
	Walls     []Cell    `json:"walls"`
	State     MazeState `json:"state"`
	Solutions []Path    `json:"solutions,omitempty"`
}

var (
	NoSolutionErr    = errors.New("maze does not have a solution")
	ManySolutionsErr = errors.New("maze has more than one solution")
)

// NewMaze initialises a new Maze by creating the actual maze as a matrix of the size provided,
// and storing the entrance of the maze and the walls
func NewMaze(id uint64, entrance, gridSize string, wall []string) (*Maze, error) {
	e, err := ParseCell(entrance)
	if err != nil {
		log.Errorln("Could not parse entrance. Error: ", err)
		return nil, err
	}

	s, err := ParseGridSize(gridSize)
	if err != nil {
		log.Errorln("Could not parse grid Size. Error: ", err)
		return nil, err
	}

	// if the grid size is not valid return an error
	if !s.IsValid() {
		log.Errorln("The grid size is too small")
		return nil, InvalidGridSizeErr
	}

	w, err := ParseWall(wall)
	if err != nil {
		log.Errorln("Could not parse the wall. Error: ", err)
		return nil, err
	}

	maze := &Maze{
		Id:        id,
		Entrance:  e,
		GridSize:  s,
		Matrix:    NewMatrix(s, w),
		State:     NotSolved,
		Walls:     w,
		Solutions: make([]Path, 0),
	}
	return maze, nil
}

// Solve will try to solve the maze by walking through the maze and finding all possible solutions if
// the maze is not already solved. If it is solved then checks for the solutions.
// If the maze has no solutions it returns NoSolutions error, if more than one different solution
// then it returns TooManySolutions error. If it only one solution or more than one but with the same
// exit then the maze is valid and no error is returned.
//
// The assumption is that the solution can only be at the bottom of the maze.
//
// So we start from the top to bottom.
// ie. imagine that we have a maze defined with the following information:
//
// entrance  : A1
// grid size : 8x8
// walls     : [C1,G1,A2,C2,E2,G2,C3,E3,B4,C4,E4,F4,G4,B5,E5,B6,D6,E6,G6,H6,B7,D7,G7,B8]
//
// the maze would look like:
//
//    A B C D E F G H
//  1 |*| |█| | | |█| |
//  2 |█| |█| |█|█| | |
//  3 | | |█| |█| | | |
//  4 | |█|█| |█|█|█| |
//  5 | |█| | |█| | | |
//  6 | |█| |█|█| |█|█|
//  7 | |█| |█| | |█| |
//  8 | |█| | | | | | |
//
//  so the solution will be:
//
//    A B C D E F G H
//  1 |░|░|█| | | |█| |
//  2 |█|░|█| |█|█| | |
//  3 |░|░|█| |█| | | |
//  4 |░|█|█| |█|█|█| |
//  5 |░|█| | |█| | | |
//  6 |░|█| |█|█| |█|█|
//  7 |░|█| |█| | |█| |
//  8 |░|█| | | | | | |
//
// which translates to : [A1, B1, B2, B3, A3, A4, A5, A6, A7, A8]
//
func (m *Maze) Solve() error {

	if m.State == NotSolved {
		path := make(Path, 0)
		path = append(path, m.Entrance)

		// walk through the maze and find all the possible solutions
		m.walkTheMaze(m.Entrance, path)
	}

	if len(m.Solutions) == 0 {
		m.State = NoSolutions
		return NoSolutionErr
	}

	if len(m.Solutions) > 1 {
		sol := m.Solutions[0].GetExitCell()

		for i := 1; i < len(m.Solutions); i++ {
			if !sol.IsSame(m.Solutions[i].GetExitCell()) {
				m.State = TooManySolutions
				return ManySolutionsErr
			}
		}
	}
	m.State = Solved
	return nil
}

// GetShortestPath return the shortest solution path for the maze. If the maze
// does not have any solutions then NoSolutionErr error is returned. If
// the maze has more than one different solutions then  TooManySolutions error is returned
func (m *Maze) GetShortestPath() (Path, error) {
	return m.GetPath(func(p1, p2 Path) bool {
		return len(p1) > len(p2)
	})
}

// GetLongestPath return the longest solution path for the maze. If the maze
// does not have any solutions then NoSolutionErr error is returned. If
// the maze has more than one different solutions then  TooManySolutions error is returned
func (m *Maze) GetLongestPath() (Path, error) {
	return m.GetPath(func(p1, p2 Path) bool {
		return len(p1) < len(p2)
	})
}

// GetPath return the solution path for the maze. If the maze
// does not have any solutions then NoSolutionErr error is returned. If
// the maze has more than one different solutions then  TooManySolutions error is returned
// If the maze is not solved, it first is solved.
func (m *Maze) GetPath(predicate PathPredicate) (Path, error) {
	if m.State == NotSolved {
		_ = m.Solve()
	}

	if m.State == TooManySolutions {
		return nil, ManySolutionsErr
	}

	if m.State == NoSolutions {
		return nil, NoSolutionErr
	}

	path := m.Solutions[0]
	for i := 1; i < len(m.Solutions); i++ {
		if predicate(path, m.Solutions[i]) {
			path = m.Solutions[i]
		}
	}
	return path, nil
}

// walkTheMaze is a recursive functions that walks through the maze finding all the possible solutions.
// For each possible cell movement (Up, Down, Left, Right) it calls itself passing the adjacent cell and
// the path travelled so far
func (m *Maze) walkTheMaze(currentCell Cell, currentPath Path) {
	m.Matrix.Visit(currentCell)

	if m.Matrix.IsSolution(currentCell) {
		m.Solutions = append(m.Solutions, currentPath)
		return
	}

	// move up
	if c, ok := currentCell.Up(m.Matrix); ok {
		p := append(currentPath, c)
		m.walkTheMaze(c, p)
	}

	// move down
	if c, ok := currentCell.Down(m.Matrix); ok {
		p := append(currentPath, c)
		m.walkTheMaze(c, p)
	}

	// move left
	if c, ok := currentCell.Left(m.Matrix); ok {
		p := append(currentPath, c)
		m.walkTheMaze(c, p)
	}

	// move right
	if c, ok := currentCell.Right(m.Matrix); ok {
		p := append(currentPath, c)
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
			log.Errorln(fmt.Sprintf("Error parsing the wall cell [%s]. Error: %s", v, err.Error()))
			return nil, err
		}
		wallCells = append(wallCells, cell)
	}
	return wallCells, nil
}
