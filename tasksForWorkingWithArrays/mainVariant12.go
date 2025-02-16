package tasks

import "fmt"

//Variant12
//Вариант 12
//Даны массивы натуральных чисел А и В, упорядоченные по возрастанию.
//Определить, верно ли, что массив В содержит каждый элемент массива А.

func main() {

	var lenArrayA, lenArrayB uint32

	fmt.Println("Введите длину массива <A>: ")
	fmt.Scan(&lenArrayA)
	fmt.Println("Введите длину массива <B>: ")
	fmt.Scan(&lenArrayB)

	var arrayA []uint32 = make([]uint32, lenArrayA)
	var arrayB []uint32 = make([]uint32, lenArrayB)

	var counter int = 1
	fmt.Println("Введите массив <A>, упорядоченный по возрастанию: ")
	for iter := 0; iter < int(lenArrayA); {
		fmt.Printf("%d элемент: \n", iter+counter)
		fmt.Scan(&arrayA[iter])

		if iter > 0 && arrayA[iter] == arrayA[iter-1] {
			arrayA = append(arrayA[:iter], arrayA[iter+1:]...)
			lenArrayA--
			counter++
		} else {
			iter++
		}
	}

	counter = 1
	fmt.Println("\nВведите массив <B>, упорядоченный по возрастанию: ")
	for iter := 0; iter < int(lenArrayB); {
		fmt.Printf("%d элемент: \n", iter+counter)
		fmt.Scan(&arrayB[iter])

		if iter > 0 && arrayB[iter] == arrayB[iter-1] {
			arrayB = append(arrayB[:iter], arrayB[iter+1:]...)
			lenArrayB--
			counter++
		} else {
			iter++
		}
	}

	result := workingWithArrays(arrayA, arrayB)

	if result {
		fmt.Println("Верно. Массив В содержит каждый элемент массива А.")
	} else {
		fmt.Println("Не верно. Массив В не содержит каждый элемент массива А.")
	}

}

func workingWithArrays(arrayA []uint32, arrayB []uint32) bool {

	ii, jj := 0, 0
	containsAll := true

	for ii < len(arrayA) && jj < len(arrayB) {
		if arrayA[ii] > arrayB[jj] {
			jj++
		} else if arrayA[ii] == arrayB[jj] {
			ii++
			jj++
		} else {
			containsAll = false
			break
		}
	}

	if ii < len(arrayA) { //это означает что счетчик по массиву А, не закончился,
		// а значит есть уникальные элементы в массиве А
		containsAll = false
	}

	return containsAll
}
