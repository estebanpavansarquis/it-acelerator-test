package utils

import (
	"fmt"
	"unicode"
)

const MostRepeatedLetterInput = "Hola, mi nombre es Ariel, soy una sirena, vivo bajo el mar"

//comentario malardo

func MostRepeatedLetter(s string) string {
	letterCounter := make(map[rune]int)
	max := struct {
		count  int
		letter rune
	}{}

	for _, c := range s {
		if unicode.IsLetter(c) {
			c = unicode.ToUpper(c)
			letterCounter[c]++
			if letterCounter[c] > max.count {
				max.count = letterCounter[c]
				max.letter = c
			}
		}
	}

	return fmt.Sprintf("La letra que más se repite es la %c con %d repeticiones", max.letter, max.count)
}

func notUsed(){
	print("I am never used!")
}