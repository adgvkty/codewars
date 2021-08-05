package six

import "fmt"

// Tickets ...
func Tickets(peopleInLine []int) string {
	var bills []int

	for i := 0; i < len(peopleInLine); i++ {
		fmt.Println(bills, peopleInLine[i])
		if peopleInLine[i] == 25 {
			bills = append(bills, peopleInLine[i])
			continue
		} else {
			if peopleInLine[i] == 50 {
				if findInArray(bills, 25) {
					bills = deleteFromArray(bills, 25)
					bills = append(bills, 50)
					continue
				} else {
					return "NO"
				}
			} else if peopleInLine[i] == 100 {
				if findInArray(bills, 50) && findInArray(bills, 25) {
					bills = deleteFromArray(bills, 50)
					bills = deleteFromArray(bills, 25)
					continue
				} else if countInArray(bills, 25) >= 3 {
					for x := 0; x < 3; x++ {
						bills = deleteFromArray(bills, 25)
					}
					continue
				} else {
					return "NO"
				}
			}
		}
	}

	return "YES"
}

func countInArray(array []int, digit int) int {
	var count int
	for _, number := range array {
		if number == digit {
			count++
		}
	}
	return count
}

func findInArray(array []int, digit int) bool {
	for _, number := range array {
		if number == digit {
			return true
		}
	}
	return false
}

func deleteFromArray(array []int, digit int) []int {
	for id, number := range array {
		if number == digit {
			array[id] = array[len(array)-1]
			return array[:len(array)-1]
		}
	}
	return []int{}
}
