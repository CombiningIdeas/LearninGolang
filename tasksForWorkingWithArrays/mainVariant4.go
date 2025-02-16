package tasks

import "fmt"

//Variant4
//Вариант 4
//Дано:
//А – массив натуральных чисел, в котором нет одинаковых элементов;
//В – массив натуральных чисел, в котором нет одинаковых элементов.
//Получить массив С, содержащий все элементы массива А, которых нет в В и все
//элементы массива В, которых нет в А.

func main() {

	var lenArrayA, lenArrayB, lenArrayC uint32

	fmt.Println("Введите длину массива <A>: ")
	fmt.Scan(&lenArrayA)
	fmt.Println("Введите длину массива <B>: ")
	fmt.Scan(&lenArrayB)

	lenArrayC = lenArrayA + lenArrayB //длина нашего массива в этой задачке не может быть
	// больше длины массива А.

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

	//после прошлого варианта я понял, что сделал не совсем оптимизировано вариант 3, лучше всего
	//стоило через флаг булевый true и false сделать, а не счетчиком, но зато сейчас так сделаю.

	var flagForA, flagForB bool
	*lenArrayC = 0 //обнуляем длину массива С

	//сначала проверяем массив А, на сколько там есть элементов, которых нет в массиве В
	for ii := 0; ii < int(lenArrayA); ii++ {
		flagForA = true

		for jj := 0; jj < int(lenArrayB); jj++ {
			if arrayA[ii] == arrayB[jj] {
				flagForA = false //это означает что у нас есть совпадения.
				break
			}
		}

		if flagForA == true { //а это значит что повторений какого то числа из массива А,
			//в массиве В не было.
			(*arrayC)[int(*lenArrayC)] = arrayA[ii]
			(*lenArrayC)++ //увеличиваем длину массива С
		}

	}

	//А теперь проверяем массив В, сколько тут элементов, которых нет в массиве А
	for ii := 0; ii < int(lenArrayB); ii++ {
		flagForB = true

		for jj := 0; jj < int(lenArrayA); jj++ {
			if arrayB[ii] == arrayA[jj] {
				flagForB = false //это означает что у нас есть совпадения.
				break
			}
		}

		if flagForB == true { //а это значит что повторений какого то числа из массива А,
			// в массиве В не было.
			(*arrayC)[int(*lenArrayC)] = arrayB[ii]
			(*lenArrayC)++ //увеличиваем длину массива С
		}

	}

}
