package maze

import (
	"fmt"
	"testing"
)

func TestNewMaze(t *testing.T) {
	testData := []struct {
		entrance         string
		size             string
		walls            []string
		expectedEntrance Cell
		expectedSize     Size
		expectedError    error
	}{
		{"A1", "8x8", []string{"C1", "G1", "A2", "C2", "E2"}, Cell{Row: 0, Col: 0}, Size{Rows: 8, Columns: 8}, nil},
		{"AA", "8x8", []string{"C1", "G1", "A2", "C2", "E2"}, Cell{Row: 0, Col: 0}, Size{Rows: 8, Columns: 8}, InvalidRowErr},
		{"12", "8x8", []string{"C1", "G1", "A2", "C2", "E2"}, Cell{Row: 0, Col: 0}, Size{Rows: 8, Columns: 8}, InvalidColumnNameErr},
		{"B2", "0x8", []string{"C1", "G1", "A2", "C2", "E2"}, Cell{Row: 0, Col: 0}, Size{Rows: 8, Columns: 8}, InvalidGridSizeErr},
		{"B2", "8x0", []string{"C1", "G1", "A2", "C2", "E2"}, Cell{Row: 0, Col: 0}, Size{Rows: 8, Columns: 8}, InvalidGridSizeErr},
		{"B2", "27x27", []string{"C1", "G1", "A2", "C2", "E2"}, Cell{Row: 0, Col: 0}, Size{Rows: 8, Columns: 8}, InvalidGridSizeErr},
		{"B2", "8x8", []string{"C!", "G1", "A2", "C2", "E2"}, Cell{Row: 0, Col: 0}, Size{Rows: 8, Columns: 8}, InvalidRowErr},
		{"B2", "8x8", []string{"1A", "G1", "A2", "C2", "E2"}, Cell{Row: 0, Col: 0}, Size{Rows: 8, Columns: 8}, InvalidColumnNameErr},
	}

	for _, v := range testData {
		m, err := NewMaze(1, v.entrance, v.size, v.walls)

		if v.expectedError != nil {
			if err != v.expectedError {
				t.Error(fmt.Sprintf("Expeced: %v, Got: %v", v.expectedError, err))
			}
			continue
		}

		if m.State != NotSolved {
			t.Error(fmt.Sprintf("Invalid State. Expeced: NotSolved (0), Got: %d", m.State))
			continue
		}

		if !m.Entrance.IsSame(v.expectedEntrance) {
			t.Error(fmt.Sprintf("Invalid enterance. Expeced: %s, Got: %s", v.expectedEntrance.String(), m.Entrance.String()))
			continue
		}

		if !m.GridSize.IsSame(v.expectedSize) {
			t.Error(fmt.Sprintf("Invalid Size. Expeced: %s, Got: %s", v.expectedSize.String(), m.GridSize.String()))
			continue
		}

	}
}

func TestParseWall(t *testing.T) {
	testData := []struct {
		wall          []string
		expectedError error
	}{
		{[]string{"A1", "B1", "C1"}, nil},
		{[]string{"AA", "B1", "C1"}, InvalidRowErr},
		{[]string{"3A", "B1", "C1"}, InvalidColumnNameErr},
		{[]string{"a1", "B1", "C1"}, InvalidColumnNameErr},
	}

	for _, v := range testData {

		// we do not need to test the wall cells as the ParseWall uses ParseCell to parse each individual cell
		// of the wall, and ParseCell is covered in unit tests
		_, err := ParseWall(v.wall)

		if err != v.expectedError {
			t.Error(fmt.Sprintf("Expeced: %v, Got: %v", v.expectedError, err))
		}
	}
}