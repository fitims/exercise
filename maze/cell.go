package maze

import (
	"errors"
	"fmt"
)

var (
	colNames = []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}

	InvalidColumnNameErr = errors.New("invalid column name")
)

const (
	Letters = 26
)

// Cell represents a cell in a maze
type Cell struct {
	Row int
	Col int
}

// Up checks if we can move once cell up in the maze. A move is valid if the
// maze has a row above the current cell, and if the cell above is not a wall,
//and the cell hasn't been visited before.

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

// String returns a string representation fo the cell
func (c Cell) String() string {
	return fmt.Sprintf("%s%d", ConvertIntToColumn(c.Col), c.Row)
}

// ConvertIntToColumn converts an integer value to Column Name. Column names
// are in the format of:
// A-Z, AA-AZ, BA-BZ, CA-CZ, ... , AAA..AAZ, ABA..ABZ, ... , AAAA...
func ConvertIntToColumn(col int) string {
	if col == 0 {
		return colNames[col]
	}

	count := 0
	result := ""
	for col > 0 {
		mod := col % Letters
		div := col / Letters

		if count > 0 {
			if mod > 0 {
				result = colNames[mod-1] + result
			} else {
				result = colNames[mod] + result
			}
		} else {
			result = colNames[mod] + result
		}
		col = div
		count += 1
	}
	return result
}

// ConvertColumnToInt converts colName to integer value.
// If the conversion succeeds the integer value is returned, otherwise
// InvalidColumnErr is returned.
//
// Column names are in the format of:
// A-Z, AA-AZ, BA-BZ, CA-CZ, ... , AAA..AAZ, ABA..ABZ, ... , AAAA...
func ConvertColumnToInt(colName string) (int, error) {
	if len(colName) == 1 {
		idx := IndexOf(colNames, colName)
		if idx == -1 {
			return 0, InvalidColumnNameErr
		}
		return idx, nil
	}

	sum := 0
	for i := 0; i < len(colName); i++ {
		idx := IndexOf(colNames, string(colName[i]))
		if idx == -1 {
			return 0, InvalidColumnNameErr
		}

		if i == 0 {
			sum = sum*Letters + (idx + 1)
		} else {
			sum = sum*Letters + idx
		}
	}
	return sum, nil
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
