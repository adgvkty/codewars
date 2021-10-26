package kata

import (
	"strings"
)

// ToNato ...
func ToNato(words string) string {
	var result string

	nato := []string{"Alfa", "Bravo", "Charlie", "Delta", "Echo", "Foxtrot", "Golf", "Hotel", "India", "Juliett", "Kilo", "Lima", "Mike", "November", "Oscar", "Papa", "Quebec", "Romeo", "Sierra", "Tango", "Uniform", "Victor", "Whiskey", "X-ray", "Yankee", "Zulu"}
	badSymbolds := []string{",", ".", "!", "?"}

	for i := 0; i < len(words); i++ {
		for _, badSymbol := range badSymbolds {
			if string(words[i]) == badSymbol {
				result += badSymbol + " "
				break
			}
		}

		for _, natoWord := range nato {
			if strings.ToUpper(string(words[i])) == string(natoWord[0]) {
				result += natoWord + " "
				break
			}
		}
	}
	return result[:len(result)-1]
}
