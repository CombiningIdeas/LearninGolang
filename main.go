//package main

/*
func add(x int, y int) int {
	return x + y
}
func multiply(x int, y int) int {
	return x * y
}
func action(n1 int, n2 int, operation func(int, int) int) {
	result := operation(n1, n2)
	fmt.Println(result)
}
*/

//func main() {
/*var niga int32 = 5
niga++
fmt.Println("hello ффф", niga)
var tmp bool = 3 > 2 && 5 < 10
fmt.Println(tmp)

for i := 1; i < 11; i++ {
	fmt.Println(i * i)
}

var number [10]int
for ii := 0; ii < 10; ii++ {
	number[ii] = ii
}

fmt.Println()

for ii := 0; ii < 10; ii++ {
	fmt.Println(number[ii])
}

fmt.Println()

tmp_a := 10

switch tmp_a {
case 1:
	fmt.Println(1)
case 2, 3, 4:
	fmt.Println(234)
case 5, 6, 7:
	fmt.Println(567)
case 8, 9, 10:
	fmt.Println(8910)
case 1000:
	fmt.Println(90909)
default:
	fmt.Println("значение переменной a не определено")
}

for i := 1; i < 10; i++ {
	for j := 1; j < 10; j++ {
		fmt.Print(i*j, "\t")
	}
	fmt.Println()
}

var users = [5]string{"Tom", "Alice", "Kate"}
for index, value := range users {
	fmt.Println(index, value)
}

var numbers = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
var sum = 0

for _, value := range numbers {
	if value > 10 {
		break // если число больше 4 выходим из цикла
	}
	sum += value
}
fmt.Println("Sum:", sum) // Sum: 10

var digit = [10]int{1, -2, 3, -4, 5, -6, -7, 8, -9, 10}
var sum_digit = 0

for _, value := range digit {
	if value < 0 {
		continue // переходим к следующей итерации
	}
	sum_digit += value
}
fmt.Println("sum_digit:", sum_digit)

action(10, 25, add)    // 35
action(5, 6, multiply) // 30
*/

/*var p *int
y := 10
p = &y
fmt.Println(p)
p = nil
fmt.Println(p)


var ptr = new(int)
fmt.Println(*ptr)
*ptr = 10
fmt.Println(*ptr)





d := 5
fmt.Println("d before:", d) // 5
changeValue(&d)             // изменяем значение
fmt.Println("d after:", d)*/
//array := []int{3, 2, 4}
//target := 6
//
//fmt.Println(twoSum(array, target))

//var str string = "AaA BBB DD"
//
//fmt.Println(findSymbolCount(str))

//fmt.Println("Задание с бутылками:\n", 15%4)
//fmt.Println(numWaterBottles(15, 4))

//}

//func findSymbolCount(text string) map[string]int {
//	text = strings.ToUpper(text)
//	var tmp = map[string]int{}
//	for ii := 0; ii < len(text); ii++ {
//		if val, ok := tmp[string(text[ii])]; ok {
//			tmp[string(text[ii])] = val + 1
//		} else {
//			tmp[string(text[ii])] = 1
//		}
//	}
//	return tmp
//}

//func isPalindrome(x int) bool {
//	var str string = strconv.Itoa(x)
//	for ii := 0; ii < len(str)/2; ii++ {
//		if str[ii] != str[len(str)-1-ii] {
//			return false
//		}
//	}
//	return true
//}

//func numWaterBottles(numBottles int, numExchange int) int {
//	var result int = numBottles
//	for ii := numBottles; ii >= numExchange; ii = (ii / numExchange) + (ii % numExchange) {
//		result += ii / numExchange
//	}
//	return result
//}

/*
func changeValue(x *int) {
	*x = (*x) * (*x)
}
*/

//func twoSum(nums []int, target int) []int {
//	for ii := 0; ii < len(nums); ii++ {
//		for jj := ii + 1; jj < len(nums); jj++ {
//			if (nums[ii] + nums[jj]) == target {
//				var array = []int{ii, jj}
//				return (array)
//			}
//		}
//	}
//	return nil
//}

//func findSymbolCount(text string) {
//
//}
//
//Дана строка. Нужно посчитать количество символов в этой строке
//
//Пример:
//
//Ввод: "Hello"
//Вывод: [h:1, e:1, l:2, o:1]

/*package main

import (
	"fmt"
)

func isValid(s string) bool {
	// Создаем стек для хранения открывающих скобок
	var stack []rune = []rune{}

	// Карта для хранения соответствия закрывающих и открывающих скобок
	// делаем ключем закрывающие скобки, чтобы потом по закрывающим скобкам искать открывающие
	var matchingBrackets map[rune]rune = map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	// Проходимся по каждому символу строки
	for _, char := range s {
		switch char {
		case '(', '{', '[':
			// Если символ - открывающая скобка, добавляем его в стек
			stack = append(stack, char)
		case ')', '}', ']':
			// Если символ - закрывающая скобка, проверяем соответствие последней открывающей скобки в стеке
			// Тут мы не добавляем же закрывающую скобку, поэтмоу можем проверить.
			// Первое это проверка на то, что просто идет закрывающая скобка, без открывающей
			// второе это проверка на соотвествие открытой скобке, такому же типа закрывающей скобки.
			if len(stack) == 0 || stack[len(stack)-1] != matchingBrackets[char] {
				return false // тогда сразу возвращаем значение
			}
			// Удаляем последнюю открывающую скобку из стека
			stack = stack[:len(stack)-1]
		default:
			// Если символ не является скобкой, пропускаем его
			continue
		}
	}

	// Если стек пуст, все открывающие скобки были закрыты корректно
	return len(stack) == 0
}

func main() {
	// Примеры для тестирования
	tests := []string{
		"()",
		"()[]{}",
		"(]",
		"([)]",
		"{[]}",
		"",
		"([]){}",
	}

	for _, test := range tests {
		fmt.Printf("isValid(\"%s\") = %v\n", test, isValid(test))
	}
}*/

//package main
//
//import "fmt"
//
//type Stream interface {
//	read() string
//	write(string)
//	close()
//}
//
//func writeToStream(stream Stream, text string) {
//	stream.write(text)
//}
//func closeStream(stream Stream) {
//	stream.close()
//}
//
//// структура файл
//type File struct {
//	text string
//}
//
//// структура папка
//type Folder struct{}
//
//// реализация методов для типа *File
//func (f *File) read() string {
//	return f.text
//}
//func (f *File) write(message string) {
//	f.text = message
//	fmt.Println("Запись в файл строки", message)
//}
//func (f *File) close() {
//	fmt.Println("Файл закрыт")
//}
//
//// релизация методов для типа *Folder
//func (f *Folder) close() {
//	fmt.Println("Папка закрыта")
//}
//
//func main() {
//
//	// использование указателя на обьекты File and Folder:
//	// так было написанно изначально:
//	var myFile *File = &File{}
//	var myFolder *Folder = &Folder{}
//
//	//Использование на прямую типов данных File and Folder:
//	//Главное не забыть убрать указатели в функиях в этом случае и случае ниже:
//	//// так писать можно в обоих случаях
//	//var myFile File = File{}
//	//var myFolder Folder = Folder{}
//
//	//var myFile Stream = File{} //так писать можно так как реализованы все методы в структуре соотвествующие
//	////  функкциям в интерфейсах
//	//var myFolder Stream = Folder{} // а так писать нельзя, так как методы этого класса не соотвествуют
//	//// функциям интерфейса
//
//	writeToStream(myFile, "hello world")
//	closeStream(myFile)
//	//closeStream(myFolder)     // Ошибка: тип *Folder не реализует интерфейс Stream
//	myFolder.close() // Так можно
//}

//package main
//
//import "fmt"
//
//func main() {
//	var nums []int = []int{1, 1, 1, 2}
//	fmt.Println(removeDuplicates(nums))
//}
//
//func removeDuplicates(nums []int) int {
//	var tmp int = len(nums)
//	var counter int = 0
//	for ii := 0; ii < tmp-counter; ii++ {
//		for jj := ii + 1; jj < tmp-counter; jj++ {
//			if nums[ii] == nums[jj] {
//				nums = append(nums[:jj], nums[jj+1:]...)
//				counter++
//				jj-- //возвращаемся на место удаленного элемента, вдруг там еще есть повторы
//			}
//		}
//	}
//
//	tmp = tmp - counter
//	counter = 0
//	for ii := 0; ii < tmp; ii++ {
//		counter++
//	}
//
//	return counter
//}

//package main
//
//import "fmt"
//
//func main() {
//
//	var nums []int = []int{1, 3, 5, 6}
//	fmt.Println(searchInsert(nums, -10))
//
//}
//
//func searchInsert(nums []int, target int) int {
//	for ii := 0; ii < len(nums); ii++ {
//		if nums[ii] == target {
//			return ii
//		}
//	}
//
//	for ii := 0; ii < len(nums); ii++ {
//		if nums[ii] < target {
//			if (ii + 1) >= len(nums) {
//				return ii + 1
//			} else if nums[ii+1] > target {
//				return ii + 1
//			}
//		}
//
//	}
//	return 0
//}

//package main
//
//import "fmt"
//
//func main() {
//	fmt.Println(convertTemperature(115.55))
//
//}
//
//func convertTemperature(celsius float64) []float64 {
//	return []float64{celsius + 273.15, celsius*1.80 + 32.00}
//}

//package main
//
//import (
//	"fmt"
//)
//
//func main() {
//
//	fmt.Println(scoreOfString("zaz"))
//}
//
//func scoreOfString(s string) int {
//	var summ int = 0
//	for ii := 0; ii < len(s)-1; ii++ {
//		if int(s[ii])-int(s[ii+1]) >= 0 {
//			summ += int(s[ii]) - int(s[ii+1])
//		} else if int(s[ii])-int(s[ii+1]) < 0 {
//			summ += -(int(s[ii]) - int(s[ii+1]))
//		}
//	}
//	return summ
//}

//package main
//
//import (
//	"fmt"
//)
//
//func main() {
//
//	fmt.Println(theMaximumAchievableX(10, 3))
//}
//
//func theMaximumAchievableX(num int, t int) int {
//	return num + t*2
//}

//package main
//
//import (
//	"fmt"
//)
//
//func main() {
//	var nums []int = []int{1, 3, 2, 1}
//	fmt.Println(getConcatenation(nums))
//}

//func getConcatenation(nums []int) []int {
//	var result []int = append(nums, nums...)
//	runtime.GC()
//	return result
//}

//func getConcatenation(nums []int) []int {
//	return append(nums, nums...)
//}

//func getConcatenation(nums []int) []int {
//	var result []int = make([]int, 2*len(nums))
//	for ii := 0; ii < len(nums); ii++ {
//		result[ii] = nums[ii]
//		result[ii+len(nums)] = nums[ii]
//	}
//	runtime.GC()
//	return result
//}

//func getConcatenation(nums []int) []int {
//	var tmp_len int = len(nums) //нужно записать текущую длинну массива,потому что она потом буде тизменяться
//	for ii := 0; ii < tmp_len; ii++ {
//		nums = append(nums, nums[ii])
//	}
//	return nums
//}

//package main
//
//import "fmt"
//
//func main() {
//	fmt.Println(numJewelsInStones("aA", "AfaAdgfaAAa"))
//}
//
//func numJewelsInStones(jewels string, stones string) int {
//	var counter int = 0
//	for jj := 0; jj < len(jewels); jj++ {
//		for ii := 0; ii < len(stones); ii++ {
//			if jewels[jj] == stones[ii] {
//				counter++
//			}
//		}
//	}
//	return counter
//}

//package main
//
//import "fmt"
//
//func main() {
//	var array []string = []string{"--X", "++X", "X++"}
//	fmt.Println(finalValueAfterOperations(array))
//}
//
//func finalValueAfterOperations(operations []string) int {
//	var x int = 0
//	for ii := 0; ii < len(operations); ii++ {
//		if (operations[ii] == "--X") || (operations[ii] == "X--") {
//			x = x - 1
//		}
//		if (operations[ii] == "++X") || (operations[ii] == "X++") {
//			x = x + 1
//		}
//	}
//	return x
//}

