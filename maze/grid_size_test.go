package maze

import (
	"fmt"
	"testing"
)

func TestSize_IsValid(t *testing.T) {
	testData := []struct {
		size    Size
		isValid bool
	}{
		{Size{Rows: 0, Columns: 0}, false},
		{Size{Rows: 0, Columns: 1}, false},
		{Size{Rows: 1, Columns: 0}, false},
		{Size{Rows: -1, Columns: 0}, false},
		{Size{Rows: 0, Columns: -1}, false},
		{Size{Rows: 5, Columns: -1}, false},
		{Size{Rows: -1, Columns: 10}, false},
		{Size{Rows: 27, Columns: 27}, false},
		{Size{Rows: 30, Columns: 5}, false},
		{Size{Rows: 5, Columns: 40}, false},
		{Size{Rows: 1, Columns: 1}, true},
		{Size{Rows: 5, Columns: 10}, true},
		{Size{Rows: 15, Columns: 13}, true},
		{Size{Rows: 26, Columns: 26}, true},
	}

	for _, v := range testData {
		if v.size.IsValid() != v.isValid {
			t.Error(fmt.Sprintf("Invalid result. Expected: %v, GOT: %v", v.isValid, v.size.IsValid()))
		}
	}
}

func TestParseGridSize(t *testing.T) {
	testData := []struct {
		grid          string
		expectedSize  Size
		expectedError error
	}{
		{"", Size{}, InvalidGridSizeErr},
		{"12 12", Size{}, InvalidGridSizeErr},
		{"x20", Size{}, InvalidGridSizeErr},
		{"6x", Size{}, InvalidGridSizeErr},
		{"0x0", Size{}, InvalidGridSizeErr},
		{"15x0", Size{}, InvalidGridSizeErr},
		{"0x16", Size{}, InvalidGridSizeErr},
		{"27x27", Size{}, InvalidGridSizeErr},
		{"27x5", Size{}, InvalidGridSizeErr},
		{"8x27", Size{}, InvalidGridSizeErr},
		{"1x1", Size{Rows: 1, Columns: 1}, nil},
		{"26x26", Size{Rows: 26, Columns: 26}, nil},
		{"12x16", Size{Rows: 12, Columns: 16}, nil},
		{"21x24", Size{Rows: 21, Columns: 24}, nil},
	}

	for _, v := range testData {
		s, err := ParseGridSize(v.grid)

		if err != v.expectedError {
			t.Error(fmt.Sprintf("Error. Expected: %v, GOT: %v", v.expectedError, err))
		}

		if s.Rows != v.expectedSize.Rows || s.Columns != v.expectedSize.Columns {
			t.Error(fmt.Sprintf("Size should be same. Expected: (%d x %d), GOT: (%d x %d)", v.expectedSize.Rows, v.expectedSize.Columns, s.Rows, s.Columns))
		}
	}
}
