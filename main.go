package main

import (
	"fmt"
)

//region methods
func areAnagrams(s1 string, s2 string) bool {
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
		fmt.Println(s1,"and", s2, "are anagrams.")
	} else {
		fmt.Println(s1,"and", s2, "are not anagrams.")
	}
	//endregion

}