//package main
//
//import "fmt"
//
//func main() {
//	var array []int = []int{1, 2, 3, 4}
//	fmt.Println(minimumOperations(array))
//}
//
//func minimumOperations(nums []int) int {
//	var count_operation int = 0
//	for ii := 0; ii < len(nums); ii++ {
//		if nums[ii]%3 != 0 {
//			count_operation++
//			if nums[ii]%3 == 1 {
//				nums[ii] = nums[ii] - 1
//			} else if nums[ii]%3 == 2 {
//				nums[ii] = nums[ii] + 1
//			}
//
//		}
//	}
//
//	return count_operation
//}

//package main
//
//import "fmt"
//
//func main() {
//	fmt.Println(strStr("hello", "lg"))
//}
//
//func strStr(haystack string, needle string) int {
//	//var tmp int = 0
//	//for ii := 0; ii < len(haystack); ii++ {
//	//	for jj := 0; jj < len(needle); jj++ {
//	//		if haystack[ii] == needle[jj] {
//	//			tmp++
//	//
//	//			if ii < len(haystack)-1 { //чтобы ii не выходил за пределы диапазона
//	//				ii++ // если начинает совпадать ii и jj то увеличим соотвественно и ii чтобы проверить не та ли это строка
//	//				//что нам нужна
//	//
//	//			}
//	//		} else {
//	//			if tmp == 0 {
//	//				ii = ii - tmp + 1 // обязательно прописывать логику если это оказалась не та строка
//	//			} else {
//	//				ii = ii - tmp
//	//			}
//	//			tmp = 0
//	//		}
//	//
//	//		if tmp == len(needle) {
//	//			//это нужно если окончания строк совпадают, потому что тогда в другом if ii не увиличивается как должно
//	//			if ii > len(haystack)-1 {
//	//				ii = ii + 2
//	//			} else if ii == len(haystack)-1 {
//	//				ii = ii + 1
//	//			}
//	//			return (ii - len(needle))
//	//		}
//	//	}
//	//}
//	//return -1
//	var tmp int = 0
//	for ii := 0; ii < len(haystack); ii++ {
//		if haystack[ii] == needle[tmp] {
//			tmp++
//			if tmp == len(needle) {
//				return ii - tmp + 1
//			}
//		} else {
//			ii = ii - tmp
//			tmp = 0
//		}
//	}
//	return -1
//}

//// 27. Удалить элемент
//package main
//
//import "fmt"
//
//func main() {
//	var nums []int = []int{3, 2, 3, 3, 2, 3}
//	fmt.Println(removeElement(nums, 3))
//}
//
//func removeElement(nums []int, val int) int {
//	for ii := 0; ii < len(nums); ii++ {
//		if nums[ii] == val {
//			nums = append(nums[:ii], nums[ii+1:]...)
//			ii-- //возвращаем на место наш счетчик
//		}
//	}
//	return len(nums)
//}

//66. Плюс один

//package main
//
//import "fmt"
//
//func main() {
//	nums := []int{9, 9, 9}
//	fmt.Println(plusOne(nums))
//}
//
//func plusOne(digits []int) []int {
//
//	for ii := len(digits) - 1; ii >= 0; ii-- {
//		if digits[ii] < 9 {
//			digits[ii]++
//			return digits
//		}
//		digits[ii] = 0 //если это 9 то увеличиваем постепенно
//	}
//
//	//это на тот случай иил если нет элементов, или на тот слуйчай если они есть, но увеличился разярд
//	//самого большого числа и стал 0, значит нужно в самое начало добавить новый разряд и сделать его равным 1
//	if digits[0] == 0 {
//		digits = append([]int{1}, digits...) //добавляем в начало а не в конец
//		//для этого сначала добавляем новое значение а потомуже сам срез
//	}
//
//	//можно было и так сделать:
//	//digits = make([]int, len(digits)+1)
//	//digits[0] = 1
//
//	return digits
//}

//{DB_DRIVER}://${DB_USR}:${DB_PWD}@${DB_HOST}:${DB_PORT}/${DB_NAME}
//postgresql://postgres:123@127.0.0.1:5532/training?sslmode=disable

//13. Римские числа в целые числа

//package main
//
//import "fmt"
//
//func main() {
//
//	fmt.Println("Hende hock")
//	fmt.Println(romanToInt("DMIXVLIL"))
//}
//
//func romanToInt(s string) int {
//
//	h := map[uint8]int{
//		'I': 1,
//		'V': 5,
//		'X': 10,
//		'L': 50,
//		'C': 100,
//		'D': 500,
//		'M': 1000,
//	}
//
//	var result, penultimate, last int
//	for ii := len(s) - 1; ii >= 0; ii-- {
//		last = h[s[ii]]
//		if last < penultimate {
//			result -= last
//		} else {
//			result += last
//		}
//		penultimate = last
//	}
//
//	return result
//}

//14. Самый длинный общий префекс

//package main
//
//import (
//	"fmt"
//)
//
//func main() {
//
//	//var result string = "rerwerefdh"
//	//result = result[:4]
//	//fmt.Println(result)
//	//fmt.Println("Hende hock")
//	var strs []string = []string{"flow", "flow", "flight", "ftlow", "flow", "flow"}
//	fmt.Println(longestCommonPrefix(strs))
//}

//1920. Построение массива из перестановок

//package main
//
//import "fmt"
//
//func main() {
//
//	fmt.Println("Hello World")
//
//	fmt.Println(buildArray([]int{1, 5, 0, 2, 6, 4, 3, 7}))
//}

//1108.
//Удаление IP-адреса

//package main
//
//import "fmt"
//
//func main() {
//
//	var str string = "255.100.0.1."
//	fmt.Println(defangIPaddr(str))
//}

// 1512.
//Количество хороших пар

//package main
//
//import "fmt"
//
//func main() {
//
//	var str []int = []int{1, 2, 3, 1, 1, 3}
//	fmt.Println(numIdenticalPairs(str))
//}

//2894.
//Разность делимых и неделимых сумм

//package main
//
//import "fmt"
//
//func main() {
//
//	fmt.Println(differenceOfSums(10, 2))
//}

//3146.
//Разница перестановок между двумя строками

//package main
//
//import "fmt"
//
//func main() {
//
//	fmt.Println(findPermutationDifference("abcde", "edbac"))
//}

//2942. Найдите слова, содержащие иероглиф

//package main
//
//import "fmt"
//
//func main() {
//	var words []string = []string{"dffdf", "rerewdsfdf", "dfgse", "waewae", "jyty", "ffs", "ee", "dfd"}
//	fmt.Println(findWordsContaining(words, 'f'))
//}

//1470. Перетасовать массив

//package main
//
//import "fmt"
//
//func main() {
//	var nums []int = []int{2, 5, 1, 3, 4, 7}
//	fmt.Println(shuffle(nums, 3))
//}

//1672. Богатейшее состояние клиента

//package main
//
//import "fmt"
//
//func main() {
//	var nums [][]int = [][]int{{2, 8, 7}, {7, 1, 3}, {1, 9, 5}}
//	//var nums [][]int = [][]int{{1, 2, 3}, {3, 2, 1}}
//	fmt.Println(maximumWealth(nums))
//}

//2373. Наибольшие локальные значения в матрице

//package main
//
//import "fmt"
//
//func main() {
//	var nums [][]int = [][]int{{2, 8, 7, 2}, {7, 1, 3, 11}, {1, 9, 5, 20}}
//	//var nums [][]int = [][]int{{1, 2, 3}, {3, 2, 1}}
//	fmt.Println(largestLocal(nums))
//}

//ООП в GO #1 Пример создания конструктора

//package main
//
//import "fmt"
//
//type Person struct {
//	Name string
//	Age  int
//}
//
//func (p *Person) Initialize(name string, age int) {
//	p.Name = name
//	p.Age = age
//}
//
//func main() {
//	var p *Person = &Person{} // Создание объекта
//	p.Initialize("Alice", 30) // Инициализация объекта
//	fmt.Println(p.Name)       // Output: Alice
//	fmt.Println(p.Age)        // Output: 30
//}

//ООП в GO #2 Пример создания конструктора

//package main
//
//import "fmt"
//
//type Person struct {
//	Name string
//	Age  int
//}
//
//func NewPerson(name string, age int) *Person {
//	return &Person{
//		Name: name,
//		Age:  age,
//	}
//}
//
//func (p *Person) Display() {
//	fmt.Println(p.Name, p.Age)
//}
//
//func main() {
//	var p *Person = NewPerson("Alice", 30)
//	p.Display() // Output: Alice 30
//}

//1603. Проектирование парковочной системы

//package main
//
//import "fmt"
//
//type ParkingSystem struct {
//	array_parking []int
//}
//
//func Constructor(big int, medium int, small int) ParkingSystem {
//	array_parking := []int{0, big, medium, small}      //создаем срез, и сразу нумерация парковачных мест
//	return ParkingSystem{array_parking: array_parking} // поэтому написали 0 на нулевое место, чтобы была нумерация
//}
//
//func (this *ParkingSystem) AddCar(carType int) bool {
//	if this.array_parking[carType] > 0 { //проверка на то есть ли еще места
//		this.array_parking[carType]-- // вычитаем 1 парковачное место из конкретного мест(big, medium или small)
//		return true                   //если места были, и проверка пройдена то возвращаем правду
//	} else {
//		return false //иначе места закончились, тогда возвращаем ложь
//	}
//}
//
//func main() {
//	var obj ParkingSystem = Constructor(1, 1, 0)
//	obj.AddCar(1)
//	obj.AddCar(1)
//	param_1 := obj.AddCar(2)
//	fmt.Println(param_1)
//	fmt.Println(obj)
//}

//1115. Печать FooBar поочередно

//package main
//
//import "fmt"
//
//func printFoo() {
//	fmt.Println("Foo")
//}
//
//func printBar() {
//	fmt.Println("Bar")
//}
//
//type FooBar struct {
//	n   int
//	foo bool
//}
//
//func NewFooBar(n int) *FooBar {
//	return &FooBar{n: n, foo: true}
//}
//
//func (fb *FooBar) Foo(printFoo func()) {
//	for i := 0; i < fb.n; i++ {
//		for !fb.foo {
//		}
//		// printFoo() outputs "foo". Do not change or remove this line.
//		//printFoo()
//		printFoo()
//		fb.foo = false
//	}
//}
//
//func (fb *FooBar) Bar(printBar func()) {
//	for i := 0; i < fb.n; i++ {
//		for fb.foo {
//		}
//		// printBar() outputs "bar". Do not change or remove this line.
//		//printBar()
//		printBar()
//		fb.foo = true
//	}
//}
//
//func main() {
//	var object *FooBar = NewFooBar(2)
//	object.Foo(printFoo)
//	object.Bar(printBar)
//}

//type FooBar struct {
//	n int
//	foo bool
//}
//
//func NewFooBar(n int) *FooBar {
//	return &FooBar{n: n, foo: true}
//}
//
//func (fb *FooBar) Foo(printFoo func()) {
//	for i := 0; i < fb.n; i++ {
//		// printFoo() outputs "foo". Do not change or remove this line.
//		for !fb.foo {
//		}
//		printFoo()
//		fb.foo = false
//
//
//	}
//}
//
//func (fb *FooBar) Bar(printBar func()) {
//	for i := 0; i < fb.n; i++ {
//		// printBar() outputs "bar". Do not change or remove this line.
//		for fb.foo {
//		}
//		printBar()
//		fb.foo = true
//
//	}
//}

//852. Индекс вершины в горном массиве

//package main
//
//import "fmt"
//
//func main() {
//	var number []int = []int{0, 0, 2, 5, 10, 1, 0}
//	fmt.Println(peakIndexInMountainArray(number))
//}

//package main
//
//import "fmt"
//
//func main() {
//
//	for i := 1; i < 7; i++ {
//
//		go func(n int) {
//			result := 1
//			for j := 1; j <= n; j++ {
//				result *= j
//			}
//			fmt.Println(n, "-", result)
//		}(i)
//	}
//	fmt.Scanln()
//	fmt.Println("The End")
//}

