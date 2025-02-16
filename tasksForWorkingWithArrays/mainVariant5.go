package tasks

import "fmt"

//Variant5
//Вариант 5
//Дано:
//А – массив натуральных чисел, в котором нет одинаковых элементов;
//В – массив натуральных чисел, в котором нет одинаковых элементов.
//Определить, верно ли, что массив В содержит каждый элемент массива А.

func main() {

	var lenArrayA, lenArrayB uint32

	fmt.Println("Введите длину массива <A>: ")
	fmt.Scan(&lenArrayA)
	fmt.Println("Введите длину массива <B>: ")
	fmt.Scan(&lenArrayB)

	var arrayA []uint32 = make([]uint32, lenArrayA)
	var arrayB []uint32 = make([]uint32, lenArrayB)

	fmt.Println("Введите массива <A>: ")
	for iter := 0; iter < int(lenArrayA); iter++ {
		fmt.Printf("%d элемент: \n", iter+1)
		fmt.Scan(&arrayA[iter])
	}

	fmt.Println("\nВведите массива <B>: ")
	for iter := 0; iter < int(lenArrayB); iter++ {
		fmt.Printf("%d элемент: \n", iter+1)
		fmt.Scan(&arrayB[iter])
	}

	var result bool = workingWithArrays(lenArrayA, lenArrayB, arrayA, arrayB)

	if result {
		fmt.Println("Да. Массив В содержит каждый элемент массива А.")
	} else {
		fmt.Println("Нет. Массив В не содержит каждый элемент массива А.")
	}

}

func workingWithArrays(lenArrayA uint32, lenArrayB uint32, arrayA []uint32,
	arrayB []uint32) bool {

	//по сути нам нужно проверить только 1 массив, это массив В.
	var resultFlag bool = true

	for ii := 0; ii < int(lenArrayA); ii++ {
		tmpFlag := false
		for jj := 0; jj < int(lenArrayB); jj++ {
			if arrayA[ii] == arrayB[jj] {
				tmpFlag = true //это означает что у нас есть совпадения.
				break
			}
		}

		if tmpFlag == false { //а тут у нас проверка на то, что уже хотя
			// бы один элемент массива А, не соответствует ни одному элементу из массива В
			resultFlag = false
			break
		}
	}

	return resultFlag

}
