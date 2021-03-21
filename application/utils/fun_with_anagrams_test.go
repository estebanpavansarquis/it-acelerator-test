package utils

import (
	"testing"
)

func TestGetFunWithAnagramsInput(t *testing.T) {
	testingTable := []struct {
		title          string
		input          []int
		expectedResult []string
	}{
		{"Input 0", []int{0}, FunWithAnagramsInputs[0]},
		{"Input 1", []int{1}, FunWithAnagramsInputs[1]},
		{"Input param negative", []int{-1}, FunWithAnagramsInputs[0]},
		{"Input param not available", []int{5}, FunWithAnagramsInputs[0]},
		{"Input param nil", []int{}, FunWithAnagramsInputs[0]},
		{"Many input params", []int{1, 2, 3, 4}, FunWithAnagramsInputs[1]},
	}

	for _, testCase := range testingTable {
		result := GetFunWithAnagramsInput(testCase.input...)
		if !equals(result, testCase.expectedResult) {
			t.Errorf("Test case \"%v\" failed with input: %v, expected: %v; got: %v", testCase.title, testCase.input, testCase.expectedResult, result)
		}
	}
}

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
