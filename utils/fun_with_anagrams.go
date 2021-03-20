package utils

import (
	"sort"
)

func GetFunWithAnagramsInput() []string {
	return []string{"framer", "code", "doce", "ecod", "frame", "farmer"}
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

func FunWithAnagrams(ls []string) []string {
	strList := ls

	for i := 0; i < len(strList); i++ {
		for j := i + 1; j < len(strList); {
			if areAnagrams(strList[i], strList[j]) {
				strList = append(strList[:j], strList[j+1:]...)
			} else {
				j++
			}
		}
	}
	sort.Strings(strList)
	return strList
}
