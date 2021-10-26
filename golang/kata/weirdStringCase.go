package kata

import (
	"strings"
)

// ToWeirdCase ...
func ToWeirdCase(str string) string {
	var res string

	// получаем []string слайс из lowercase слов
	strSplit := strings.Split(strings.ToLower(str), " ")

	// проходим в цикле по каждому слову
	for _, v := range strSplit {

		// получаем id символа в слове
		for id := range v {

			// если id символа четный
			if id%2 == 0 {

				// добавляем в строку uppercase символ
				res += strings.ToUpper(string(v[id]))
			} else {

				// нечетный - lowercase
				res += string(v[id])
			}
		}

		// в конце слова добавляем пробел
		res += " "
	}

	// убираем последний пробел
	return res[:len(res)-1]
}
