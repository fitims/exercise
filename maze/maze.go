package maze

import (
	"errors"
	"log"
	"strconv"
	"strings"
)

var (
	InvalidGridSizeErr = errors.New("invalid grid size")
)

type Size struct {
	Rows    int `json:"rows"`
	Columns int `json:"columns"`
}

type Maze struct {
	Id       uint   `json:"id"`
	Entrance Cell   `json:"entrance"`
	Matrix   Matrix `json:"matrix"`
	GridSize Size   `json:"gridSize"`
	Walls    []Cell `json:"walls"`
}

func ParseWall(wall []string) ([]Cell, error) {
	return nil, nil
}

// ParseGridSize parses the provided string into rows and columns. The valid format for
// gridSize string si in the format of [cols]x[rows] ie. 5x5, 8x8, 12x10, etc.
//
// If the provided string has invalid columnName (A, B, ..., AA, AB, AAA, ...) then InvalidGridSizeErr is returned.
// if the provided string has invalid row number (1, 2, 3, 10, 25, 100, ...) then InvalidGridSizeErr is returned.
// Or, if the string is not in valid format then InvalidGridSizeErr is returned.
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
