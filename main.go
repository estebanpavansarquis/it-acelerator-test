package main

import (
	"fmt"
)

//region methods
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
//endregion

func main(){
	//region inputs
	s1, s2, s3 := "ironicamente", "renacimiento", "renacimientos"
	//endregion

	//region areAnagrams method calls
	if areAnagrams(s1, s2) {
		fmt.Println(s1,"and", s2, "are anagrams.")
	} else {
		fmt.Println(s1,"and", s2, "are not anagrams.")
	}

	if areAnagrams(s1, s3) {
		fmt.Println(s1,"and", s3, "are anagrams.")
	} else {
		fmt.Println(s1,"and", s3, "are not anagrams.")
	}
	//endregion

}