package maze

import (
	"log"
	"strconv"
	"strings"
)

// Size represents the grid size of the maze in Rows and Columns
type Size struct {
	Rows    int `json:"rows"`
	Columns int `json:"columns"`
}

// ParseGridSize parses the provided string into rows and columns. The valid format for
// gridSize string si in the format of [cols]x[rows] ie. 5x5, 8x8, 12x10, etc.
// If the string is not in valid format then InvalidGridSizeErr is returned.
//
// The minimum valid grid size is 1x1 and maximum valid grid size is 26x26. If the grid size is
// not in valid range then InvalidGridSizeErr is returned.
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

	result := Size{
		Rows:    rows,
		Columns: cols,
	}

	if !result.IsValid() {
		return Size{}, InvalidGridSizeErr
	}

	return result, nil
}

// IsValid check if the size is a valid grid size. A minimal valid grid Size
// can be a grid with 1 row and 1 column and the maximum valid Size for the grid
// can be a grid with 26 rows and 26 columns
func (s Size) IsValid() bool {
	return s.Rows >= MinRows && s.Rows <= MaxRows &&
		s.Columns >= MinCols && s.Columns <= MaxCols
}