//package main
//
//import "fmt"
//
//func main() {
//
//	intCh := make(chan int, 3)
//	intCh <- 10000
//	intCh <- 3234
//	close(intCh) // канал закрыт
//
//	for i := 0; i < cap(intCh); i++ {
//		val, opened := <-intCh
//		if opened == true {
//			fmt.Println(val)
//		} else {
//			fmt.Println("Channel closed!")
//		}
//	}
//}

//package main
//
//import "fmt"
//
//func main() {
//
//	results := make(map[int]int)
//	structCh := make(chan struct{})
//
//	go factorial(5, structCh, results)
//
//	<-structCh // ожидаем закрытия канала structCh
//
//	for i := 0; i <= 6; i++ {
//		fmt.Println(i, " - ", results[i])
//	}
//}
//
//func factorial(n int, ch chan struct{}, results map[int]int) {
//	defer close(ch) // закрываем канал после завершения горутины
//	result := 1
//	for i := 1; i <= n; i++ {
//		result *= i
//		results[i] = result
//	}
//}

//package main
//
//import "fmt"
//
//func main() {
//	intCh := make(chan int)
//
//	go factorial(7, intCh)
//
//	for {
//		num, opened := <-intCh // получаем данные из потока
//		if !opened {
//			break // если поток закрыт, выход из цикла
//		}
//		fmt.Println(num)
//	}
//}
//
//func factorial(n int, ch chan int) {
//	defer close(ch)
//	result := 1
//	for i := 1; i <= n; i++ {
//		result *= i
//		ch <- result // посылаем по числу
//	}
//}

//package main
//
//import (
//	"fmt"
//	"sync"
//)
//
//func main() {
//	intCh := make(chan int)
//
//	go factorial(7, intCh)
//
//	for num := range intCh {
//		fmt.Println(num)
//	}
//}
//
//func factorial(n int, ch chan int) {
//	defer close(ch)
//	result := 1
//	for i := 1; i <= n; i++ {
//		result *= i
//		ch <- result
//	}
//}

//package main
//
//import (
//	"fmt"
//	"sync"
//	"time"
//)
//
//var counter int = 0 //  общий ресурс
//func main() {
//
//	ch := make(chan bool) // канал
//	var mutex sync.Mutex  // определяем мьютекс
//	for i := 1; i < 5; i++ {
//		go work(i, ch, &mutex)
//	}
//
//	for i := 1; i < 5; i++ {
//		<-ch
//	}
//
//	fmt.Println("The End")
//}
//func work(number int, ch chan bool, mutex *sync.Mutex) {
//	mutex.Lock() // блокируем доступ к переменной counter
//	counter = 0
//	for k := 1; k <= 5; k++ {
//		counter++
//		fmt.Println("Goroutine", number, "-", counter)
//	}
//	mutex.Unlock() // деблокируем доступ
//	ch <- true
//}

//package main
//
//import (
//	"fmt"
//	"io"
//	"sync"
//	"time"
//)
//
//func main() {
//	var wg sync.WaitGroup
//	wg.Add(2) // в группе две горутины
//	var work func(id int) = func(id int) {
//		defer wg.Done()
//		fmt.Printf("Горутина %d начала выполнение \n", id)
//		time.Sleep(100 * time.Millisecond)
//		fmt.Printf("Горутина %d завершила выполнение \n", id)
//	}
//
//	// вызываем горутины
//	go work(2)
//	go work(1)
//
//	wg.Wait() // ожидаем завершения обоих горутин
//	fmt.Println("Горутины завершили выполнение")
//}

//package main
//
//import (
//	"fmt"
//	"io"
//)
//
//// Определяем новый тип phoneReader как строку
//type phoneReader string
//
//// Реализуем метод Read для phoneReader
//func (ph phoneReader) Read(p []byte) (int, error) {
//	var count int = 0
//	for ii := 0; ii < len(ph); ii++ {
//		if ph[ii] >= '0' && ph[ii] <= '9' {
//			if count < len(p) { // Проверяем, есть ли еще место в буфере
//				p[count] = ph[ii]
//				count++
//			}
//		}
//	}
//	if count == 0 {
//		return count, io.EOF // Если ничего не записали, возвращаем EOF
//	}
//	return count, nil // Возвращаем количество прочитанных байтов
//}
//
//func main() {
//	phone1 := phoneReader("+1(234)567 9010")
//	phone2 := phoneReader("+2-345-678-12-35")
//
//	// Создаем буфер достаточной длины
//	buffer := make([]byte, len(phone1))
//	n, err := phone1.Read(buffer)
//	if err != nil && err != io.EOF {
//		fmt.Println("Error reading phone1:", err)
//	}
//	fmt.Println(string(buffer[:n])) // 12345679010
//
//	// Создаем новый буфер для phone2
//	buffer = make([]byte, len(phone2))
//	n, err = phone2.Read(buffer)
//	if err != nil && err != io.EOF {
//		fmt.Println("Error reading phone2:", err)
//	}
//	fmt.Println(string(buffer[:n])) // 23456781235
//}

//Я изменил вывод, чтобы отображалось только то количество байтов,
//которое было фактически записано в буфер. Это делается через срез buffer[:n].

//package main
//
//import "fmt"
//
//type phoneWriter struct{}
//
//func (p phoneWriter) Write(bs []rune) (int, error) {
//	if len(bs) == 0 {
//		return 0, nil
//	}
//	for i := 0; i < len(bs); i++ {
//		if bs[i] >= '0' && bs[i] <= '9' {
//			fmt.Print(string(bs[i]))
//		}
//	}
//	fmt.Println()
//	return len(bs), nil
//}
//
//func main() {
//	bytes1 := []rune("+1(234)567 9010")
//	bytes2 := []rune("+2-345-678-12-35")
//
//	writer := phoneWriter{}
//	writer.Write(bytes1)
//	writer.Write(bytes2)
//}

//package main
//
//import (
//	"fmt"
//	"os"
//)
//
//func main() {
//	file, err := os.Create("hello.txt") // создаем файл
//	if err != nil {                     // если возникла ошибка
//		fmt.Println("Unable to create file:", err)
//		os.Exit(1) // выходим из программы
//	}
//	defer file.Close()       // закрываем файл
//	fmt.Println(file.Name()) // hello.txt
//}

//package main
//
//import (
//	"fmt"
//	"os"
//)
//
//func main() {
//	text := "Hello Gold!"
//	file, err := os.Create("hello.txt")
//
//	if err != nil {
//		fmt.Println("Unable to create file:", err)
//		os.Exit(1)
//	}
//	defer file.Close()
//	file.WriteString(text)
//
//	fmt.Println("Done.")
//}

//package main
//
//import (
//	"fmt"
//	"io"
//	"os"
//)
//
//func main() {
//	data := []byte("Hello Bold!")
//	file, err := os.Create("hello.bin")
//	if err != nil {
//		fmt.Println("Unable to create file:", err)
//		os.Exit(1)
//	}
//	defer file.Close()
//	file.Write(data)
//
//	fmt.Println("Done.")
//}

//package main
//
//import (
//	"fmt"
//	"io"
//	"os"
//)
//
//func main() {
//	file, err := os.Open("hello.txt")
//	if err != nil {
//		fmt.Println(err)
//		os.Exit(1)
//	}
//	defer file.Close()
//
//	data := make([]byte, 64) //64- число символов, которое мы прочтем из файла
//	// - плюс это размер среза.
//
//	for { //по сути делается 1 проход по циклу
//		n, erroR := file.Read(data)
//		if erroR == io.EOF { // если конец файла
//			break // выходим из цикла
//		}
//		fmt.Print(string(data[:n]))
//	}
//}

//package main
//
//import (
//	"fmt"
//	"io"
//	"os"
//)
//
//func main() {
//	file, err := os.Open("hello.txt")
//	if err != nil {
//		fmt.Println(err)
//		os.Exit(1)
//	}
//	defer file.Close()
//
//	io.Copy(os.Stdout, file)
//}

//package main
//
//import (
//	"fmt"
//	"io"
//	"os"
//)
//
//type phoneReader string
//
//func (p phoneReader) Read(bs []byte) (int, error) {
//	count := 0
//	for i := 0; i < len(p); i++ {
//		if p[i] >= '0' && p[i] <= '9' {
//			bs[count] = p[i]
//			count++
//		}
//	}
//	return count, io.EOF
//}
//
//func main() {
//	phone1 := phoneReader("+1(234)567 90-10")
//	io.Copy(os.Stdout, phone1)
//	fmt.Println()
//}

//package main
//
//import (
//	"fmt"
//	"os"
//)
//
//func main() {
//	file, err := os.Create("confeve.txt")
//	defer file.Close()
//
//	if err != nil {
//		fmt.Println(err)
//		os.Exit(1)
//	}
//
//	fmt.Fprintln(file, "Сегодня ")
//	fmt.Fprint(file, "хорошая погода")
//
//}

//package main
//
//import (
//	"fmt"
//	"os"
//)
//
//type person struct {
//	name   string
//	age    int32
//	weight float64
//}
//
//func main() {
//	var tom person = person{
//		name:   "Tom",
//		age:    24,
//		weight: 68.5,
//	}
//
//	file, err := os.Create("person.txt")
//	defer file.Close()
//
//	if err != nil {
//		fmt.Println(err)
//		os.Exit(1)
//	}
//
//	fmt.Fprintf(file,
//		"%-10s %-9d %-10.3f\n",
//		tom.name, tom.age, tom.weight)
//}

//package main
//
//import (
//	"fmt"
//	"os"
//)
//
//func main() {
//	fmt.Fprintln(os.Stdout, "hello cold")
//}

//package main
//
//import "fmt"
//
//type person struct {
//	name   string
//	age    int32
//	weight float64
//}
//
//func main() {
//	tom := person{
//		name:   "Tom",
//		age:    24,
//		weight: 68.5,
//	}
//	fmt.Printf("%-10s %-10d %-10.3f\n",
//		tom.name, tom.age, tom.weight)
//	fmt.Print("Hello ")
//	fmt.Println("cold!")
//}

//package main
//
//import (
//	"fmt"
//	"os"
//)
//
//type person struct {
//	name   string
//	age    int32
//	weight float64
//}
//
//func main() {
//	var filename string = string("hello2.txt")
//	writeData(filename)
//	readData(filename)
//}
//
//func writeData(filename string) {
//	// начальные данные
//	var name string = "Tom"
//	var age int = 24
//
//	file, err := os.Create(filename)
//	defer file.Close()
//
//	if err != nil {
//		fmt.Println(err)
//		os.Exit(1)
//	}
//
//	fmt.Fprintln(file, name)
//	fmt.Fprintln(file, age)
//}
//
//func readData(filename string) {
//
//	var name string
//	var age int
//
//	file, err := os.Open(filename)
//	defer file.Close()
//
//	if err != nil {
//		fmt.Println(err)
//		os.Exit(1)
//	}
//
//	fmt.Fscanln(file, &name)
//	fmt.Fscanln(file, &age)
//	fmt.Println(name, age)
//}

