package main

import (
	"fmt"
	"sort"
)

//region methods

func numberOfTokens(experyLimit int, commands [][]int) int {
	tokens := make(map[int]int)
	activeTokens, time := 0,  commands[len(commands)-1][2]

	for _, c := range commands {
		if c[0] == 0 { // If it's a new Token command
			if tokens[c[1]] == 0 { // If it's not initialized
				tokens[c[1]] = c[2] + experyLimit
			}
		} else { // If it's an update command
			if tokens[c[1]] != 0 {
				if c[2] <= tokens[c[1]] {
					tokens[c[1]] = tokens[c[1]] + experyLimit
				}
			}
		}
	}

	for _, token := range tokens {
		if token >= time {
			activeTokens ++
		}
	}
	return activeTokens
}


func areAnagrams(s1 string, s2 string) bool {
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
	for _, c := range count {
		if c != 0 {
			return false
		}
	}
	return true
}

func funWithAnagrams(ls []string) []string {
	strList := ls

	for i := 0; i < len(strList); i++{
		for j := i + 1; j < len(strList); {
			if areAnagrams(strList[i], strList[j]) {
				strList = append(strList[:j], strList[j+1:]...)
			}else{
				j++
			}
		}
	}
	sort.Strings(strList)
	return strList
}

//endregion

func main() {
	//region inputs
    
    //funWithAnagrams Input
	strList := []string{"framer","code", "doce", "ecod", "frame", "farmer"}	

    //authenticationTokens Input
    experyLimit := 4
	inputArray := [][]int{{0, 1, 1}, {0, 1, 3}, {0, 2, 2}, {1, 1, 5}, {1, 2, 7}, {1, 3, 7}}
    
    //endregion

    //funWithAnagrams function call
	fmt.Println("My original input:", strList)
	fmt.Println("After funWithAnagrams:", funWithAnagrams(strList))

    //authenticationTokens function call
	at:= numberOfTokens(experyLimit, inputArray)
	fmt.Println("There are", at, "active Tokens.")
}