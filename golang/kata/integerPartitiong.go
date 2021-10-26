package kata

import "fmt"

// ValidateSolution ...
func ValidateSolution(m [][]int) bool {
	for x := 0; x < len(m); x++ {
		for y := 0; y < len(m[x]); y++ {
			if m[x][y] == 0 {
				return false
			}
			for a := 0; a < len(m); a++ {
				fmt.Println(m[a][y], m[x][y], x, y)
				if m[a][y] == m[x][y] && a != x {
					return false
				}
			}
			for b := 0; b < len(m); b++ {
				if m[x][b] == m[x][y] && b != y {
					return false
				}
			}
		}
	}

	return true
}