//package main
//
//import (
//	"fmt"
//	"io"
//	"os"
//)
//
//type person struct {
//	name   string
//	age    int32
//	weight float64
//}
//
//func main() {
//	filename := "person.txt"
//	writeData(filename)
//	readData(filename)
//}
//
//func writeData(filename string) {
//	// начальные данные
//	var tom person = person{
//		name:   "Tom",
//		age:    24,
//		weight: 68.5,
//	}
//
//	//более подробное обьявление для наглядности и демонстрации типа данных
//	var file *os.File
//	var err error
//
//	//инициализация выше обьявленных перменных
//	file, err = os.Create(filename)
//	defer file.Close() //в конце всех операций закрфыываем файл
//
//	//проверка на ошибку
//	if err != nil {
//		fmt.Println(err)
//		os.Exit(1)
//	}
//
//	// сохраняем данные в файл
//	fmt.Fprintf(file, "%s %d %.2f\n", tom.name, tom.age, tom.weight)
//}
//
////func readData(filename string) {
////
////	// переменные для считывания данных
////	var name string
////	var age int
////	var weight float64
////
////	//обьявление и инициализация перменных
////	file, err := os.Open(filename)
////	defer file.Close() //в конце всех операций закрфыываем файл
////
////	//проверка на ошибку
////	if err != nil {
////		fmt.Println(err)
////		os.Exit(1)
////	}
////
////	// считывание данных из файла
////	_, err = fmt.Fscanf(file, "%s %d %f \n", &name, &age, &weight)
////	//считываем используя ссылки, чтобы передать саму перменную а не копию
////	//по сути мы передаем адреса обьектов, в которые нужно считать данные
////
////	if err != nil {
////		fmt.Println(err)
////		os.Exit(1)
////	}
////
////	//выводи на экран
////	fmt.Printf("\"%s %d %.2f \n", name, age, weight)
////}
//
//// но можно функцию readDatа сделать и такой:
//// тут вместо отдельных примитивных типов мы используем обьект структуры
//func readData(filename string) {
//
//	// переменная для считывания данных
//	tom := person{}
//
//	file, err := os.Open(filename)
//	if err != nil {
//		fmt.Println(err)
//		os.Exit(1)
//	}
//	defer file.Close()
//
//	// считывание данных из файла
//	_, err = fmt.Fscanf(file, "%s %d %f\n", &tom.name, &tom.age, &tom.weight)
//
//	if err != nil {
//		fmt.Println(err)
//		os.Exit(1)
//	}
//	fmt.Printf("%-8s %-8d %-8.2f\n", tom.name, tom.age, tom.weight)
//}

//// Рассмотрим более сложный пример,
//// когда файл содержит набор данных структур person:
//package main
//
//import (
//	"fmt"
//	"io"
//	"os"
//)
//
//type person struct {
//	name   string
//	age    int32
//	weight float64
//}
//
//func main() {
//	filename := "people.txt"
//	writeData(filename)
//	readData(filename)
//}
//
//func writeData(filename string) {
//
//	// начальные данные
//	var people = []person{
//		{"Tom", 24, 68.5},
//		{"Bob", 25, 64.2},
//		{"Sam", 27, 73.6},
//	}
//
//	file, err := os.Create(filename)
//	defer file.Close()
//	if err != nil {
//		fmt.Println(err)
//		os.Exit(1)
//	}
//
//	for _, p := range people {
//		fmt.Fprintf(file, "%s %d %.2f\n", p.name, p.age, p.weight)
//	}
//
//}
//func readData(filename string) {
//
//	var name string
//	var age int
//	var weight float64
//
//	file, err := os.Open(filename)
//	defer file.Close()
//	if err != nil {
//		fmt.Println(err)
//		os.Exit(1)
//	}
//
//	for {
//		_, err = fmt.Fscanf(file, "%s %d %f\n", &name, &age, &weight)
//		if err != nil {
//			if err == io.EOF {
//				break
//			} else {
//				fmt.Println(err)
//				os.Exit(1)
//			}
//		}
//		fmt.Printf("%-8s %-8d %-8.2f\n", name, age, weight)
//	}
//
//}

//package main
//
//import (
//	"fmt"
//	"os"
//)
//
//func main() {
//	var name string
//	var age int
//	fmt.Print("Введите имя: ")
//	fmt.Fscan(os.Stdin, &name)
//
//	fmt.Print("Введите возраст: ")
//	fmt.Fscan(os.Stdin, &age)
//
//	fmt.Println(name, age)
//}

//package main
//
//import (
//	"fmt"
//)
//
//func main() {
//	var name string
//	var age int
//	fmt.Print("Введите имя: ")
//	fmt.Scan(&name)
//	fmt.Print("Введите возраст: ")
//	fmt.Scan(&age)
//
//	fmt.Println(name, age)
//}

//package main
//
//import (
//	"fmt"
//)
//
//func main() {
//	var name string
//	var age int
//	fmt.Print("Введите имя и возраст: ")
//	fmt.Scan(&name, &age)
//	fmt.Println(name, age)
//
//	// альтернативный вариант
//	//fmt.Println("Введите имя и возраст:")
//	//fmt.Scanf("%s %d", &name, &age)
//	//fmt.Println(name, age)
//}

//package main
//
//import (
//	"bufio"
//	"fmt"
//	"io"
//	"os"
//)
//
//func main() {
//	var rows []string = []string{
//		"Hello Go!",
//		"Welcome to Golang",
//	}
//
//	file, err := os.Create("some.txt")
//	var writer *bufio.Writer = bufio.NewWriter(file) // расписано полностью для примера
//	defer file.Close()
//
//	//проверка есть ли ошибка
//	if err != nil {
//		fmt.Println(err)
//		os.Exit(1)
//	}
//
//	for _, row := range rows {
//		writer.WriteString(row)  // запись строки
//		writer.WriteString("\n") // перевод строки
//	}
//
//	writer.Flush() // сбрасываем данные из буфера в файл
//}

//package main
//
//import (
//	"bufio"
//	"fmt"
//	"io"
//	"os"
//)
//
//func main() {
//
//	file, err := os.Open("some.txt")
//	defer file.Close()
//
//	if err != nil {
//		fmt.Println("Unable to open file:", err)
//		os.Exit(1)
//	}
//
//	reader := bufio.NewReader(file)
//
//	for {
//
//		//для наглядности какие типы данных:
//		var line string
//		var err error
//
//		line, err = reader.ReadString('\n')
//
//		if err != nil {
//			//расматриваем ошибку если строка закончилась:
//			if err == io.EOF {
//				break
//			} else { //если иная ошибка то так быть не должно:
//				fmt.Println(err)
//				os.Exit(1)
//			}
//		}
//
//		fmt.Print(line)
//
//	}
//
//}

// Сетевое программирование
// Отправка запросов:

// Сетевое программирование
// Отправка запросов:

// Сетевое программирование
// Отправка запросов:

// Сетевое программирование
// Отправка запросов:

// Сетевое программирование
// Отправка запросов:

// Сетевое программирование
// Отправка запросов:

// Сетевое программирование
// Отправка запросов:

// Сетевое программирование
// Отправка запросов:

// Сетевое программирование
// Отправка запросов:

// Сетевое программирование
// Отправка запросов:

// Сетевое программирование
// Отправка запросов:

// Сетевое программирование
// Отправка запросов:

//package main
//
//import (
//	"fmt"
//	"io"
//	"net"
//	"os"
//)
//
//func main() {
//
//	var httpRequest string
//	httpRequest = "GET / HTTP/1.1\n" + "Host: golang.org\n\n" //здесь мы пишем запрос на то
//	//что хотим получить информацию о хосте, где находится официальный сайт с документацией go
//
//	//для наглядности показываем типы данных:
//	var conn net.Conn
//	var err error
//
//	conn, err = net.Dial("tcp", "golang.org:80")
//	defer conn.Close()
//
//	//обрабатываем ошибку:
//	if err != nil {
//		fmt.Println(err)
//		os.Exit(1)
//	}
//
//	//типичный синтаксис в go, где все максимально сокращается
//	//у нас в if попадает сразу проверка на ошибку, которая(переменная
//	//отвечающая за ошибку) тут и обьявляется и инициализируется сразу
//	//и получается сразу проверяется её значение, это просто гениально
//	//ПРОСТО ГЕНИАЛЬНО!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
//	//отправляем HTTP запрос(httpRequest. Отправляем в виде среза байт),
//	//с помощью метода Write по установленому соединению - conn.
//	if _, err = conn.Write([]byte(httpRequest)); err != nil {
//		//Write([]byte(httpRequest)) - записываеам в уже известный нам метод write
//		//срез байтов. В байты мы преобразуем переменную httpRequest
//
//		//обрабатываем ошибку:
//		fmt.Println(err)
//		os.Exit(1)
//	}
//
//	//копируем данные из одного потока в другой
//	//первый поток conn - это сеть(это поток откуда парсятся данные)
//	// conn - это объект удовлетворяющий интерфейсу io.Reader
//	//Интерфейс io.Reader предназначен для считывания данных. Он имеет следующее определение:
//	//type Reader interface {
//	//    Read(p []byte) (n int, err error)
//	//}
//
//	// os.Stdout - это объект удовлетворяющий интерфейсу io.Writer
//	//io.Writer
//	//Интерфейс io.Writer предназначен для записи в поток. Он определяет метод Write():
//	//type Writer interface {
//	//    Write(p []byte) (n int, err error)
//	//}
//
//	//пример с прошлых глав(название: тандартные потоки ввода-вывода и io.Copy)
//	//n, err = io.Copy(io.Writer, io.Reader)
//	io.Copy(os.Stdout, conn)
//	fmt.Println("Done")
//
//	//_, err = conn.Write([]byte(httpRequest)) - в этом коде,
//	//что описан выше мы используем именно
//	//метод write, который удовлетворяет интерфейсу Writer,
//	// в котором есть метод write который принимает именно срез байтов и ничего другого
//	//а любой байт из среза байтов это символ в кодировки ASCII, именно из таких символов
//	//и состоит срез байтов, и с такими отдельными байтами можно работать сравнивая
//	//с отдельными индексами строки типа string в go, для среза байтов отдельными байтами.
//}

// сервер
//package main
//
//import (
//	"fmt"
//	"net"
//)
//
//func main() {
//
//	message := "Hello, I am a server" // отправляемое сообщение
//
//	listener, err := net.Listen("tcp", ":4545")
//	defer listener.Close()
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//
//	fmt.Println("Server is listening...")
//
//	for {
//		conn, err := listener.Accept()
//
//		if err != nil {
//			fmt.Println(err)
//			return
//		}
//
//		conn.Write([]byte(message))
//		conn.Close()
//	}
//
//}

//клиент
//package main
//import (
//"fmt"
//"os"
//"net"
//"io"
//)
//func main() {
//
//	conn, err := net.Dial("tcp", "127.0.0.1:4545")
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	defer conn.Close()
//
//	io.Copy(os.Stdout, conn)
//	fmt.Println("\nDone")
//}

// Взаимодействие клиента и сервера
// Сервер:
//package main
//
//import (
//	"fmt"
//	"net"
//	"time"
//)
//
//var dict = map[string]string{
//	"red":    "красный",
//	"green":  "зеленый",
//	"blue":   "синий",
//	"yellow": "желтый",
//}
//
//func main() {
//	var listener net.Listener
//	var err error
//
//	listener, err = net.Listen("tcp", ":4545")
//	defer listener.Close()
//
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//
//	fmt.Println("Server is listening...")
//
//	for {
//
//		var conn net.Conn
//		var err error
//		conn, err = listener.Accept()
//
//		if err != nil {
//			fmt.Println(err)
//			conn.Close() //закрываем поток если ошибка
//			continue
//		}
//
//		go handleConnection(conn) // запускаем горутину для обработки запроса
//	}
//}
//
//// обработка подключения
//func handleConnection(conn net.Conn) {
//	defer conn.Close()
//
//	for {
//
//		// считываем полученные в запросе данные
//		var input []byte = make([]byte, (1024 * 4))
//
//		var n int
//		var err error
//		n, err = conn.Read(input)
//
//		//обработка ошибки:
//		if n == 0 || err != nil {
//			fmt.Println("Read error:", err)
//			break
//		}
//
//		//преобразуем из массива байт в строку
//		var source string = string(input[0:n])
//
//		// на основании полученных данных получаем из словаря перевод
//		var target string
//		var err2 bool // в оригинале была переменная ok
//		target, err2 = dict[source]
//
//		if err2 == false { // если данные не найдены в словаре
//			target = "undefined"
//		}
//
//		// выводим на консоль сервера диагностическую информацию
//		fmt.Println(source, "-", target)
//		// отправляем данные клиенту
//		conn.Write([]byte(target))
//	}
//
//	//source - это то что отправил клиент и мы превратили в строку,
//	//для того чтобы по словарю получить значения так как там везде string
//	//target - это итоговая стркоа получения при переводе через словарь,
//	//которую мы возвращаем преобразуя предварительно в массив байтов.
//	//так как именно с байтами работают методы Write и Read
//
//}

//Клиент:
//package main
//import (
//"fmt"
//"net"
//)
//func main() {
//
//	conn, err := net.Dial("tcp", "127.0.0.1:4545")
//	if err != nil {
//		fmt.Println(err)
//

