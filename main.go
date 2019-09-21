package main

func numberOfTokens(experyLimit int, commands [][]int) (int, int) {
	at, t := 0, 0
	return at, t
}

func main(){
	experyLimit := 4
	inputArray := [][]int{{0, 1, 1}, {0, 1, 3}, {0, 2, 2}, {1, 1, 5}, {1, 2, 7}, {1, 3, 7}}

	at, t := numberOfTokens(experyLimit, inputArray)
	fmt.Println("There are", at, "active Tokens at time", t)
}