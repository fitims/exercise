package maze

import "testing"

func TestPath_GetExitCell(t *testing.T) {
	p := Path{
		{Row: 0, Col: 0}, {Row: 1, Col: 0}, {Row: 1, Col: 1}, {Row: 2, Col: 1},
	}

	actual := p.GetExitCell()

	if !actual.IsSame(Cell{Row: 2, Col: 1}) {
		t.Error("exit cell is not valid")
	}
}