//	}
//	defer conn.Close()
//	for{
//		var source string
//		fmt.Print("Введите слово: ")
//		_, err := fmt.Scanln(&source)
//		if err != nil {
//			fmt.Println("Некорректный ввод", err)
//			continue
//		}
//		// отправляем сообщение серверу
//		if n, err := conn.Write([]byte(source));
//			n == 0 || err != nil {
//			fmt.Println(err)
//			return
//		}
//		// получем ответ
//		fmt.Print("Перевод:")
//		buff := make([]byte, 1024)
//		n, err := conn.Read(buff)
//		if err !=nil{ break}
//		fmt.Print(string(buff[0:n]))
//		fmt.Println()
//	}
//}

// Установка таймаута
// Однако в ряде ситуаций это не лучший способ, особенно когда размер
// передаваемых данных превышает размер буфера. Мы можем точно не знать,
// сколько данных возвратит нам сервер. Поэтому определим следующий код клиента:
//package main
//
//import (
//	"fmt"
//	"net"
//	"net/http"
//	"time"
//)
//
//func main() {
//
//	conn, err := net.Dial("tcp", "192.168.1.157:4545")
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	defer conn.Close()
//
//	for {
//		var source string
//
//		fmt.Print("Введите слово: ")
//
//		_, err := fmt.Scanln(&source)
//		if err != nil {
//			fmt.Println("Некорректный ввод", err)
//			continue
//		}
//
//		// отправляем сообщение серверу
//		if n, err := conn.Write([]byte(source)); n == 0 || err != nil {
//			fmt.Println(err)
//			return
//		}
//
//		// получем ответ
//		fmt.Print("Перевод: ")
//		conn.SetReadDeadline(time.Now().Add(time.Second * 5))
//		//Это значит, что клиент может ожидать данные на чтение от сервера в течении 5
//		//секунд. По истечении этого времени операция чтения генерирует ошибку и
//		//соответственно происходит выход из цикла, где мы пытаемся прочитать данные от
//		//сервера. 5 секунд - довольно большой период, но в начале перед первым
//		//взаимодействием лучше устанавливать период побольше. И после прочтения первых
//		//1024 байт таймаут сбрасывается до 700 миллисекунд. То есть если в течение
//		//последующих 700 милисекунд сервер не пришлет никаких данных, то происходит
//		//выход из цикла и соответственно чтение данных заканчивается.
//		//
//		//Важно понимать роль подобных задержек, так как они позволяют сгенерировать ошибку
//		//при чтении данных. А значит мы можем получить эту ошибку и должным образом
//		//обработать ее, например, выйти из бесконечного цикла. Если бы мы не использовали
//		//установку таймаута, то могла бы сложиться ситуация, когда сервер ожидал данных
//		//от клиента в операции чтения, а клиент ожидал данных от сервера также в операции
//		//чтения. И была бы своего рода блокировка.
//
//		for {
//			buff := make([]byte, 1024)
//
//			n, err := conn.Read(buff)
//			if err != nil {
//				break
//			}
//
//			fmt.Print(string(buff[0:n]))
//
//			conn.SetReadDeadline(time.Now().Add(time.Millisecond * 700))
//		}
//
//		fmt.Print("\n")
//	}
//
//}

//Код сервера остается тем же, что и в прошлой теме:
//package main
//import (
//"fmt"
//"net"
//)
//var dict = map[string]string{
//	"red": "красный",
//	"green": "зеленый",
//	"blue": "синий",
//	"yellow": "желтый",
//}
//
//func main() {
//	listener, err := net.Listen("tcp", ":4545")
//
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	defer listener.Close()
//	fmt.Println("Server is listening...")
//	for {
//		conn, err := listener.Accept()
//		if err != nil {
//			fmt.Println(err)
//			conn.Close()
//			continue
//		}
//		go handleConnection(conn)  // запускаем горутину для обработки запроса
//	}
//}
//// обработка подключения
//func handleConnection(conn net.Conn) {
//	defer conn.Close()
//	for {
//		// считываем полученные в запросе данные
//		input := make([]byte, (1024 * 4))
//		n, err := conn.Read(input)
//		if n == 0 || err != nil {
//			fmt.Println("Read error:", err)
//			break
//		}
//		source := string(input[0:n])
//		// на основании полученных данных получаем из словаря перевод
//		target, ok := dict[source]
//		if ok == false{             // если данные не найдены в словаре
//			target = "undefined"
//		}
//		// выводим на консоль сервера диагностическую информацию
//		fmt.Println(source, "-", target)
//		// отправляем данные клиенту
//		conn.Write([]byte(target))
//	}
//}

// Отправка запросов по HTTP:
//Для отправки запросов в пакете net/http определен ряд функций:
//
//func Get(url string) (resp *Response, err error)
//func Head(url string) (resp *Response, err error)
//func Post(url string, contentType string, body io.Reader)
//(resp *Response, err error)
//func PostForm(url string, data url.Values) (resp *Response, err error)

//Расшифровка:
//Get(): отправляет запрос GET
//Head(): отправляет запрос HEAD
//Post(): отправляет запрос POST
//PostForm(): отправляет форму в запросе POST

// Рассмотрим выполнение самого простого запроса - запроса GET,
// для которого применяется одноименный метод:
//package main
//
//import (
//	"fmt"
//	"io"
//	"net/http"
//	"os"
//)
//
//func main() {
//	resp, err := http.Get("https://google.com")
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	defer resp.Body.Close()
//	for true {
//
//		bs := make([]byte, 1014)
//		n, err := resp.Body.Read(bs)
//		fmt.Println(string(bs[:n]))
//
//		if n == 0 || err != nil {
//			break
//		}
//	}
//}

// Поскольку в данном случае ответ от веб-ресурса все равно
// выводится на консоль, то мы можем сократить код:
//package main
//
//import (
//	"fmt"
//	"io"
//	"net/http"
//	"os"
//)
//
//func main() {
//	resp, err := http.Get("https://google.com")
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	defer resp.Body.Close()
//	io.Copy(os.Stdout, resp.Body)
//}

//http.Client
//Для осуществления HTTP-запросов также может применяться структура http.Client. Чтобы отправить запрос к веб-ресурсу, можно использовать один из ее методов:
//
//func (c *Client) Do(req *Request) (*Response, error)
//func (c *Client) Get(url string) (resp *Response, err error)
//func (c *Client) Head(url string) (resp *Response, err error)
//func (c *Client) Post(url string, contentType string, body io.Reader)
//(resp *Response, err error)
//func (c *Client) PostForm(url string, data url.Values)
//(resp *Response, err error)

// Во многом они аналогичны тем одноименным функциям (за исключением метода Do),
// которые определены в пакете net/http и которые были рассмотрены в прошлой
// теме. Например, выполнение самого простого запроса GET:

//package main
//
//import (
//	"fmt"
//	"io"
//	"net/http"
//	"os"
//	"time"
//)
//
//func main() {
//	client := http.Client{}
//
//	resp, err := client.Get("https://google.com")
//	defer resp.Body.Close()
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//
//	io.Copy(os.Stdout, resp.Body)
//}

//Настройка клиента
//Структура http.Client имеет ряд полей, которые позволяют
//настроить ее поведение:

//Timeout: устанавливает таймаут для запроса
//Jar: устанавливает куки, отправляемые в запросе
//Transport: определяет механиз выполнения запроса

//Установка таймаута:А

//package main
//
//import (
//	"fmt"
//	"io"
//	"net/http"
//	"os"
//	"time"
//)
//
//func main() {
//	client := http.Client{
//		Timeout: 6 * time.Second,
//		//Свойство Timeout представляет объект time.Duration,
//		//и в данном случае оно равно 6 секундам.
//	}
//
//	//для примера показываю тип данных:
//	var resp *http.Response
//	var err error
//
//	resp, err = client.Get("https://google.com")
//	defer resp.Body.Close()
//
//	if err != nil {
//		fmt.Println(err)
//		os.Exit(1)
//	}
//
//	io.Copy(os.Stdout, resp.Body)
//}

//Для отправки объекта Request можно применять метод Do():
//Do(req *http.Request) (*http.Response, error)
//Например:

//package main
//
//import (
//	"fmt"
//	"io"
//	"net/http"
//	"os"
//)
//
//func main() {
//
//	var client *http.Client = &http.Client{}
//
//	var req *http.Request
//	var err error
//
//	req, err = http.NewRequest(
//		"GET", "https://google.com", nil,
//	)
//
//	// добавляем заголовки
//	req.Header.Add("Accept", "text/html")     // добавляем заголовок Accept
//	req.Header.Add("User-Agent", "MSIE/15.0") // добавляем заголовок User-Agent
//
//	var resp *http.Response
//
//	resp, err = client.Do(req)
//	defer resp.Body.Close()
//
//	if err != nil {
//		fmt.Println(err)
//		os.Exit(1)
//	}
//
//	io.Copy(os.Stdout, resp.Body)
//
//}

//PostgreSQL:
//PostgreSQL:
//PostgreSQL:
//PostgreSQL:
//PostgreSQL:
//
//
//
//
//
// Добавление данных
// Для добавления используется метод Exec():
//package main
//
//import (
//	"database/sql"
//	"fmt"
//	_ "github.com/lib/pq"
//	"log"
//)
//
//func main() {
//
//	//открытие соединения с баззой данных:
//
//	connStr := "user=postgres password=123 dbname=postgres sslmode=disable"
//	db, err := sql.Open("postgres", connStr)
//
//	defer db.Close()
//	if err != nil {
//		panic(err)
//	}
//
//	//создаем базу данных и проверяем на ошибки:
//	if _, err = db.Exec("CREATE DATABASE productdb"); err != nil {
//		log.Fatal(err)
//	}
//
//	fmt.Println("Database created successfully")
//
//	_, err = db.Exec(`
//   	CREATE TABLE Products (
//       	id      integer PRIMARY KEY GENERATED BY DEFAULT AS IDENTITY,
//       	model   varchar(30) NOT NULL,
//       	company varchar(30) NOT NULL,
//       	price   integer NOT NULL
//   	)
//	`)
//	if err != nil {
//		panic(err)
//	}
//
//	result, err := db.Exec(`
//		INSERT INTO Products (model, company, price)
//		VALUES ('iPhone X', "Apple", 72000)
//	`)
//	if err != nil {
//		panic(err)
//	}
//
//	fmt.Println(result.LastInsertId()) // не поддерживается
//	fmt.Println(result.RowsAffected()) // количество добавленных строк
//}

// А теперь исправленный код:
//package main
//
//import (
//	"database/sql"
//	"fmt"
//	"log"
//
//	_ "github.com/lib/pq"
//)
//
//func main() {
//	// Подключение к PostgreSQL
//	connStr := "user=postgres password=123 dbname=postgres sslmode=disable"
//	db, err := sql.Open("postgres", connStr)
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer db.Close()
//
//	// Создание базы данных
//	_, err = db.Exec("CREATE DATABASE productdb")
//	if err != nil {
//		log.Fatal(err)
//	}
//	fmt.Println("Database created successfully")
//
//	// Подключение к новой базе данных
//	productDbConnStr := "user=postgres password=123 dbname=productdb sslmode=disable"
//	productDb, err := sql.Open("postgres", productDbConnStr)
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer productDb.Close()
//
//	// Создание таблицы Products
//	_, err = productDb.Exec(`
//		CREATE TABLE products (
//			id      SERIAL PRIMARY KEY,
//			model   VARCHAR(30) NOT NULL,
//			company VARCHAR(30) NOT NULL,
//			price   INTEGER NOT NULL
//		)
//	`)
//	if err != nil {
//		log.Fatal(err)
//	}
//	fmt.Println("Table created successfully")
//
//	// Вставка данных в таблицу Products
//	result, err := productDb.Exec(`
//		INSERT INTO products (model, company, price)
//		VALUES ('iPhone X', 'Apple', 72000)
//	`)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	// Получение количества добавленных строк
//	rowsAffected, err := result.RowsAffected()
//	if err != nil {
//		log.Fatal(err)
//	}
//	fmt.Printf("Rows affected: %d\n", rowsAffected)
//}

