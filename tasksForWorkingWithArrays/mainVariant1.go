package tasks

//Variant1
//Вариант 1
//Дано:
//А – массив натуральных чисел, в котором нет одинаковых элементов;
//В – массив натуральных чисел, в котором нет одинаковых элементов.
//Получить массив С, содержащий все элементы массивов А и В без повторений.

import "fmt"

func main() {

	//Узнаем длину массивов и создаем эти массивы.
	var lenArrayA, lenArrayB, lenArrayC uint32

	fmt.Println("Введите длину массива <A>: ")
	fmt.Scan(&lenArrayA)
	fmt.Println("Введите длину массива <B>: ")
	fmt.Scan(&lenArrayB)

	lenArrayC = lenArrayA + lenArrayB

	var arrayA []uint32 = make([]uint32, lenArrayA)
	var arrayB []uint32 = make([]uint32, lenArrayB)
	var arrayC []uint32 = make([]uint32, lenArrayC)

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

	workingWithArrays(lenArrayA, lenArrayB, &lenArrayC, arrayA, arrayB, &arrayC)

	//fmt.Println(lenArrayC) //тестирование длинны массива С
	fmt.Println("\n\nМассив <С>: ")
	for iter := 0; iter < int(lenArrayC); iter++ {
		fmt.Printf("%d элемент: %d \n", iter+1, arrayC[iter])
	}
}

func workingWithArrays(lenArrayA uint32, lenArrayB uint32, lenArrayC *uint32, arrayA []uint32,
	arrayB []uint32, arrayC *[]uint32) {

	var flag bool
	var counter int = 0

	for iter := 0; iter < int(lenArrayA); iter++ {
		(*arrayC)[iter] = arrayA[iter]
	}

	for iter := 0; iter < int(lenArrayB); iter++ {
		flag = true

		for ii := 0; ii < int(lenArrayA); ii++ {
			if ((*arrayC)[ii]) == arrayB[iter] {
				flag = false
			}
		}

		if flag == true {
			(*arrayC)[int(lenArrayA)+counter] = arrayB[iter]
		} else {
			counter--
			*lenArrayC--
		}

		counter++
	}
}
