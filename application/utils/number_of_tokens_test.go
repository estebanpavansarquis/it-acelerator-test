package utils

import (
	"testing"
)

func TestNumberOfTokens(t *testing.T) {
	input := GetNumberOfTokensInput()
	expectedResult := 1

	result := NumberOfTokens(input.ExperationLimit, input.InputArray)

	if result != expectedResult {
		t.Errorf("Test failed with input: %v, expected: %v; got: %v", input, expectedResult, result)
	}
}
