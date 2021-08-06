package six

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

// Stati ...
func Stati(str string) string {
	var allTeamTimes []int
	fmt.Println(str)

	// Проверка на пустую строку
	if str == "" {
		return ""
	}

	// Получаем отдельно время каждой команды
	teams := strings.Split(str, ",")

	// Идём циклом по каждому времени
	for _, team := range teams {

		// Обнуляем служебную переменную
		time := 0

		// Разделяем время на часы, минуты и секунды
		teamTime := strings.Split(team, "|")

		// Идём циклом по каждому времени
		for id, number := range teamTime {

			// Убираем пробелы, если есть
			number = strings.Replace(number, " ", "", -1)

			// Часы всегда имеют 0-й индекс, минуты всегда 1-й и секунды 2-й
			// Поэтому в свитче определяем какое у нас время
			// И переводим в секунды
			switch id {
			case 0:
				hours, _ := strconv.Atoi(number)
				time += hours * 60 * 60
			case 1:
				minutes, _ := strconv.Atoi(number)
				time += minutes * 60
			case 2:
				seconds, _ := strconv.Atoi(number)
				time += seconds
			}
		}

		// Добавляем время в секундах в массив всех времен
		allTeamTimes = append(allTeamTimes, time)
	}

	// Находим максимальное и минимальное значение в массиве времени
	maxTime, minTime := findMax(allTeamTimes), findMin(allTeamTimes)

	// Считаем range
	rangeValue := maxTime - minTime

	// Находим среднее значение функцией
	averageValue := findAverage(allTeamTimes)

	// Находим медиану функцией
	medianValue := findMedian(allTeamTimes)

	// Собираем вместе результат, как форматированную строку
	// Секунды переводим в нужный нам формат функцией
	resultString := fmt.Sprintf("Range: %v Average: %v Median: %v",
		secToTime(rangeValue), secToTime(averageValue), secToTime(medianValue))

	return resultString
}

// secToTime переводит секунды в форматированную строку времени
func secToTime(time int) string {
	hours := time / 60 / 60
	time -= hours * 60 * 60

	minutes := time / 60
	time -= minutes * 60

	seconds := time

	return fmt.Sprintf("%v|%v|%v",
		beatifyNumber(strconv.Itoa(hours)),
		beatifyNumber(strconv.Itoa(minutes)),
		beatifyNumber(strconv.Itoa(seconds)))
}

// beatifyNumber добавляет 0 к числу, если оно < 10
func beatifyNumber(number string) string {
	if len(number) == 1 {
		number = "0" + number
	}

	return number
}

// findMedian ищет медиану в []int
func findMedian(array []int) int {

	// Сортируем слайс
	sort.Slice(array, func(i, j int) bool {
		return array[i] < array[j]
	})

	// Находим ID медианного значения
	medianID := len(array) / 2

	// Если в массиве нечетное количество элементов
	// Значит есть "неоспоримая" медиана
	// Например массив []int{2, 5, 8, 14, 32}
	// В нём медианна 8, поскольку находится посередине
	// medianID в таком случае равен 2.5, но в интовой форме 2
	// Поскольку массив начинается с 0-го индекса
	// То у 8-ми как раз индекс 2
	if len(array)%2 != 0 {
		return array[medianID]
	}

	// Если же у нас четное количество элементов
	// Мы ищем среднее арифметическое
	// Между двумя "центральными" элементами
	return (array[medianID] + array[medianID-1]) / 2
}

// findAverage находим среднеарифметическое массива
func findAverage(array []int) int {
	var allTime int

	// циклом добавляем в переменную всё время
	for _, number := range array {
		allTime += number
	}

	// делим всё время на кол-во элементов = длину массива
	return allTime / len(array)
}

// findMin находит минимальное значение в массива
func findMin(array []int) int {
	// присваиваем минимум первому числу массива
	var min = array[0]

	// если в цикле находим меньше - переприсваиваем
	for _, number := range array {
		if number < min {
			min = number
		}
	}

	return min
}

// findMax находит максимальное значение массива
func findMax(array []int) int {
	// присваиваем максимум первому числу массива
	var max = array[0]

	// если в цикле находим больше - переприсваиваем
	for _, number := range array {
		if number > max {
			max = number
		}
	}

	return max
}
