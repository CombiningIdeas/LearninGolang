package tasks

import "fmt"

//Variant7
//Вариант 7
//Дано:
//А – массив натуральных чисел, в котором нет одинаковых элементов;
//В – массив натуральных чисел, в котором нет одинаковых элементов.
//Определить, верно ли, что в массивах А и В нет общих элементов.

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
		fmt.Println("Да. В массивах А и В нет общих элементов.")
	} else {
		fmt.Println("Нет. В массивах А и В есть общие элементы.")
	}

}

func workingWithArrays(lenArrayA uint32, lenArrayB uint32, arrayA []uint32,
	arrayB []uint32) bool {

	//для того, чтобы проверить есть ли общие элемента в массивах
	//нам понадобятся вложенные циклы и флаги

	var resultFlag bool = true //true - означает что общих элементов нету

	for ii := 0; ii < int(lenArrayA); ii++ {
		tmpFlag := true
		for jj := 0; jj < int(lenArrayB); jj++ {
			if arrayA[ii] == arrayB[jj] {
				tmpFlag = false //это означает что у нас есть общие элементы.
				break
			}
		}

		if tmpFlag == false { //если есть хотя бы один общий элемент, то это
			// значит, что в массивах уже как минимум есть 1 общий элемент
			resultFlag = false //общий элемент найден
			break
		}
	}

	return resultFlag

}
