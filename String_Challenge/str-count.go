package main

func is_consonants(char rune) bool {
	if (char == 'a') || (char == 'i') || (char == 'e') || (char == 'o') || (char == 'u') || (char == 'A') || (char == 'I') || (char == 'E') || (char == 'U') || (char == 'O') || (char == ' ') {

		return false
	} else {
		return true
	}
}

func StringChallenge(str string) int {
	count := 0
	for _, char := range str {
		if is_consonants(char) {
			count = count + 1
		}
	}
	return count
}
