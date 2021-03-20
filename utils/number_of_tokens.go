package utils

func NumberOfTokens(experyLimit int, commands [][]int) int {
	tokens := make(map[int]int)
	activeTokens, time := 0, commands[len(commands)-1][2]

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
			activeTokens++
		}
	}
	return activeTokens
}
