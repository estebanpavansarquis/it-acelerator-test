package utils

import (
	"testing"
)

func TestFunWithAnagramsInput(t *testing.T) {
	testingTable := []struct {
		title          string
		input          []string
		expectedResult []string
	}{
		{"Inpunt A", []string{"framer", "code", "doce", "ecod", "frame", "farmer"}, []string{"code", "frame", "framer"}},
		{"Inpunt B", []string{"framer", "code", "doce", "ecod", "frame", "farmer"}, []string{"code", "frame", "framer"}},
		{"All anagrams", []string{"amor", "roma", "mora", "ramo", "omar", "orma"}, []string{"amor"}},
		{"No anagrams", []string{"efg", "fgh", "abc", "def", "bcd", "cde", "hij", "ghi"}, []string{"abc", "bcd", "cde", "def", "efg", "fgh", "ghi", "hij"}},
		{"No anagrams and ordered", []string{"abc", "bcd", "cde", "def", "efg", "fgh", "ghi", "hij"}, []string{"abc", "bcd", "cde", "def", "efg", "fgh", "ghi", "hij"}},
		{"List with only one element", []string{"qwerty"}, []string{"qwerty"}},
		{"Empty list", []string{}, []string{}},
	}

	for _, testCase := range testingTable {
		result := FunWithAnagrams(testCase.input)
		if !equals(result, testCase.expectedResult) {
			t.Errorf("Test case \"%v\" failed with input: %v, expected: %v; got: %v", testCase.title, testCase.input, testCase.expectedResult, result)
		}
	}
}

func equals(slice1 []string, slice2 []string) bool {
	if len(slice1) != len(slice2) {
		return false
	}
	for i, s := range slice1 {
		if s != slice2[i] {
			return false
		}
	}
	return true
}
