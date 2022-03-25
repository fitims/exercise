package maze

import (
	"fmt"
	"testing"
)

func TestConvertIntToColumn(t *testing.T) {
	testData := []struct {
		ColNo   int
		ColName string
	}{
		{ColNo: 0, ColName: "A"}, {ColNo: 1, ColName: "B"}, {ColNo: 2, ColName: "C"}, {ColNo: 3, ColName: "D"}, {ColNo: 4, ColName: "E"},
		{ColNo: 5, ColName: "F"}, {ColNo: 6, ColName: "G"}, {ColNo: 7, ColName: "H"}, {ColNo: 8, ColName: "I"}, {ColNo: 9, ColName: "J"},
		{ColNo: 10, ColName: "K"}, {ColNo: 11, ColName: "L"}, {ColNo: 12, ColName: "M"}, {ColNo: 13, ColName: "N"}, {ColNo: 14, ColName: "O"},
		{ColNo: 15, ColName: "P"}, {ColNo: 16, ColName: "Q"}, {ColNo: 17, ColName: "R"}, {ColNo: 18, ColName: "S"}, {ColNo: 19, ColName: "T"},
		{ColNo: 20, ColName: "U"}, {ColNo: 21, ColName: "V"}, {ColNo: 22, ColName: "W"}, {ColNo: 23, ColName: "X"}, {ColNo: 24, ColName: "Y"},
		{ColNo: 25, ColName: "Z"},

		{ColNo: 26, ColName: "AA"},
		{ColNo: 27, ColName: "AB"},
		{ColNo: 26 * 2, ColName: "BA"},
		{ColNo: 26*2 + 5, ColName: "BF"},
		{ColNo: 26 * 3, ColName: "CA"},
		{ColNo: 26 * 26, ColName: "AAA"},
		{ColNo: 26 * 26 * 26, ColName: "AAAA"},
	}

	for _, v := range testData {

		expected := v.ColName
		actual := ConvertIntToColumn(v.ColNo)

		if actual != expected {
			t.Error(fmt.Sprintf("Error. GOT : %s, EXPECTED: %s", actual, expected))
		}
	}
}

func TestConvertColumnToInt(t *testing.T) {
	testData := []struct {
		ColNo   int
		ColName string
	}{
		{ColNo: 0, ColName: "A"}, {ColNo: 1, ColName: "B"}, {ColNo: 2, ColName: "C"}, {ColNo: 3, ColName: "D"}, {ColNo: 4, ColName: "E"},
		{ColNo: 5, ColName: "F"}, {ColNo: 6, ColName: "G"}, {ColNo: 7, ColName: "H"}, {ColNo: 8, ColName: "I"}, {ColNo: 9, ColName: "J"},
		{ColNo: 10, ColName: "K"}, {ColNo: 11, ColName: "L"}, {ColNo: 12, ColName: "M"}, {ColNo: 13, ColName: "N"}, {ColNo: 14, ColName: "O"},
		{ColNo: 15, ColName: "P"}, {ColNo: 16, ColName: "Q"}, {ColNo: 17, ColName: "R"}, {ColNo: 18, ColName: "S"}, {ColNo: 19, ColName: "T"},
		{ColNo: 20, ColName: "U"}, {ColNo: 21, ColName: "V"}, {ColNo: 22, ColName: "W"}, {ColNo: 23, ColName: "X"}, {ColNo: 24, ColName: "Y"},
		{ColNo: 25, ColName: "Z"},

		{ColNo: 26, ColName: "AA"},
		{ColNo: 27, ColName: "AB"},
		{ColNo: 26 * 2, ColName: "BA"},
		{ColNo: 26*2 + 5, ColName: "BF"},
		{ColNo: 26 * 3, ColName: "CA"},
		{ColNo: 26 * 26, ColName: "AAA"},
		{ColNo: 26*26 + 5, ColName: "AAF"},
		{ColNo: 26 * 26 * 26, ColName: "AAAA"},
		{ColNo: 26*26*26 + 2, ColName: "AAAC"},
	}

	for _, v := range testData {

		expected := v.ColNo
		actual, _ := ConvertColumnToInt(v.ColName)

		if actual != expected {
			t.Error(fmt.Sprintf("Error. GOT : %d, EXPECTED: %d", actual, expected))
		}
	}
}

func TestCell_Up(t *testing.T) {
	matrix := Matrix{
		{Free, Wall, Wall},
		{Free, Free, Visited},
		{Wall, Free, Free},
		{Wall, Solution, Wall},
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
		{Wall, Solution, Wall},
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
		{Wall, Solution, Wall},
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
		{Wall, Solution, Wall},
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
