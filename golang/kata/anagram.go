package kata

// Anagrams ...
func Anagrams(word string, words []string) []string {
	var result []string = []string{}
	var wordValue byte
	var tempWordValue byte

	for id := range word {
		wordValue += word[id]
	}

	for _, word := range words {
		tempWordValue = 0
		for id := range word {
			tempWordValue += word[id]
		}
		if wordValue == tempWordValue {
			result = append(result, word)
		}
	}

	return result
}
