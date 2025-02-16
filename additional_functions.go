package main

import (
	"math"
	"sort"
)

// это функция для поиска одинаковой самой длинной подстроки в массиве строк
func finding_substrings_array(strs []string) string {
	if len(strs) == 1 {
		return strs[0]
	}

	var count int = 0
	var max_count int = 0
	var count_ii int = 1
	var tmp string = strs[0]
	var result string = ""
	for ii := 1; ii < len(strs); ii++ {
		var tmp_len int = 0
		if len(strs[ii]) <= len(tmp) {
			tmp_len = len(strs[ii])
		} else {
			tmp_len = len(tmp)
		}

		for jj := 0; jj < tmp_len; jj++ {
			if tmp[jj] == strs[ii][jj] {
				count++
			} else if tmp[jj] != strs[ii][jj] {
				if count_ii == ii {
					if max_count < count {
						max_count = count
						result = tmp[jj-max_count : jj]
					}
					count = 0

				} else if count_ii != ii {
					count_ii = ii
					max_count = count
					if max_count >= count {
						result = tmp[jj-max_count : jj]
					}
				}
			}
			//на последок проверка если вдруг не успело зайти в раздел tmp[jj] != strs[ii][jj]
			if jj == tmp_len-1 {
				if max_count < count {
					max_count = count
					result = tmp[jj+1-max_count : jj+1]
				}
			}
		}
		count = 0
	}

	return result
}

// #14
// это функция для поиска самого длинный общего префекса(начала строки)
func longestCommonPrefix(strs []string) string {
	if len(strs) == 1 {
		return strs[0]
	}

	var result string = ""
	var count int = 0

	//короткая проверка на наличие пустых строк иначе крашнется программа
	for _, str := range strs {
		if str == result {
			return result
		}
	}

	for ii := 0; ii < len(strs)-1; ii++ {

		//нужно чтобы не делать линих проверок
		if strs[ii][0] == strs[ii+1][0] {

			//выбираем длинну для внутреннего цкила,
			//чтобы в будущем не выходить за пределы строки(стркоа будет массивом у нас):
			var tmp_len int = 0
			if len(strs[ii]) <= len(strs[ii+1]) {
				tmp_len = len(strs[ii])
			} else {
				tmp_len = len(strs[ii+1])
			}

			for jj := 0; jj < tmp_len; jj++ {
				if strs[ii][jj] == strs[ii+1][jj] {
					count++
					if strs[ii][:count] == result {
						break
					}
				} else {
					break
				}
			}
			//проверяем совпадает ли новый результат со старым result если нет то это значит
			// что в коде выше он совсем не сходиться ни 1 буквой
			if (strs[ii][:count] != result) && (result != "") && (len(strs[ii][:count]) == 0) {
				result = ""
				break
			} else if (strs[ii][:count] != result) && (result != "") && (len(strs[ii][:count]) != 0) {
				//это нужно на тот случай если сначало было fl а потом сократилось до f
				result = strs[ii][:count]
				//если наоброт было f и стало fl то код выше не даст сюда попасть и остановит выполнение на f
			} else {
				//данный вариант будет в том случае если все совпадает (это вариант особенно нужен для
				//самого первого прохода по циклу, чтобы значение установилось.
				result = strs[ii][:count]
			}
			count = 0
		} else {
			//если хоть раз самые первые элементы не сошлись то это red флаг
			result = ""
			break
		}

	}

	return result
}

//#1920
//Построение массива из перестановок

func buildArray(nums []int) []int {
	var result []int //создание пустого массива

	for _, number := range nums {
		result = append(result, nums[number])
	}

	return result
}

//1108.
//Удаление IP-адреса

func defangIPaddr(address string) string {

	var result []rune = []rune{}
	addressRuns := []rune(address)
	for ii := 0; ii < len(addressRuns); ii++ {
		if address[ii] == '.' {
			result = append(result, '[', '.', ']')
		} else {
			result = append(result, addressRuns[ii])
		}
	}

	return string(result)
}

