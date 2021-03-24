package utils

import (
	"testing"
)

func TestMostRepeatedLetter(t *testing.T) {
	testingTable := []struct {
		title          string
		input          string
		expectedResult string
	}{
		{"exercise input 1", MostRepeatedLetterInput, "La letra que más se repite es la A con 6 repeticiones"},
		{"special characters", "qqweer12  5$^^ ¨¨ ¨¨^ ^ }[(((((((({{ …–„;:;_:_^^^     31 24ty", "La letra que más se repite es la Q con 2 repeticiones"},
		{"casing", "  HolA, mi nombr12e es Ariel, SOy una sirena,,,,,, vivo bajo el mar   ", "La letra que más se repite es la A con 6 repeticiones"},
	}

	for _, testCase := range testingTable {
		t.Run(testCase.title, func(t *testing.T) {
			result := MostRepeatedLetter(testCase.input)
			if result != testCase.expectedResult {
				t.Errorf("Test case failed with input: %v, expected: %v; got: %v", testCase.input, testCase.expectedResult, result)
			}
		})
	}
}
