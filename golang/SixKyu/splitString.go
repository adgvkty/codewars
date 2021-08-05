package six

import "fmt"

// Split is splitting string into []string
func Split(str string) []string {
	var result []string

	for i := 0; i < len(str); i += 2 {
		if i == len(str)-1 {
			result = append(result, fmt.Sprintf("%v_", str[i:i+1]))
			continue
		}
		result = append(result, str[i:i+2])
	}
	return result
}
