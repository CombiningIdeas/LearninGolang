package tasks

import "fmt"

//Variant6
//Вариант 6
//Дано:
//А – массив натуральных чисел, в котором нет одинаковых элементов;
//В – массив натуральных чисел, в котором нет одинаковых элементов.
//Определить, верно ли, что массивы А и В состоят из одинаковых элементов.

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
		fmt.Println("Да. Массивы А и В состоят из одинаковых элементов")
	} else {
		fmt.Println("Нет. Массивы А и В не состоят из одинаковых элементов")
	}

}

func workingWithArrays(lenArrayA uint32, lenArrayB uint32, arrayA []uint32,
	arrayB []uint32) bool {

	var resultFlag bool = true

	for ii := 0; ii < int(lenArrayA); ii++ {
		tmpFlag := false
		for jj := 0; jj < int(lenArrayB); jj++ {
			if arrayA[ii] == arrayB[jj] {
				tmpFlag = true //это означает что у нас есть совпадения.
				break
			}
		}

		if tmpFlag == false {
			resultFlag = false
			break
		}
	}

	//если у нас уже хотя бы один элемент массива А, не соответствует ни одному элементу массива В,
	//то можно сразу завершать функцию, а если нет то проверим еще и массив В
	if resultFlag == false {
		return resultFlag
	}

	for ii := 0; ii < int(lenArrayB); ii++ {
		tmpFlag := false
		for jj := 0; jj < int(lenArrayA); jj++ {
			if arrayB[ii] == arrayA[jj] {
				tmpFlag = true //это означает что у нас есть совпадения.
				break
			}
		}

		if tmpFlag == false {
			resultFlag = false
			break
		}
	}

	return resultFlag

}
