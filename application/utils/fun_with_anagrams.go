package utils

import (
	"sort"
)

var (
	FunWithAnagramsInputs = [][]string{

		// expected result: code frame framer
		{"framer", "code", "doce", "ecod", "frame", "farmer"},

		// expected result: cero, code, oran, roma
		{"roma", "ramo", "amor", "mora", "oran", "code", "doce", "cero"},

		// expected result: abc, amor, bcd, cde, code, def, efg, fgh, frame, framer, ghi, hij, qwerty
		{"framer", "code", "doce", "ecod", "frame", "farmer", "amor", "roma", "mora", "ramo", "omar", "orma", "efg", "fgh", "abc", "def", "bcd", "cde", "hij", "ghi", "code", "doce", "ecod", "frame", "farmer", "efg", "fgh", "abc", "def", "bcd", "cde", "hij", "ghi", "code", "doce", "ecod", "frame", "farmer", "efg", "fgh", "abc", "def", "bcd", "cde", "hij", "ghi", "code", "doce", "ecod", "frame", "farmer", "efg", "fgh", "abc", "def", "bcd", "cde", "hij", "ghi", "code", "doce", "ecod", "doce", "ecod", "frame", "farmer", "amor", "roma", "mora", "ramo", "omar", "orma", "efg", "fgh", "abc", "def", "bcd", "cde", "hij", "ghi", "code", "doce", "ecod", "frame", "farmer", "efg", "fgh", "abc", "def", "bcd", "cde", "hij", "ghi", "code", "doce", "ecod", "frame", "farmer", "efg", "fgh", "abc", "def", "bcd", "cde", "hij", "ghi", "code", "doce", "ecod", "frame", "farmer", "efg", "fgh", "abc", "def", "bcd", "cde", "hij", "ghi", "code", "doce", "ecod", "doce", "ecod", "frame", "farmer", "amor", "roma", "mora", "ramo", "omar", "orma", "efg", "fgh", "abc", "def", "bcd", "cde", "hij", "ghi", "code", "doce", "ecod", "frame", "farmer", "efg", "fgh", "abc", "def", "bcd", "cde", "hij", "ghi", "code", "doce", "ecod", "frame", "farmer", "efg", "fgh", "abc", "def", "bcd", "cde", "hij", "ghi", "code", "doce", "ecod", "frame", "farmer", "efg", "fgh", "abc", "def", "bcd", "cde", "hij", "ghi", "code", "doce", "ecod", "frame", "farmer", "efg", "fgh", "abc", "def", "bcd", "cde", "hij", "ghi", "code", "doce", "ecod", "frame", "farmer", "efg", "fgh", "abc", "def", "bcd", "cde", "hij", "ghi", "code", "doce", "ecod", "frame", "farmer", "efg", "fgh", "abc", "def", "bcd", "cde", "hij", "ghi", "code", "doce", "ecod", "frame", "farmer", "efg", "fgh", "abc", "def", "bcd", "cde", "hij", "ghi", "ghi", "code", "doce", "ecod", "frame", "farmer", "efg", "fgh", "abc", "def", "bcd", "cde", "hij", "ghi", "code", "doce", "ecod", "frame", "farmer", "efg", "fgh", "abc", "def", "bcd", "cde", "hij", "ghi", "code", "doce", "ecod", "frame", "farmer", "efg", "fgh", "abc", "def", "bcd", "cde", "hij", "ghi", "ghi", "code", "doce", "ecod", "frame", "farmer", "efg", "fgh", "abc", "def", "bcd", "cde", "hij", "ghi", "code", "doce", "ecod", "frame", "farmer", "efg", "fgh", "abc", "def", "bcd", "cde", "hij", "ghi", "code", "doce", "ecod", "frame", "farmer", "efg", "fgh", "abc", "def", "bcd", "cde", "hij", "ghi", "ghi", "code", "doce", "ecod", "frame", "farmer", "efg", "fgh", "abc", "def", "bcd", "cde", "hij", "ghi", "code", "doce", "ecod", "frame", "farmer", "efg", "fgh", "abc", "def", "bcd", "cde", "hij", "ghi", "code", "doce", "ecod", "frame", "farmer", "efg", "fgh", "abc", "def", "bcd", "cde", "hij", "ghi", "ghi", "code", "doce", "ecod", "frame", "farmer", "efg", "fgh", "abc", "def", "bcd", "cde", "hij", "ghi", "code", "doce", "ecod", "frame", "farmer", "efg", "fgh", "abc", "def", "bcd", "cde", "hij", "ghi", "code", "doce", "ecod", "frame", "farmer", "efg", "fgh", "abc", "def", "bcd", "cde", "hij", "ghi", "ghi", "code", "doce", "ecod", "frame", "farmer", "efg", "fgh", "abc", "def", "bcd", "cde", "hij", "ghi", "code", "doce", "ecod", "frame", "farmer", "efg", "fgh", "abc", "def", "bcd", "cde", "hij", "ghi", "code", "doce", "ecod", "frame", "farmer", "efg", "fgh", "abc", "def", "bcd", "cde", "hij", "ghi", "ghi", "code", "doce", "ecod", "frame", "farmer", "efg", "fgh", "abc", "def", "bcd", "cde", "hij", "ghi", "code", "doce", "ecod", "frame", "farmer", "efg", "fgh", "abc", "def", "bcd", "cde", "hij", "ghi", "code", "doce", "ecod", "frame", "farmer", "efg", "fgh", "abc", "def", "bcd", "cde", "hij", "ghi", "ghi", "code", "doce", "ecod", "frame", "farmer", "efg", "fgh", "abc", "def", "bcd", "cde", "hij", "ghi", "code", "doce", "ecod", "frame", "farmer", "efg", "fgh", "abc", "def", "bcd", "cde", "hij", "ghi", "code", "doce", "ecod", "frame", "farmer", "efg", "fgh", "abc", "def", "bcd", "cde", "hij", "ghi", "ghi", "code", "doce", "ecod", "frame", "farmer", "efg", "fgh", "abc", "def", "bcd", "cde", "hij", "ghi", "code", "doce", "ecod", "frame", "farmer", "efg", "fgh", "abc", "def", "bcd", "cde", "hij", "ghi", "code", "doce", "ecod", "frame", "farmer", "efg", "fgh", "abc", "def", "bcd", "cde", "hij", "ghi", "ghi", "code", "doce", "ecod", "frame", "farmer", "efg", "fgh", "abc", "def", "bcd", "cde", "hij", "ghi", "code", "doce", "ecod", "frame", "farmer", "efg", "fgh", "abc", "def", "bcd", "cde", "hij", "ghi", "code", "doce", "ecod", "frame", "farmer", "efg", "fgh", "abc", "def", "bcd", "cde", "hij", "ghi", "qwerty"},
	}
)

func GetFunWithAnagramsInput(inputNumber ...int) []string {
	inputWanted := 0
	if len(inputNumber) > 0 && inputNumber[0] >= 0 && inputNumber[0] <= len(FunWithAnagramsInputs) {
		inputWanted = inputNumber[0]
	}
	return FunWithAnagramsInputs[inputWanted]
}

//
func FunWithAnagrams (slice []string) []string    {
	// to avoid changes on the original slice
		stringList := make([]string, len(slice))
	copy(stringList, slice)

	for i := 0; i < len(stringList); i++ {
		for j := i + 1; j < len(stringList); {
			if areAnagrams(stringList[i], stringList[j]) {
				stringList = removeIndexFromSlice(stringList, j)
			} else {
				j++
			}
		}
	}
		sort.Strings(stringList)
	return stringList
}

func areAnagrams(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	count := make(map[rune]int)

	for _, c := range s1 {
		count[c] = count[c] + 1
	}
	for _, c := range s2 {
		count[c] = count[c] - 1
		if count[c] < 0 {
			return false
		}
	}
	return true
}

func removeIndexFromSlice(slice []string, index int) []string {
	return append(slice[:index], slice[index+1:]...)
}
