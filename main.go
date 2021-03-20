package main

import (
	"fmt"

	"it-acelerator-test/utils"
)

func main() {
	// region inputs

	// funWithAnagrams Input
	strList := []string{"framer", "code", "doce", "ecod", "frame", "farmer"}

	// authenticationTokens Input
	experationLimit := 4
	inputArray := [][]int{{0, 1, 1}, {0, 1, 3}, {0, 2, 2}, {1, 1, 5}, {1, 2, 7}, {1, 3, 7}}

	// endregion

	// funWithAnagrams function call
	fmt.Println("My original input:", strList)
	fmt.Println("After funWithAnagrams:", utils.FunWithAnagrams(strList))

	// authenticationTokens function call
	numberOfTokens := utils.NumberOfTokens(experationLimit, inputArray)
	fmt.Println("There are", numberOfTokens, "active Tokens.")
}
