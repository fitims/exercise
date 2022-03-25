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
		//{ColNo: 0, ColName: "A"}, {ColNo: 1, ColName: "B"}, {ColNo: 2, ColName: "C"}, {ColNo: 3, ColName: "D"}, {ColNo: 4, ColName: "E"},
		//{ColNo: 5, ColName: "F"}, {ColNo: 6, ColName: "G"}, {ColNo: 7, ColName: "H"}, {ColNo: 8, ColName: "I"}, {ColNo: 9, ColName: "J"},
		//{ColNo: 10, ColName: "K"}, {ColNo: 11, ColName: "L"}, {ColNo: 12, ColName: "M"}, {ColNo: 13, ColName: "N"}, {ColNo: 14, ColName: "O"},
		//{ColNo: 15, ColName: "P"}, {ColNo: 16, ColName: "Q"}, {ColNo: 17, ColName: "R"}, {ColNo: 18, ColName: "S"}, {ColNo: 19, ColName: "T"},
		//{ColNo: 20, ColName: "U"}, {ColNo: 21, ColName: "V"}, {ColNo: 22, ColName: "W"}, {ColNo: 23, ColName: "X"}, {ColNo: 24, ColName: "Y"},
		//{ColNo: 25, ColName: "Z"},
		//
		//{ColNo: 26, ColName: "AA"},
		//{ColNo: 27, ColName: "AB"},
		//{ColNo: 26 * 2, ColName: "BA"},
		//{ColNo: 26*2 + 5, ColName: "BF"},
		//{ColNo: 26 * 3, ColName: "CA"},
		//{ColNo: 26 * 26, ColName: "AAA"},
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
