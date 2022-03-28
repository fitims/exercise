package maze

import (
	"fmt"
	"testing"
)

func TestNewMatrix(t *testing.T) {
	size := Size{Rows: 4, Columns: 4}
	wall := []Cell{{Row: 0, Col: 1}, {Row: 1, Col: 1}, {Row: 2, Col: 1}, {Row: 3, Col: 1}}
	m := NewMatrix(size, wall)

	for r := 0; r < size.Rows; r++ {
		for c := 0; c < size.Columns; c++ {
			if r == 0 && c == 1 && m[r][c] != Wall {
				t.Error(fmt.Sprintf("Matrix is not initialised correctly. The cell [%d,%d] should be Wall (0), but it is %d", r, c, m[r][c]))
			}
		}
	}
}

func TestMatrix_Visit(t *testing.T) {
	size := Size{Rows: 4, Columns: 4}
	m := NewMatrix(size, []Cell{})

	if m[0][0] != Free {
		t.Error("Matrix is invalid. Expected: Free (1). Got: ", m[0][0])
	}

	m.Visit(Cell{Row: 0, Col: 0})

	if m[0][0] != Visited {
		t.Error("Matrix is invalid. Expected : Visited (2). Got: ", m[0][0])
	}
}

func TestMatrix_IsInside(t *testing.T) {
	testDate := []struct {
		cell     Cell
		isInside bool
	}{
		{Cell{Row: 0, Col: 0}, true},
		{Cell{Row: 0, Col: 1}, true},
		{Cell{Row: 0, Col: 2}, true},
		{Cell{Row: 0, Col: 3}, true},
		{Cell{Row: 1, Col: 0}, true},
		{Cell{Row: 1, Col: 1}, true},
		{Cell{Row: 1, Col: 2}, true},
		{Cell{Row: 1, Col: 3}, true},
		{Cell{Row: 2, Col: 0}, true},
		{Cell{Row: 2, Col: 1}, true},
		{Cell{Row: 2, Col: 2}, true},
		{Cell{Row: 2, Col: 3}, true},
		{Cell{Row: 3, Col: 0}, true},
		{Cell{Row: 3, Col: 1}, true},
		{Cell{Row: 3, Col: 2}, true},
		{Cell{Row: 3, Col: 3}, true},
		{Cell{Row: 4, Col: 0}, false},
		{Cell{Row: 0, Col: 4}, false},
		{Cell{Row: 4, Col: 4}, false},
	}

	size := Size{Rows: 4, Columns: 4}
	m := NewMatrix(size, []Cell{})

	for _, v := range testDate {
		if m.IsInside(v.cell) != v.isInside {
			t.Error(fmt.Sprintf("Expected : %v, GOT: %v", v.isInside, m.IsInside(v.cell)))
		}
	}
}

func TestMatrix_IsSolution(t *testing.T) {
	size := Size{Rows: 4, Columns: 4}
	wall := []Cell{{Row: 0, Col: 1}, {Row: 1, Col: 1}, {Row: 2, Col: 1}, {Row: 3, Col: 1}}
	m := NewMatrix(size, wall)

	testDate := []struct {
		cell       Cell
		isSolution bool
	}{
		{Cell{Row: 0, Col: 0}, true},
		{Cell{Row: 0, Col: 1}, false},
		{Cell{Row: 0, Col: 2}, true},
		{Cell{Row: 0, Col: 3}, true},
		{Cell{Row: 1, Col: 0}, true},
		{Cell{Row: 1, Col: 1}, false},
		{Cell{Row: 1, Col: 2}, false},
		{Cell{Row: 1, Col: 3}, true},
		{Cell{Row: 2, Col: 0}, true},
		{Cell{Row: 2, Col: 1}, false},
		{Cell{Row: 2, Col: 2}, false},
		{Cell{Row: 2, Col: 3}, true},
		{Cell{Row: 3, Col: 0}, true},
		{Cell{Row: 3, Col: 1}, false},
		{Cell{Row: 3, Col: 2}, true},
		{Cell{Row: 3, Col: 3}, true},
	}

	for _, v := range testDate {
		if m.IsSolution(v.cell) != v.isSolution {
			t.Error(fmt.Sprintf("Cell: %s, Expected : %v, GOT: %v", v.cell.String(), v.isSolution, m.IsSolution(v.cell)))
		}
	}
}
