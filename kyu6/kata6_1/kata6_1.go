package kata6_1

import "strings"

func DuplicateEncode(word string) (newWord string) {

	lowerWord := (strings.ToLower(word))

	minimap := make(map[rune]int)
	for _, char := range lowerWord {
		minimap[char]++
	}

	for _, char := range lowerWord {
		if minimap[char] > 1 {
			newWord = newWord + ")"
		} else {
			newWord = newWord + "("
		}
	}

	return newWord
}
