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
		{"input 0", []int{0}, FunWithAnagramsInputs[0]},
		{"input 1", []int{1}, FunWithAnagramsInputs[1]},
		{"input param negative", []int{-1}, FunWithAnagramsInputs[0]},
		{"input param not available", []int{5}, FunWithAnagramsInputs[0]},
		{"input param nil", []int{}, FunWithAnagramsInputs[0]},
		{"many input params", []int{1, 2, 3, 4}, FunWithAnagramsInputs[1]},
	}

	for _, testCase := range testingTable {
		t.Run(testCase.title, func(t *testing.T) {
			result := GetFunWithAnagramsInput(testCase.input...)
			if !equals(result, testCase.expectedResult) {
				t.Errorf("Test case failed with input: %v, expected: %v; got: %v", testCase.input, testCase.expectedResult, result)
			}
		})
	}
}

func TestFunWithAnagramsInput(t *testing.T) {
	testingTable := []struct {
		title          string
		input          []string
		expectedResult []string
	}{
		{"inpunt A", []string{"framer", "code", "doce", "ecod", "frame", "farmer"}, []string{"code", "frame", "framer"}},
		{"inpunt B", []string{"framer", "code", "doce", "ecod", "frame", "farmer"}, []string{"code", "frame", "framer"}},
		{"all anagrams", []string{"amor", "roma", "mora", "ramo", "omar", "orma"}, []string{"amor"}},
		{"no anagrams", []string{"efg", "fgh", "abc", "def", "bcd", "cde", "hij", "ghi"}, []string{"abc", "bcd", "cde", "def", "efg", "fgh", "ghi", "hij"}},
		{"no anagrams and ordered", []string{"abc", "bcd", "cde", "def", "efg", "fgh", "ghi", "hij"}, []string{"abc", "bcd", "cde", "def", "efg", "fgh", "ghi", "hij"}},
		{"list with only one element", []string{"qwerty"}, []string{"qwerty"}},
		{"empty list", []string{}, []string{}},
	}

	for _, testCase := range testingTable {
		t.Run(testCase.title, func(t *testing.T) {
			result := FunWithAnagrams(testCase.input)
			if !equals(result, testCase.expectedResult) {
				t.Errorf("Test failed with input: %v, expected: %v; got: %v", testCase.input, testCase.expectedResult, result)
			}
		})
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
