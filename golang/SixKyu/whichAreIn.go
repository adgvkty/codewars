package six

import (
	"sort"
	"strings"
)

// InArray ...
func InArray(array1 []string, array2 []string) []string {
	var result []string
	for _, word := range array1 {
		for _, secondWord := range array2 {
			if strings.Contains(secondWord, word) {
				if !findInStringArray(result, word) {
					result = append(result, word)
				}
			}
		}
	}
	sort.Strings(result)
	return result
}

func findInStringArray(array []string, element string) bool {
	for _, word := range array {
		if word == element {
			return true
		}
	}
	return false
}
