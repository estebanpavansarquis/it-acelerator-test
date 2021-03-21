package utils

import (
	"testing"
)

func TestFunWithAnagramsInput1(t *testing.T) {
	testInput := []string{"framer", "code", "doce", "ecod", "frame", "farmer"}
	expectedResult := []string{"code", "frame", "framer"}

	result := FunWithAnagrams(testInput)

	if !equals(result, expectedResult) {
		t.Fatalf("Test failed with input: %v, expected: %v; got: %v", testInput, expectedResult, result)
	}
}

func TestFunWithAnagramsInput2(t *testing.T) {
	testInput := []string{"roma", "ramo", "amor", "mora", "oran", "code", "doce", "cero"}
	expectedResult := []string{"cero", "code", "oran", "roma"}
	result := FunWithAnagrams(testInput)

	if !equals(result, expectedResult) {
		t.Fatalf("Test failed with input: %v, expected: %v; got: %v", testInput, expectedResult, result)
	}
}

func TestFunWithAnagramsInput3(t *testing.T) {
	testInput := []string{"roma", "roma", "roma", "roma", "roma", "roma", "roma", "roma"}
	expectedResult := []string{"roma"}
	result := FunWithAnagrams(testInput)

	if !equals(result, expectedResult) {
		t.Fatalf("Test failed with input: %v, expected: %v; got: %v", testInput, expectedResult, result)
	}
}

func TestFunWithAnagramsInput4(t *testing.T) {
	testInput := []string{"abc", "bcd", "cde", "def", "efg", "fgh", "ghi", "hij"}
	expectedResult := []string{"abc", "bcd", "cde", "def", "efg", "fgh", "ghi", "hij"}
	result := FunWithAnagrams(testInput)

	if !equals(result, expectedResult) {
		t.Fatalf("Test failed with input: %v, expected: %v; got: %v", testInput, expectedResult, result)
	}
}

func TestFunWithAnagramsInput5(t *testing.T) {
	testInput := []string{}
	expectedResult := []string{}
	result := FunWithAnagrams(testInput)

	if !equals(result, expectedResult) {
		t.Fatalf("Test failed with input: %v, expected: %v; got: %v", testInput, expectedResult, result)
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