//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//Если нам обязательно нужно получить id добавленного объекта, то мы можем использовать метод db.QueryRow(),
//который выполняет запрос и возвращает определенный объект:
//var id int
//db.QueryRow("insert into Products (model, company, price) values ('Mate 10 Pro', $1, $2) returning id",
//    "Huawei", 35000).Scan(&id)
//fmt.Println(id)

//В само sql выражение вводится подвыражение "returning id".
//И с помощью метода Scan() полученное значение считывается в переменную id.
//
//
//
//
//
//

//Получение данных:
//Для получения данных применяется метод db.Query(), который возващает набор строк, либо db.QueryRow(), который возвращает одну строку:

//package main
//
//import (
//	"database/sql"
//	"fmt"
//	_ "github.com/lib/pq"
//)
//
//type product struct {
//	id      int
//	model   string
//	company string
//	price   int
//}
//
//func main() {
//
//	connStr := "user=postgres password=123 dbname=productdb sslmode=disable"
//	db, err := sql.Open("postgres", connStr)
//	defer db.Close()
//	if err != nil {
//		panic(err)
//	}
//
//	//выводим все колонки на экран:
//	rows, err := db.Query("select * from Products")
//	defer rows.Close()
//	if err != nil {
//		panic(err)
//	}
//
//	var products []product = []product{} //создается срез обьектов(или же срез структуры, мы создали срез переменных
//	//с типом данных product) - а это и есть срез структур
//
//	for rows.Next() == true {
//		//создаем еще 1 обьект(временный) для записи в срез обьектов
//		var p product = product{}
//
//		//заполняем поля структуры через ссылку на них самих, это нужно чтобы не передовались и незаполнялись копии:
//		if err := rows.Scan(&p.id, &p.model, &p.company, &p.price); err != nil {
//			fmt.Println(err)
//			continue
//		}
//
//		// в цикле пока есть строке в таблице(указателе на колонки), мы перебираем их
//		// и записываем в другой обьект,
//		products = append(products, p)
//	}
//
//	//выводим на экран через цикл все строки, которые записали в структуру
//	for _, p := range products {
//		fmt.Println(p.id, p.model, p.company, p.price)
//	}
//}

//Обновление
// Для обновления данных применяется метод Exec:
//package main
//
//import (
//	"database/sql"
//	"fmt"
//	_ "github.com/lib/pq"
//)
//
//func main() {
//
//	connStr := "user=postgres password=mypass dbname=productdb sslmode=disable"
//	db, err := sql.Open("postgres", connStr)
//	if err != nil {
//		panic(err)
//	}
//	defer db.Close()
//
//	// обновляем строку с id=1
//	result, err := db.Exec("update Products set price = $1 where id = $2", 69000, 1)
//	if err != nil {
//		panic(err)
//	}
//	fmt.Println(result.RowsAffected()) // количество обновленных строк
//}

// Удаление
// Для удаления также применяется метод Exec:
//package main
//
//import (
//	"database/sql"
//	"fmt"
//	_ "github.com/lib/pq"
//)
//
//func main() {
//
//	connStr := "user=postgres password=mypass dbname=productdb sslmode=disable"
//	db, err := sql.Open("postgres", connStr)
//	defer db.Close()
//	if err != nil {
//		panic(err)
//	}
//
//	// удаляем строку с id=2
//	result, err := db.Exec("delete from Products where id = $1", 2)
//	if err != nil {
//		panic(err)
//	}
//	fmt.Println(result.RowsAffected()) // количество удаленных строк
//}

//package main
//
//import "fmt"
//
////Builder
//
//type ComputerBuilderI interface {
//	CPU(string) ComputerBuilderI
//	RAM(int) ComputerBuilderI
//	MB(string) ComputerBuilderI
//
//	Build() Computer
//}
//
//type Computer struct {
//	CPU string
//	RAM int
//	MB  string
//}
//
//// для этой структуры будем реализовывать методы интерфейса ComputerBuilderI
//type computerBuilder struct {
//	cpu string
//	ram int
//	mb  string
//}
//
//// constructor
//func NewComputerBuilder() ComputerBuilderI {
//	return &computerBuilder{}
//}
//
//// реализовываем методы:
//func (b *computerBuilder) CPU(val string) ComputerBuilderI {
//	b.cpu = val
//	return b
//}
//func (b *computerBuilder) RAM(val int) ComputerBuilderI {
//	b.ram = val
//	return b
//}
//func (b *computerBuilder) MB(val string) ComputerBuilderI {
//	b.mb = val
//	return b
//}
//
//func (b *computerBuilder) Build() Computer {
//	return Computer{
//		CPU: b.cpu,
//		RAM: b.ram,
//		MB:  b.mb,
//	}
//}
//
//func main() {
//	computerBuilder := NewComputerBuilder()
//	computer := computerBuilder.CPU("core i3").RAM(8).MB("gigabity").Build()
//	fmt.Println(computer, computerBuilder)
//}
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//

//package main
//
//import "fmt"
//
////Builder
//
//type ComputerBuilderI interface {
//	CPU(string) ComputerBuilderI
//	RAM(int) ComputerBuilderI
//	MB(string) ComputerBuilderI
//
//	Build() Computer
//}
//
//type Computer struct {
//	CPU string
//	RAM int
//	MB  string
//}
//
//// для этой структуры будем реализовывать методы интерфейса ComputerBuilderI
//type computerBuilder struct {
//	cpu string
//	ram int
//	mb  string
//}
//
//// constructor - возвращает computerBuilderI - интерфейс
//func NewComputerBuilder() *computerBuilder {
//	return &computerBuilder{}
//}
//
//// мы в методах возвращали интерфейс, чтобы потом через точку можно было все написать:
//// computerBuilder.CPU("core i3").RAM(8).MB("gigabity").Build()
//// реализовываем методы:
//func (b *computerBuilder) CPU(val string) ComputerBuilderI {
//	b.cpu = val
//	return b
//}
//func (b *computerBuilder) RAM(val int) ComputerBuilderI {
//	b.ram = val
//	return b
//}
//func (b *computerBuilder) MB(val string) ComputerBuilderI {
//	b.mb = val
//	return b
//}
//
//func (b *computerBuilder) Build() Computer {
//	return Computer{
//		CPU: b.cpu,
//		RAM: b.ram,
//		MB:  b.mb,
//	}
//}
//
////
////
////
////
////
////
////
//
////Когда вы создаете встроенную структуру, вы можете использовать поля встроенной структуры
////как если бы они были полями самой структуры. Это означает, что вы можете достучаться к полям
////computerBuilder структуры через officeComputerBuilder структуру, как если бы они были ее
////собственными полями. Но при этом если до этого поля прошлой структуры имели какие либо значения
////то в новой структуре они бы не сохранились бы(если бы там было тоже самое что в методе Build()
////так как мы там принимаем другой обьект, обьект типа officeComputerBuilder, а не computerBuilder
//
////но если бы этот метод так же принимал бы тот же тип данных (тоесть не изменялся бы)
////то тогда было бы возможно использовать тот же метод Build() и получить те же дефолтные поля
////которые были бы установлены в структуре computerBuilder в методе Build()
//
//// В этом примере, officeComputerBuilder структура содержит поля computerBuilder структуры,
//// что позволяет ей наследовать поведение computerBuilder структуры. Это означает,
//// что officeComputerBuilder структура может использовать методы computerBuilder структуры,
//// такие как CPU, RAM и MB, как если бы они были ее собственными методами.
//
////Например, в коде officeComputerBuilder := NewOfficeComputerBuilder();
////officeComputer := officeComputerBuilder.RAM(16).Build(), вы используете метод RAM
////из computerBuilder структуры, как если бы он был методом officeComputerBuilder структуры.
//
//type officeComputerBuilder struct {
//	computerBuilder
//}
//
//// нужен еще конструктор:
//func NewOfficeComputerBuilder() ComputerBuilderI {
//	return &officeComputerBuilder{}
//}
//
//func (b *officeComputerBuilder) Build() Computer {
//	//описываем дефолтное создание компьютера, которое потом можно будет изменить:
//	if b.cpu == "" {
//		b.cpu = "pentium 3"
//	}
//	if b.ram == 0 {
//		b.ram = 2
//	}
//	if b.mb == "" {
//		b.mb = "asrock"
//	}
//
//	return Computer{
//		CPU: b.cpu,
//		RAM: b.ram,
//		MB:  b.mb,
//	}
//}
//
////
////
////
////
////
////
////
//
//func main() {
//	var compBuilder *computerBuilder = NewComputerBuilder()
//	var computer Computer = compBuilder.CPU("core i3").RAM(8).MB("gigabity").Build()
//	fmt.Println(computer, compBuilder)
//
//	//вторым параметром выступает структура с которой
//	// мы изначально работали и тут вообще тип данных будет *computerBuilder(конкретная реализация)
//	//просто программа намерено использует абстрактный тип данных по умолчанию, а не конкретный,
//	//потому что в конструкторе у нас указан асбтрактных тип данных
//	//который реализует сам интерфейс, поэтому мы можем руками это поменять и написать
//	// *computerBuilder вместо ComputerBuilderI, но тогда придется и поменять в самом конструкторе
//	//тип данных с ComputerBuilderI на *computerBuilder, и тут уже все же будут различия, смотря
//	//для чего нам это нужно. вообще рекомендуется указывать интерфейс так как это
//	//более гибкий поддход
//
//	//Помимо своего компьютера делаем еще и офисные компьютер:
//	//вообще у нас тип данных officeComputerBuilder, но программа пишет асбтрактный тип данных
//	//для нашего обьекта тоесть ComputerBuilderI, так как она не знает какая структура будет
//	//соотвествовать данном интерфейсу во время выполнения программы, и если мы не используем
//	//указатели то работаем с копией структуры, и если работая с копиями
//	//Из копии сразу структуры данные сразу передать в нужную нам структуру как написано в коде выше
//	//то тогда скомпилируется копия структуры, которая будет удовлетворять нашему интерфейсу
//
//	//
//	//
//
//	//тоже самое как и в простом примере если через поле поменять значение то оно изменится даже без использования указателя
//	computer.CPU = "Ryzen 9 5500fh"
//	compBuilder.cpu = "Ryzen 9 5500fh"
//	//но если сделать так без указателя то значени не изменится:
//	compBuilder.CPU("Ryzen 9 5500fh") //так как это изменение полей через методы в go без указателей, а без ссылки конечное
//	//значение не изменится, так как результат вернется другим это да, но конкретно то что мы изменяли по итогу не изменится,
//	//так как без указателя создастся копия передаваемой переменной и итоговые изменения не будут видны в
//	//оригинальной(нужной нам не в копие) структуре, так как нет указателя на эти поля, которые должны были бы изменится
//	//изменится чисто созданная во время компиляции кода копия этих полей, а соотвествено и копия структуры, где
//	//и будут видны изменения
//	fmt.Println(computer, compBuilder)
//
//	//
//	//
//
//	officeCompBuilder := NewOfficeComputerBuilder()
//	officeCompBuilder.CPU("Ryzen 7500f").RAM(16).MB("gigabity")
//	officeComputer := officeCompBuilder.Build()
//	fmt.Println(officeComputer, officeCompBuilder)
//	//а если как тут использование метода и передачи данных из одной структуры в другую делать
//	//без указателей, то результат методов никуда не запишется, так как была изменена копия,
//	//которая автоматически создавалась без указателей, и результат никуда не успели перенести
//	//и при переносе данных из одной структуры в другую мы не будем знать как изменялась копия
//	//ведь данные копии структур были потеряны(так как мы работаем вообще с указателем на обьект
//	//структуры, и поля структуры мы изменяем через обьект(этой же соотвествено структуры)
//	//с именем b и поэтому во время вызова метода Build все будет зависить от того что было написано
//	//в самой структуре а не при вызове методов обьекта officeCompBuilder.
//}

