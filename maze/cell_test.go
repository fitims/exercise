package maze

import (
	"fmt"
	"testing"
)

func TestCell_Up(t *testing.T) {
	matrix := Matrix{
		{Free, Wall, Wall},
		{Free, Free, Visited},
		{Wall, Free, Free},
		{Wall, Free, Wall},
	}

	// test for the top of the maze
	c := Cell{Row: 0, Col: 0}
	_, success := c.Up(matrix)

	if success {
		t.Error("Cannot move Up. Cell is at the top of the maze")
	}

	// test for the wall
	c = Cell{Row: 1, Col: 1}
	_, success = c.Up(matrix)

	if success {
		t.Error("Cannot move Up. Cell up is a wall")
	}

	// test for visited
	c = Cell{Row: 2, Col: 2}
	_, success = c.Up(matrix)

	if success {
		t.Error("Cannot move Up. Cell up has been visited before")
	}

	// test for the free cell
	c = Cell{Row: 1, Col: 0}
	up, success := c.Up(matrix)

	if !success {
		t.Error("Move up should be possible as the cell up is free")
	}

	if up.Row != 0 && up.Col != 0 {
		t.Error(fmt.Sprintf("Invalid cell returned. Expected: (r:0, c:0), GOT: (r:%d, c:%d)", up.Row, up.Col))
	}
}

func TestCell_Down(t *testing.T) {
	matrix := Matrix{
		{Free, Wall, Free},
		{Free, Free, Visited},
		{Wall, Free, Free},
		{Wall, Free, Wall},
	}

	// test for the bottom of the maze
	c := Cell{Row: 3, Col: 1}
	_, success := c.Down(matrix)

	if success {
		t.Error("Cannot move Down. Cell is at the bottom of the maze")
	}

	// test for the wall
	c = Cell{Row: 1, Col: 0}
	_, success = c.Down(matrix)

	if success {
		t.Error("Cannot move Down. Cell down is a wall")
	}

	// test for visited
	c = Cell{Row: 0, Col: 2}
	_, success = c.Down(matrix)

	if success {
		t.Error("Cannot move Down. Cell down has been visited before")
	}

	// test for the free cell
	c = Cell{Row: 0, Col: 0}
	up, success := c.Down(matrix)

	if !success {
		t.Error("Move up should be possible as the cell down is free")
	}

	if up.Row != 1 && up.Col != 0 {
		t.Error(fmt.Sprintf("Invalid cell returned. Expected: (r:1, c:0), GOT: (r:%d, c:%d)", up.Row, up.Col))
	}
}

func TestCell_Left(t *testing.T) {
	matrix := Matrix{
		{Free, Wall, Free},
		{Free, Visited, Free},
		{Wall, Free, Free},
		{Free, Free, Wall},
		{Wall, Free, Wall},
	}

	// test for the left edge of the maze
	c := Cell{Row: 1, Col: 0}
	_, success := c.Left(matrix)

	if success {
		t.Error("Cannot move Left. Cell is at the left edge of the maze")
	}

	// test for the wall
	c = Cell{Row: 2, Col: 1}
	_, success = c.Left(matrix)

	if success {
		t.Error("Cannot move Left. Cell to the left is a wall")
	}

	// test for visited
	c = Cell{Row: 1, Col: 2}
	_, success = c.Left(matrix)

	if success {
		t.Error("Cannot move Left. Cell to the left has been visited before")
	}

	// test for the free cell
	c = Cell{Row: 3, Col: 1}
	up, success := c.Left(matrix)

	if !success {
		t.Error("Move to the left should be possible as the cell to the left is free")
	}

	if up.Row != 3 && up.Col != 0 {
		t.Error(fmt.Sprintf("Invalid cell returned. Expected: (r:3, c:0), GOT: (r:%d, c:%d)", up.Row, up.Col))
	}
}

func TestCell_Right(t *testing.T) {
	matrix := Matrix{
		{Free, Wall, Free},
		{Free, Visited, Free},
		{Wall, Free, Free},
		{Free, Free, Wall},
		{Wall, Free, Wall},
	}

	// test for the right edge of the maze
	c := Cell{Row: 0, Col: 2}
	_, success := c.Right(matrix)

	if success {
		t.Error("Cannot move Right. Cell is at the Right edge of the maze")
	}

	// test for the wall
	c = Cell{Row: 3, Col: 1}
	_, success = c.Right(matrix)

	if success {
		t.Error("Cannot move Right. Cell to the Right is a wall")
	}

	// test for visited
	c = Cell{Row: 1, Col: 0}
	_, success = c.Right(matrix)

	if success {
		t.Error("Cannot move Right. Cell to the Right has been visited before")
	}

	// test for the free cell
	c = Cell{Row: 2, Col: 1}
	up, success := c.Right(matrix)

	if !success {
		t.Error("Move to the Right should be possible as the cell to the Right is free")
	}

	if up.Row != 2 && up.Col != 2 {
		t.Error(fmt.Sprintf("Invalid cell returned. Expected: (r:2, c:2), GOT: (r:%d, c:%d)", up.Row, up.Col))
	}
}

func TestParseCell(t *testing.T) {
	testData := []struct {
		cellStr       string
		expectedCell  Cell
		expectedError error
	}{
		{cellStr: "A1", expectedCell: Cell{Row: 0, Col: 0}, expectedError: nil},
		{cellStr: "B1", expectedCell: Cell{Row: 0, Col: 1}, expectedError: nil},
		{cellStr: "C3", expectedCell: Cell{Row: 2, Col: 2}, expectedError: nil},
		{cellStr: "Z20", expectedCell: Cell{Row: 19, Col: 25}, expectedError: nil},
		{cellStr: "G45", expectedCell: Cell{}, expectedError: InvalidRowErr},
		{cellStr: "AA2", expectedCell: Cell{}, expectedError: InvalidRowErr},
		{cellStr: "#5", expectedCell: Cell{}, expectedError: InvalidColumnNameErr},
		{cellStr: "A", expectedCell: Cell{}, expectedError: InvalidRowErr},
	}

	for _, v := range testData {
		cell, err := ParseCell(v.cellStr)

		if err != v.expectedError {
			t.Error(fmt.Sprintf("Error. Expected: %v, GOT: %v", v.expectedError, err))
		}

		if !cell.IsSame(v.expectedCell) {
			t.Error(fmt.Sprintf("Cells should be same. Expected: (c:%d, r:%d), GOT: (c:%d, r:%d)", v.expectedCell.Row, v.expectedCell.Col, cell.Col, cell.Row))
		}
	}
}
