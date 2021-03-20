package main

import (
	"fmt"

	"it-acelerator-test/utils"
)

func main() {

	// funWithAnagrams function call
	funWithAnagramsInput := utils.GetFunWithAnagramsInput()
	fmt.Println("\nMy funWithAnagrams original input:", funWithAnagramsInput)
	fmt.Println("Result of running funWithAnagrams:", utils.FunWithAnagrams(funWithAnagramsInput))

	// authenticationTokens function call
	numberOfTokensInput := utils.GetNumberOfTokensInput()
	fmt.Println("\nMy numberOfTokens original input:", numberOfTokensInput)
	fmt.Println("Result of running numberOfTokens: ", utils.NumberOfTokens(numberOfTokensInput.ExperationLimit, numberOfTokensInput.InputArray), " active Tokens.")
}
