package utils

import (
	"testing"
)

var testingTable = []struct {
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
	{"big list", []string{"framer", "code", "doce", "ecod", "frame", "farmer", "amor", "roma", "mora", "ramo", "omar", "orma", "efg", "fgh", "abc", "def", "bcd", "cde", "hij", "ghi", "code", "doce", "ecod", "frame", "farmer", "efg", "fgh", "abc", "def", "bcd", "cde", "hij", "ghi", "code", "doce", "ecod", "frame", "farmer", "efg", "fgh", "abc", "def", "bcd", "cde", "hij", "ghi", "code", "doce", "ecod", "frame", "farmer", "efg", "fgh", "abc", "def", "bcd", "cde", "hij", "ghi", "code", "doce", "ecod", "frame", "farmer", "efg", "fgh", "abc", "def", "bcd", "cde", "hij", "ghi", "code", "doce", "ecod", "frame", "farmer", "efg", "fgh", "abc", "def", "bcd", "cde", "hij", "ghi", "code", "doce", "ecod", "frame", "farmer", "efg", "fgh", "abc", "def", "bcd", "cde", "hij", "ghi", "code", "doce", "ecod", "frame", "farmer", "efg", "fgh", "abc", "def", "bcd", "cde", "hij", "ghi", "ghi", "code", "doce", "ecod", "frame", "farmer", "efg", "fgh", "abc", "def", "bcd", "cde", "hij", "ghi", "code", "doce", "ecod", "frame", "farmer", "efg", "fgh", "abc", "def", "bcd", "cde", "hij", "ghi", "code", "doce", "ecod", "frame", "farmer", "efg", "fgh", "abc", "def", "bcd", "cde", "hij", "ghi", "ghi", "code", "doce", "ecod", "frame", "farmer", "efg", "fgh", "abc", "def", "bcd", "cde", "hij", "ghi", "code", "doce", "ecod", "frame", "farmer", "efg", "fgh", "abc", "def", "bcd", "cde", "hij", "ghi", "code", "doce", "ecod", "frame", "farmer", "efg", "fgh", "abc", "def", "bcd", "cde", "hij", "ghi", "ghi", "code", "doce", "ecod", "frame", "farmer", "efg", "fgh", "abc", "def", "bcd", "cde", "hij", "ghi", "code", "doce", "ecod", "frame", "farmer", "efg", "fgh", "abc", "def", "bcd", "cde", "hij", "ghi", "code", "doce", "ecod", "frame", "farmer", "efg", "fgh", "abc", "def", "bcd", "cde", "hij", "ghi", "ghi", "code", "doce", "ecod", "frame", "farmer", "efg", "fgh", "abc", "def", "bcd", "cde", "hij", "ghi", "code", "doce", "ecod", "frame", "farmer", "efg", "fgh", "abc", "def", "bcd", "cde", "hij", "ghi", "code", "doce", "ecod", "frame", "farmer", "efg", "fgh", "abc", "def", "bcd", "cde", "hij", "ghi", "ghi", "code", "doce", "ecod", "frame", "farmer", "efg", "fgh", "abc", "def", "bcd", "cde", "hij", "ghi", "code", "doce", "ecod", "frame", "farmer", "efg", "fgh", "abc", "def", "bcd", "cde", "hij", "ghi", "code", "doce", "ecod", "frame", "farmer", "efg", "fgh", "abc", "def", "bcd", "cde", "hij", "ghi", "ghi", "code", "doce", "ecod", "frame", "farmer", "efg", "fgh", "abc", "def", "bcd", "cde", "hij", "ghi", "code", "doce", "ecod", "frame", "farmer", "efg", "fgh", "abc", "def", "bcd", "cde", "hij", "ghi", "code", "doce", "ecod", "frame", "farmer", "efg", "fgh", "abc", "def", "bcd", "cde", "hij", "ghi", "ghi", "code", "doce", "ecod", "frame", "farmer", "efg", "fgh", "abc", "def", "bcd", "cde", "hij", "ghi", "code", "doce", "ecod", "frame", "farmer", "efg", "fgh", "abc", "def", "bcd", "cde", "hij", "ghi", "code", "doce", "ecod", "frame", "farmer", "efg", "fgh", "abc", "def", "bcd", "cde", "hij", "ghi", "ghi", "code", "doce", "ecod", "frame", "farmer", "efg", "fgh", "abc", "def", "bcd", "cde", "hij", "ghi", "code", "doce", "ecod", "frame", "farmer", "efg", "fgh", "abc", "def", "bcd", "cde", "hij", "ghi", "code", "doce", "ecod", "frame", "farmer", "efg", "fgh", "abc", "def", "bcd", "cde", "hij", "ghi", "ghi", "code", "doce", "ecod", "frame", "farmer", "efg", "fgh", "abc", "def", "bcd", "cde", "hij", "ghi", "code", "doce", "ecod", "frame", "farmer", "efg", "fgh", "abc", "def", "bcd", "cde", "hij", "ghi", "code", "doce", "ecod", "frame", "farmer", "efg", "fgh", "abc", "def", "bcd", "cde", "hij", "ghi", "qwerty"}, []string{"abc", "amor", "bcd", "cde", "code", "def", "efg", "fgh", "frame", "framer", "ghi", "hij", "qwerty"}},
}

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
	for _, testCase := range testingTable {
		t.Run(testCase.title, func(t *testing.T) {
			result := FunWithAnagrams(testCase.input)
			if !equals(result, testCase.expectedResult) {
				t.Errorf("Test failed with input: %v, expected: %v; got: %v", testCase.input, testCase.expectedResult, result)
			}
		})
	}
}

func BenchmarkAreAnagrams(b *testing.B) {
	benchmarkingTable := []struct {
		title    string
		function func(string, string) bool
	}{
		{"AreAnagrams", AreAnagrams},
	}

	for _, benchmarkingCase := range benchmarkingTable {
		for _, testingCase := range testingTable {
			b.Run(benchmarkingCase.title+"_with_input_"+testingCase.title, func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					FunWithAnagrams(testingCase.input)
				}
			})
		}
	}
}

func BenchmarkFunWithAnagrams(b *testing.B) {
	benchmarkingTable := []struct {
		title    string
		function func(slice []string) []string
	}{
		{"FunWithAnagrams", FunWithAnagrams},
	}

	for _, benchmarkingCase := range benchmarkingTable {
		for _, testingCase := range testingTable {
			b.Run(benchmarkingCase.title+"_with_input_"+testingCase.title, func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					benchmarkingCase.function(testingCase.input)
				}
			})
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