// 1512. Количество хороших пар
func numIdenticalPairs(nums []int) int {
	var result int = 0

	for ii := 0; ii < len(nums); ii++ {
		for jj := ii + 1; jj < len(nums); jj++ {
			if nums[ii] == nums[jj] && ii < jj {
				result++
			}
		}
	}
	return result
}

// 2894. Разность делимых и неделимых сумм
func differenceOfSums(n int, m int) int {
	number1 := 0
	number2 := 0
	for ii := 0; ii <= n; ii++ {
		if ii%m != 0 {
			number1 += ii
		} else if ii%m == 0 {
			number2 += ii
		}
	}
	return number1 - number2
}

// 3146. Разница перестановок между двумя строками
func sortedString(s string, t string) string {
	// Преобразуем строку в срез рун
	var runes []rune = []rune(t)

	// Сортируем срез рун
	sort.SliceStable(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})

	// Преобразуем обратно в строку
	return string(runes)
}

func countingPermutations(s string, t string) int {
	var runes []rune = []rune(t)
	var n int = len(runes)
	var swapCount int = 0

	for ii := 0; ii < n-1; ii++ {
		for jj := 0; jj < n-ii-1; jj++ {
			if runes[jj] > runes[jj+1] {
				// Выполняем перестановку
				runes[jj], runes[jj+1] = runes[jj+1], runes[jj]
				swapCount++
				swapCount++
			}
		}
	}

	return swapCount
}

func findPermutationDifference(s string, t string) int {
	var result int = 0
	for ii := range s {
		for jj := range t {
			if s[ii] == t[jj] {
				result += int(math.Abs(float64(ii) - float64(jj)))
			}
		}
	}

	return result
}

//2942. Найдите слова, содержащие иероглиф

func findWordsContaining(words []string, x byte) []int {
	var arr []int
	for ii := 0; ii < len(words); ii++ {
		//отделяем по отдельному симолу в каждой строке:
		for _, jj := range words[ii] {
			if jj == rune(x) {
				arr = append(arr, ii)
				break
			}
		}

	}

	return arr
}

// 1470. Перетасовать массив
func shuffle(nums []int, n int) []int {
	var result []int = make([]int, 2*n) // Создаем массив
	for ii := 0; ii < n; ii++ {
		result[2*ii] = nums[ii]     // Элементы из первой половины
		result[2*ii+1] = nums[ii+n] // Элементы из второй половины
	}
	return result
}

func maximumWealth(accounts [][]int) int {
	var result int = 0
	var tmp_max int = 0
	for ii := 0; ii < len(accounts); ii++ {
		for jj := 0; jj < len(accounts[ii]); jj++ {
			tmp_max += accounts[ii][jj]

		}
		if tmp_max > result {
			result = tmp_max
		}
		tmp_max = 0
	}

	return result
}

func largestLocal(grid [][]int) [][]int {
	var lenght int = len(grid)
	var result [][]int = make([][]int, lenght-2) //создаем итоговую матрицу которая на 2 раза меньше в длинну и ширину

	for ii := range result { // проходимся по строкам и создаем отдельно массив для строк
		result[ii] = make([]int, lenght-2)
	}

	for ii := 0; ii < lenght-2; ii++ { //идем в длинну по срокам в 2 раза меньше как и надо
		for jj := 0; jj < lenght-2; jj++ { //идем в длинну по столбцам в 2 раза меньше как и надо
			for x := 0; x < 3; x++ { // выбираем наибольшее число в квадрате 3 на 3
				for y := 0; y < 3; y++ { // выбираем наибольшее число в квадрате 3 на 3
					// заполняем итоговую матриу максимальными значениями, выбирая самое больше через проверку
					result[ii][jj] = max(result[ii][jj], grid[ii+x][jj+y]) // плюс нужен чтобы этот квадрат двигался как надо
				}
			}
		}
	}
	return result
}

func peakIndexInMountainArray(arr []int) int {
	var left int = 0
	var right int = len(arr) - 1
	for left < right {
		var mid int = left + (right-left)/2
		if arr[mid] < arr[mid+1] {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}
