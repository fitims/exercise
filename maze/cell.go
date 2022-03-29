package maze

import (
	"errors"
	"fmt"
	"log"
	"strconv"
)

const (
	MinRows = 1
	MaxRows = 26

	MinCols = 1
	MaxCols = 26
)

var (
	colNames = []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}

	InvalidColumnNameErr = errors.New("invalid column name. Columns can be A-Z")
	InvalidRowErr        = errors.New("invalid row number. Rows can be 1-26")
	InvalidCellErr       = errors.New("invalid cell. Valid cells can be from A1-Z26")
)

// Cell represents a cell in a maze
type Cell struct {
	Row int `json:"row"`
	Col int `json:"col"`
}

// Up checks if we can move once cell up in the maze. A move is valid if the
// maze has a row above the current cell, and if the cell above is not a wall,
//and the cell hasn't been visited before.
//
// If the move up is valid, then the cell above is returned along with a true value
// specifying that the move is valid
func (c Cell) Up(maze Matrix) (Cell, bool) {
	if c.Row-1 >= 0 && maze[c.Row-1][c.Col] != Visited && maze[c.Row-1][c.Col] != Wall {
		return Cell{
			Row: c.Row - 1,
			Col: c.Col,
		}, true
	}

	return Cell{}, false
}

// Down checks if we can move once cell down in the maze. A move is valid if the
// maze has a row blow the current cell, and if the cell below is not a wall,
//and the cell hasn't been visited before.
//
// If the move down is valid, then the cell below is returned along with a true value
// specifying that the move is valid
func (c Cell) Down(maze Matrix) (Cell, bool) {
	if c.Row+1 < len(maze) && maze[c.Row+1][c.Col] != Visited && maze[c.Row+1][c.Col] != Wall {
		return Cell{
			Row: c.Row + 1,
			Col: c.Col,
		}, true
	}

	return Cell{}, false
}

// Left checks if we can move once cell to the left in the maze. A move is valid if the
// maze has a column to the left of the current cell, and if the cell to the left is not a wall,
//and the cell hasn't been visited before.
//
// If the move to the left  is valid, then the cell to the left is returned along with a true value
// specifying that the move is valid
func (c Cell) Left(maze Matrix) (Cell, bool) {
	if c.Col-1 >= 0 && maze[c.Row][c.Col-1] != Visited && maze[c.Row][c.Col-1] != Wall {
		return Cell{
			Row: c.Row,
			Col: c.Col - 1,
		}, true
	}

	return Cell{}, false
}

// Right checks if we can move once cell to the right in the maze. A move is valid if the
// maze has a column to the right of the current cell, and if the cell to the right is not a wall,
//and the cell hasn't been visited before.
//
// If the move to the right  is valid, then the cell to the right is returned along with a true value
// specifying that the move is valid
func (c Cell) Right(maze Matrix) (Cell, bool) {
	if c.Col+1 < len(maze[0]) && maze[c.Row][c.Col+1] != Visited && maze[c.Row][c.Col+1] != Wall {
		return Cell{
			Row: c.Row,
			Col: c.Col + 1,
		}, true
	}

	return Cell{}, false
}

// IsSame compares the cell with the provided one. If the Row and Col are the same
// then the cells are the same
func (c Cell) IsSame(cell Cell) bool {
	return c.Row == cell.Row && c.Col == cell.Col
}

// String returns a string representation fo the cell
func (c Cell) String() string {
	return fmt.Sprintf("%s%d", colNames[c.Col], c.Row)
}

// IndexOf returns the index of the string in the slice of strings.
// If the string is not in the slice then -1 is returned
func IndexOf(slice []string, item string) int {
	for i, v := range slice {
		if v == item {
			return i
		}
	}
	return -1
}

// ParseCell parses the provided string to a valid Cell.
//
// A valid cell is: A1, B5, Z30, etc. Ie. The Columns can be capital letters A-Z, and
// rows can be valid numbers 1-26.
//
// If the provided string does not have a column or the columnName is invalid, then
// InvalidColumnNameErr is returned.
//
// If the provided string does not have a row, the row is not valid number, or the number is
// less than 1 and greater than 26 then InvalidRowErr is returned.
//
// If the string provided does not have neither columnName nor row number then
// InvalidCellErr is returned.
func ParseCell(cell string) (Cell, error) {
	colName := string(cell[0])
	row := cell[1:]

	colNo := IndexOf(colNames, colName)
	if colNo < 0 {
		return Cell{}, InvalidColumnNameErr
	}

	rowNo, err := strconv.Atoi(row)
	if err != nil {
		log.Println("Row is not valid number. Error: ", err)
		return Cell{}, InvalidRowErr
	}

	if rowNo < MinRows || rowNo > MaxRows {
		log.Println("Row number has to be 1-26")
		return Cell{}, InvalidRowErr
	}

	return Cell{
		Row: rowNo - 1,
		Col: colNo,
	}, nil
}
