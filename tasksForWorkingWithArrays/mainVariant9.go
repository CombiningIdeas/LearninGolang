package tasks

import "fmt"

//Variant9
//Вариант 9
//Даны массивы натуральных чисел А и В, упорядоченные по возрастанию.
//Получить упорядоченный по возрастанию массив С, содержащий все такие
//элементы, которые есть и в массиве А и в массиве В.

func main() {

	var lenArrayA, lenArrayB, lenArrayC uint32

	fmt.Println("Введите длину массива <A>: ")
	fmt.Scan(&lenArrayA)
	fmt.Println("Введите длину массива <B>: ")
	fmt.Scan(&lenArrayB)

	lenArrayC = lenArrayA + lenArrayB

	var arrayA []uint32 = make([]uint32, lenArrayA)
	var arrayB []uint32 = make([]uint32, lenArrayB)
	var arrayC []uint32 = make([]uint32, lenArrayC)

	var counter int = 1 //это счетчик отвечающий за изображение номера элемента в массиве
	fmt.Println("Введите массив <A>, упорядоченный по возрастанию: ")
	for iter := 0; iter < int(lenArrayA); {
		fmt.Printf("%d элемент: \n", iter+counter)
		fmt.Scan(&arrayA[iter])

		if iter > 0 && arrayA[iter] == arrayA[iter-1] {
			arrayA = append(arrayA[:iter], arrayA[iter+1:]...) // Удаляем текущий дубликат
			lenArrayA--                                        // Уменьшаем длину массива
			counter++
		} else {
			iter++ // Только если элемент не дубликат, двигаемся дальше
		}
	}

	counter = 1 //тут используем тот же счетчик
	fmt.Println("\nВведите массив <B>, упорядоченный по возрастанию: ")
	for iter := 0; iter < int(lenArrayB); {
		fmt.Printf("%d элемент: \n", iter+counter)
		fmt.Scan(&arrayB[iter])

		if iter > 0 && arrayB[iter] == arrayB[iter-1] {
			arrayB = append(arrayB[:iter], arrayB[iter+1:]...) // Удаляем текущий дубликат
			lenArrayB--                                        // Уменьшаем длину массива
			counter++
		} else {
			iter++ // Только если элемент не дубликат, двигаемся дальше
		}
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

	ii, jj := 0, 0 //это счетчики, которые нужны для прохода по массивам
	*lenArrayC = 0

	for ii < int(lenArrayA) && jj < int(lenArrayB) {
		if arrayA[ii] < arrayB[jj] { //если элемент массива В больше
			ii++
		} else if arrayA[ii] > arrayB[jj] { //если элемент массива А больше
			jj++
		} else { //если элемент массива А равен элементу массива В

			//Добавляем элемент в C только если он не дублируется
			//и если элементов еще не было в массиве С:
			if len(*arrayC) == 0 || (*arrayC)[int(*lenArrayC)] != arrayA[ii] {
				(*arrayC)[(*lenArrayC)] = arrayA[ii]
				(*lenArrayC)++
			}
			ii++
			jj++
		}
	}

}