//
//
//
//
//
//
//
//
//
//
//
//
//package main
//
//import "fmt"
//
////Builder
//
//type ComputerBuilderI interface {
//	CPU(string) ComputerBuilderI
//	RAM(int) ComputerBuilderI
//	MB(string) ComputerBuilderI
//
//	Build() Computer
//}
//
//type Computer struct {
//	CPU string
//	RAM int
//	MB  string
//}
//
//type computerBuilder struct {
//	cpu string
//	ram int
//	mb  string
//}
//
//func NewComputerBuilder() *computerBuilder {
//	return &computerBuilder{}
//}
//
//func (b *computerBuilder) CPU(val string) ComputerBuilderI {
//	b.cpu = val
//	return b
//}
//func (b *computerBuilder) RAM(val int) ComputerBuilderI {
//	b.ram = val
//	return b
//}
//func (b *computerBuilder) MB(val string) ComputerBuilderI {
//	b.mb = val
//	return b
//}
//
//func (b *computerBuilder) Build() Computer {
//	if b.cpu == "" {
//		b.cpu = "pentium 3"
//	}
//	if b.ram == 0 {
//		b.ram = 2
//	}
//	if b.mb == "" {
//		b.mb = "asrock"
//	}
//	return Computer{
//		CPU: b.cpu,
//		RAM: b.ram,
//		MB:  b.mb,
//	}
//}
//
//type officeComputerBuilder struct {
//	computerBuilder
//}
//
//// нужен еще конструктор:
//func NewOfficeComputerBuilder() ComputerBuilderI {
//	return &officeComputerBuilder{}
//}
//
//func main() {
//	var compBuilder *computerBuilder = NewComputerBuilder()
//	var computer Computer = compBuilder.CPU("core i3").RAM(8).MB("gigabity").Build()
//	fmt.Println(computer, compBuilder) //вторым параметром выступает структура с которой
//
//	officeCompBuilder := NewOfficeComputerBuilder()
//	fmt.Println(officeCompBuilder.Build())
//	officeCompBuilder.CPU("Ryzen 7500f").RAM(16).MB("gigabity")
//	officeComputer := officeCompBuilder.Build()
//	fmt.Println(officeComputer, officeCompBuilder)
//
//}

// разбор данного кода на простом примере:
//package main
//
//import "fmt"
//
//type person struct {
//	name string
//	age  int
//}
//
//func (p person) print() {
//	fmt.Println("Имя:", p.name)
//	fmt.Println("Возраст:", p.age)
//}
//
//// если тут не передавать указатель то поле не поменяется
//func (p *person) eat(meal string) {
//	p.name = "Timati"
//	fmt.Println(p.name, "ест", meal)
//}
//
//func main() {
//
//	var tom123 = person{name: "Tom", age: 24}
//	tom123.print()
//	tom123.eat("борщ с капустой, но не красный")
//	////Но если после даже изменения через указатель поменять поле просто через значение то оно изменится:
//	//tom.name = "Piter"
//	fmt.Println(tom123.name)
//}

//
//
//
//
//
//
//
//
//
//Директор в патерне билдер:

//package main
//
//import (
//	"fmt"
//)
//
//func main() {
//	//Обычные компьютеры:
//	var computerBuilder BuilderI = NewBuilder()
//	computerBuilder = computerBuilder.CPU("core_i9").RAM(32).MB("gigabyte")
//	director := NewDirector(computerBuilder)
//
//	compA := director.BuildComputer()
//	fmt.Println(compA) //{core_i9 32 gigabyte }
//
//	//
//	//
//	//
//
//	//Офисные компьютеры:
//	var officeCompBuilder BuilderI = NewOfficeComputerBuilder()
//	officeCompBuilder = officeCompBuilder.RAM(32)
//	director.SetBuilder(officeCompBuilder)
//	compB := director.BuildComputer()
//	fmt.Println(compB) //{pentium III 32 asrock logitech}
//
//}

//патерн стратегия:

//package main
//
//func main() {
//	product := "vehicle"
//	payWay := 3
//
//	var payment Payment
//	//выбор способа оплаты:
//	switch payWay {
//	case 1:
//		payment = NewCardPayment("12345", "12345")
//	case 2:
//		payment = NewPayPalPayment()
//	case 3:
//		payment = NewQIWIPayment()
//	}
//
//	processOrder(product, payment)
//}
//
//func processOrder(product string, payment Payment) {
//	// ... implementation
//	err := payment.Pay()
//	if err != nil {
//		return
//	}
//}
//
//// ----
//
//type Payment interface {
//	Pay() error
//}
//
//// ----
//
//type cardPayment struct {
//	cardNumber, cvv string
//}
//
//func NewCardPayment(cardNumber, cvv string) Payment {
//	return &cardPayment{
//		cardNumber: cardNumber,
//		cvv:        cvv,
//	}
//}
//
//func (p *cardPayment) Pay() error {
//	// ... implementation
//	return nil
//}
//
//type payPalPayment struct {
//}
//
//func NewPayPalPayment() Payment {
//	return &payPalPayment{}
//}
//
//func (p *payPalPayment) Pay() error {
//	// ... implementation
//	return nil
//}
//
//type qiwiPayment struct {
//}
//
//func NewQIWIPayment() Payment {
//	return &qiwiPayment{}
//}
//
//func (p *qiwiPayment) Pay() error {
//	// ... implementation
//	return nil
//}

//package main
//
//import (
//	"fmt"
//	"net/http"
//	"time"
//)
//
//func main() {
//	urls := []string{
//		"http://youtube.com",
//		"http://google.com",
//		"http://github.com",
//		"http://instagram.com",
//		"http://medium.com",
//	}
//
//	// создаем канал для синхронизации горутин
//	ch := make(chan string)
//
//	for _, url := range urls {
//		go func() {
//			doHTTP(url)
//			// отправляем сигнал о завершении горутины
//			ch <- "connection established"
//		}()
//	}
//
//	// ожидаем завершения всех горутин
//	for url := range urls {
//		msg, open := <-ch
//		if !open {
//			break
//		}
//		fmt.Println(msg, "; number -", url)
//		//fmt.Println(msg, "; number -", urls[url])
//	}
//}
//
//func doHTTP(url string) {
//	t := time.Now()
//
//	resp, err := http.Get(url)
//	if err != nil {
//		fmt.Printf("Failed to get <%s>: %s\n", url, err.Error())
//		return
//	}
//
//	defer resp.Body.Close()
//
//	fmt.Printf("<%s> - Status Code [%d] – Latency %d ms\n", url, resp.StatusCode, time.Since(t).Milliseconds())
//}

//// А чтобы решить проблему раннего завершения программы, мы можем использовать горутины
//package main
//
//import (
//	"fmt"
//	"net/http"
//	"time"
//)
//
//func main() {
//	urls := []string{
//		"http://google.com",
//		"http://youtube.com",
//		"http://github.com",
//		"http://instagram.com",
//		"http://medium.com",
//	}
//
//	for _, url := range urls {
//		go doHTTP(url)
//	}
//
//	time.Sleep(2 * time.Second)
//}
//
//func doHTTP(url string) {
//	t := time.Now()
//
//	resp, err := http.Get(url)
//	if err != nil {
//		fmt.Printf("Failed to get <%s>: %s\n", url, err.Error())
//	}
//
//	defer resp.Body.Close()
//
//	fmt.Printf("<%s> - Status Code [%d] – Latency %d ms\n", url, resp.StatusCode, time.Since(t).Milliseconds())
//
//}

//package main
//
//import (
//	"fmt"
//	"net/http"
//	"sync"
//	"time"
//)
//
//func main() {
//	urls := []string{
//		"http://google.com",
//		"http://youtube.com",
//		"http://github.com",
//		"http://instagram.com",
//		"http://medium.com",
//	}
//
//	var wg sync.WaitGroup
//
//	for _, url := range urls {
//		wg.Add(1)
//
//		go func() {
//			doHTTP(url)
//			wg.Done()
//		}()
//	}
//
//	wg.Wait()
//}
//
//func doHTTP(url string) {
//	t := time.Now()
//
//	resp, err := http.Get(url)
//	if err != nil {
//		fmt.Printf("Failed to get <%s>: %s\n", url, err.Error())
//	}
//
//	defer resp.Body.Close()
//
//	fmt.Printf("<%s> - Status Code [%d] – Latency %d ms\n", url, resp.StatusCode, time.Since(t).Milliseconds())
//
//}

// в этом случае мы параллельно запускаем горутины и отправляем данные в главную горутину main,
// но в main происходит последовательное чтение из каналов, а не параллельно, как должно быть
// и как было в прошлом случае, мы считывали через цикл, где после чтения одной горутины сразу
// начиналось чтение другой горутины, и не важно какой по счету, а важно какой по скорости,
// какая была быстрее та и запускалась
//package main
//
//import (
//	"fmt"
//	"time"
//)
//
//func main() {
//	message1 := make(chan string)
//	message2 := make(chan string)
//
//	//Первая горутина
//	go func() {
//		for {
//			time.Sleep(250 * time.Millisecond)
//			message1 <- "Прошла четверть секунды"
//		}
//	}()
//
//	//Вторая горутина
//	go func() {
//		for {
//			time.Sleep(time.Second)
//			message2 <- "Прошла 1 секунда"
//		}
//	}()
//
//	for {
//		select {
//		case msg := <-message1:
//			fmt.Println(msg)
//		case msg := <-message2:
//			fmt.Println(msg)
//		}
//	}
//}

// Программа для перевода чисел в различные системы счисления:

//package main
//
//import (
//	"fmt"
//)
//
//// Функция для перевода десятичного числа в двоичное
//func decimalToAny(n int, notation int) string {
//
//	if n == 0 { //проверяем не равно ли 0, иначе если убрать проверку
//		//то потом будет ошибка, так как будет пустая строка конвертироваться в число
//		return "0" //
//	}
//
//	//Наш словарь состоящий из строки:
//	dictionary := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
//	//Это наше число n, будет в нужной нам системе счисления:
//	var result string = ""
//	for n > 0 {
//		remainder := n % notation
//		//преобразуем в строку так как иначе это будет считаться символом: UTF8
//		result = string(dictionary[remainder]) + result // добавляем остаток в начало строки
//		n = n / notation
//	}
//
//	return result
//}
//
//func main() {
//	fmt.Println("Практическая №2")
//
//	var number, notation int
//	fmt.Print("Введите десятичное число: ")
//	fmt.Scan(&number)
//	fmt.Print("Введите желаемую систему счисления: ")
//	fmt.Scan(&notation)
//
//	//возвращаем именно строку так как может быть и система счисления
//	//больше чем 9, тогда будут нужны буквы:
//	var result string = decimalToAny(number, notation)
//	fmt.Printf("Двоичное представление числа %d: %s\n", number, result)
//}

// Программа для перевода чисел из различных системы счисления в десятичную систему счисления с реализацией
//4 основных действий(сложение, вычитание, умножение, деление):

//package main
//
//import (
//	"fmt"
//	"math"
//)
//
//// Функция для перевода любого числа в десятичное
//func AnyToDecimal(n string, notation int) float64 {
//
//	if n == "0" { //проверяем не равно ли 0, иначе если убрать проверку
//		//то потом будет ошибка, так как будет пустая строка конвертироваться в число
//		return 0 //
//	}
//
//	//Наш словарь состоящий из строки:
//	dictionary := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
//
//	// Создаем map с ключом типа string и значением типа int
//	indexMap := make(map[string]int)
//
//	// Заполняем map, используя цикл
//	for ii, symbol_rune := range dictionary {
//		indexMap[string(symbol_rune)] = ii
//	}
//
//	//Это наше число n, будет в десятичной системе счисления:
//	var result float64 = 0
//
//	//начинаем перебор строки, на отдельные символы(данный вид цикла разложжит на руны нашу байтовую строку)
//	for ii, symbol_rune := range n {
//		result += float64(indexMap[string(symbol_rune)]) * (math.Pow(float64(notation), float64(len(n)-ii-1))) //так как умножаем длину минус 1
//		// на самый первый элемент n
//	}
//
//	return result
//}
//
//// сложение
//func addition(x float64, y float64) float64 {
//	return x + y
//}
//
//// вычитание
//func subtraction(x float64, y float64) float64 {
//	return x - y
//}
//
//func main() {
//	fmt.Println("Практическая №2. Пункт 4.")
//
//	var number1, number2, operator string // так как может быть система счисления больше 10:
//	var notation1, notation2 int
//	fmt.Print("Введите свое число: ")
//	fmt.Scan(&number1)
//	fmt.Print("Введите систему счисления для введенного числа: ")
//	fmt.Scan(&notation1)
//
//	fmt.Print("Введите свое число: ")
//	fmt.Scan(&number2)
//	fmt.Print("Введите систему счисления для введенного числа: ")
//	fmt.Scan(&notation2)
//
//	fmt.Println("Введите знак между числами: ")
//	fmt.Scan(&operator)
//
//	var number1_decimal float64 = AnyToDecimal(number1, notation1)
//	var number2_decimal float64 = AnyToDecimal(number2, notation2)
//
//	if operator == "+" {
//		fmt.Println(addition(number1_decimal, number2_decimal))
//	} else if operator == "-" {
//		fmt.Println(subtraction(number1_decimal, number2_decimal))
//	} else {
//		fmt.Println("Непредусмотренная операция!")
//	}
//}

//
//Практическая №3. Пункт 6.1
//Практическая №3. Пункт 6.1
//Практическая №3. Пункт 6.1:
//
//package main
//
//import (
//	"fmt"
//	"math"
//	"strconv"
//	"strings"
//)
//
//func main() {
//
//	fmt.Println("Практическая №3. Пункт 6.1: Разработать программу, моделирующую выполнение основных арифметических " +
//		"операций: умножение и деление над числами, представленными в двоичной системе счисления.")
//
//	var number1_str, number2_str string
//
//	fmt.Println("Введите два произвольных двоичных числа(через пробел): ")
//	fmt.Scan(&number1_str, &number2_str)
//	number1_int := BinaryStringToFloat(number1_str)
//	number2_int := BinaryStringToFloat(number2_str)
//	fmt.Println("Само числа А в десятичной системе счисления: ", number1_int)
//	fmt.Println("Само числа B в десятичной системе счисления: ", number2_int)
//	res_number1, res_number2 := Operation(number1_int, number2_int)
//
//	fmt.Println("Результат: ")
//
//	fmt.Printf("Умножение чисел в двоичной системе счисления(a * b): %s; \nДеление чисел в двоичной системе счисления(a / b): %s \n",
//		FloatToBinaryString(res_number1), FloatToBinaryString(res_number2))
//
//}
//
//func Operation(number1 float64, number2 float64) (float64, float64) {
//
//	result1 := number1 * number2
//	result2 := number1 / number2
//
//	fmt.Println("Результат умножения чисел в десятичной ситеме счисления(a * b): ", result1)
//	fmt.Println("Результат умножения чисел в десятичной ситеме счисления(a * b): ", result2)
//
//	return result1, result2
//}
//
//// Функция для преобразования двоичной строки в десятичную строку с типом float64
//func BinaryStringToFloat(binaryStr string) float64 {
//
//	parts := strings.Split(binaryStr, ".")
//	var intPart float64
//	var fracPart float64
//	var fracPartFlag bool
//
//	//проверка на существование дробной части числа:
//	for ii, _ := range parts {
//		//fmt.Println(ii) //будет вывод 0 и 1
//		if ii == 1 {
//			fracPartFlag = true
//		}
//	}
//
//	// Преобразуем целую часть с помощью ParseInt
//	if true {
//		intValue, err := strconv.ParseInt(parts[0], 2, 64)
//		if err != nil {
//			return 0
//		}
//		intPart = float64(intValue)
//	}
//
//	// Преобразуем дробную часть
//	if fracPartFlag == true {
//		for i, char := range parts[1] {
//			if char == '1' {
//				fracPart += 1 / math.Pow(2, float64(i+1))
//			}
//		}
//	} else {
//		fracPart = 0
//	}
//
//	return intPart + fracPart
//}
//
//// Функция для преобразования десятичной строки с типом float64 в двоичную строку
//func FloatToBinaryString(num float64) string {
//
//	if num == 0 {
//		return "0"
//	}
//
//	//проверка на отрицательное число, если отрицательное, то будет положительным,
//	//и в конце уже припишем символом минус.
//	var flag bool = false
//	if num < 0 {
//		flag = true
//		num = -num
//	}
//
//	// Разделяем число на целую и дробную части
//	intPart := int64(num)
//	fracPart := num - float64(intPart)
//
//	// Преобразуем целую часть
//	intBinary := ""
//	for intPart > 0 {
//		intBinary = fmt.Sprintf("%d", intPart%2) + intBinary
//		intPart /= 2
//	}
//
//	// Преобразуем дробную часть
//	fracBinary := ""
//	for fracPart > 0 {
//		fracPart *= 2
//		bit := int64(fracPart)
//		if bit == 1 {
//			fracBinary += "1"
//			fracPart -= float64(bit)
//		} else {
//			fracBinary += "0"
//		}
//		// Ограничиваем длину дробной части для избежания бесконечных циклов
//		if len(fracBinary) > 30 { // например, 10 битов после запятой
//			break
//		}
//	}
//
//	if intBinary == "" {
//		intBinary = "0"
//	}
//
//	//проверка на случай если и число отрицательное и нету дробной части у числа
//	if fracBinary == "" && flag == true {
//		return "-" + intBinary
//	}
//
//	//если число было отрицательным, то вот как раз приписываем минус.
//	if flag == true {
//		return "-" + intBinary + "." + fracBinary
//	}
//
//	// Собираем результат
//	if fracBinary == "" {
//		return intBinary
//	}
//
//	return intBinary + "." + fracBinary
//}

//package main
//
//import "fmt"
//
//func main() {
//	fmt.Println("Task number A6: ")
//	var n int
//	fmt.Println("Enter number n: ")
//	fmt.Scan(&n)
//
//	result := Sum(n)
//	fmt.Println("Result of addition in sequence n: ", result)
//}
//
//func Sum(n int) float64 {
//	var summ, an float64
//
//	for i := 0; i < n; i++ {
//		an = 1.0 / float64(i*i*i+i+5)
//		//golang может вывести изначально сразу около 23 тысяч строк в консоль и потом чистит, но если число очень большое и много
//		//нужно выводить, то он сокращает число выводимых строк на консоль самостоятельно с каждой итерацией, что ускоряет процесс вычисления.
//		//но немного урезает вывод в консоль до 22 или 21 тысяч строки(от 10 миллионов выводов в консоль).
//		fmt.Println("Result iter(", i, "): ", an)
//		summ += an
//	}
//
//	return summ
//}

package main

import (
	"fmt"
	"os"
	"strconv"
	"text/tabwriter"
)

// вывод матрицы на экран
func matrix_output(matrix *[][]int, size int) {
	for ii := 0; ii < size; ii++ {
		w := tabwriter.NewWriter(os.Stdout, 4, 0, 0, ' ', tabwriter.AlignRight)
		for jj := 0; jj < size; jj++ {
			str := strconv.Itoa((*matrix)[ii][jj])
			fmt.Fprintf(w, "%s\t\t", str)
		}
		fmt.Println()
		w.Flush()
	}
}

// заполнение матрицы с консоли
func matrix_input(matrix *[][]int, size int) {
	fmt.Println("Начинайте вводить значения матрицы: ")
	for ii := 0; ii < size; ii++ {
		for jj := 0; jj < size; jj++ {
			fmt.Scanf("%d", &(*matrix)[ii][jj])
		}
		fmt.Println()
	}
}

// вычисление самого определителя
func determinant_calculation(matrix *[][]int, size int) int {
	// Описываем Базовый случай для матрицы 1x1
	if size == 1 {
		return (*matrix)[0][0]
	}

	// Так же описываем Базовый случай для матрицы 2x2
	if size == 2 {
		tmp1 := (*matrix)[0][0] * (*matrix)[1][1]
		tmp2 := (*matrix)[0][1] * (*matrix)[1][0]
		return (tmp1 - tmp2)
	}

	// Рекурсивный случай для матриц размером больше 2x2
	res := 0
	for col := 0; col < size; col++ { //выбор элементов, которые надо исключать будем делать по колонкам(и соответственно исключать верхнюю строчку
		//будут менять лишь её колонки, и так для каждой меньшей матрице, которую мы будем получать в ходе рекурсии)

		// Создаем минор для текущего элемента
		minor := getMinor(matrix, 0, col, size)
		// Вычисляем алгебраическое дополнение и добавляем его к результату

		number_and_minor := (*matrix)[0][col] * determinant_calculation(&minor, size-1) //уменьшаем на 1 и колонки и столбцы
		//где (*matrix)[0][col] - это элемент, который исключили, а это минор к этому элементу:
		//determinant_calculation(&subMatrix, rows-1, cols-1), а number_and_minor - это перемноженное исключенное число на
		//соотвествующий этому числу минор

		if col%2 != 0 { // в конце добавляем знак если позиция столбца не четная то ставим знак минус,
			// прямо как в формуле по вычислению миноров
			number_and_minor = -number_and_minor
		}

		//Выведем на экран наши действия:
		fmt.Printf("%d * %d = %d\n", (*matrix)[0][col], minor, number_and_minor)

		//И по итогу цикл повторяется до тех пор, пока не попадется на случай 2 на 2, который описан выше.
		//Так и работает в нашем случае рекурсия, постоянно уменьшая матрицу и приводя её к случаю 2 на 2.
		res += number_and_minor
	}

	return res
}

// Создание минора для заданной строки и столбца
func getMinor(matrix *[][]int, rowToRemove int, colToRemove int, size int) [][]int {
	//где rowToRemove - это индекс исключаемой строки, а colToRemove - это индекс исключаемого столбца,
	//а size - размер квадратной матрицы, тут уже полностью решил не расписывать.
	minor := [][]int{}

	for ii := 0; ii < size; ii++ {
		if ii == rowToRemove { //у нас как раз всегда будет браться нулевая строка, так как в rowToRemove, всегда передается 0.
			continue // а это случай мы как раз таки пропускаем.
		}

		row := []int{} //создаем массив столбцов

		for jj := 0; jj < size; jj++ {
			if jj == colToRemove { //а тут у нас элемент исключаемого столбца всегда будет меняться и с каждым
				// разом становиться на 1 больше чем до этого
				continue //поэтому пропускаем этот случай
			}
			row = append(row, (*matrix)[ii][jj]) //начинаем заполнять новый массив, элементами из нашего минора(то есть всеми,
			//кроме тех, которые мы сами убрали)
		}

		minor = append(minor, row) //а теперь в наш массив массивов, или же в матрицу, записываем такой небольшой массив,
		// в котором нету лишнего исключенного столбца, и так с каждой строкой, в каждой строке не будет элемента
		//определённого столбца
	}

	//И так наш алгоритм повторяется всю длину - size( при этом при передачи в эту функцию она на 1 с каждым рекурсивным выполнением
	//функции - determinant_calculation, будет сокращаться на 1, так как там прописано size - 1, в функции: determinant_calculation,
	//а по скольку это квадратная матрица, то сократился как раз таки 1 столбец и сократилась 1 строка)
	return minor
}

func main() {
	fmt.Println("Задание на вычисление определителя любой квадратной матрицы через метод минора.")
	fmt.Println("Введите квадратную матрицу: ")

	//Инициализация переменных отвечающих за размеры матрицы
	var size int

	fmt.Println("Введите размеры строк на стобцов одной цифрой(если 3 на 3, то вводите 3 и тд): ")
	fmt.Scanf("%d", &size)

	var matrix [][]int //инициализация матрицы

	//создаем строки матрицы
	for ii := 0; ii < size; ii++ {
		//создаем столбцы матрицы, или же массивы внутри массива
		row := make([]int, size)
		matrix = append(matrix, row)
	}

	matrix_input(&matrix, size)

	fmt.Println("Ваша исходная матрица: ")
	matrix_output(&matrix, size)
	fmt.Println("\n\nНахождение определителя через миноры: ") //делаем переход на след строку, чтобы отделить текст

	result := determinant_calculation(&matrix, size)

	fmt.Println("\nОпределитель вашей матрицы: ", result)
}
