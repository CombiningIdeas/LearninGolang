package tasks

import "fmt"

//Variant8
//Вариант 8
//Даны массивы натуральных чисел А и В, упорядоченные по возрастанию.
//Получить упорядоченный по возрастанию массив С, содержащий все элементы
//массивов А и В.

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

	fmt.Println("Введите массива <A>, упорядоченный по возрастанию: ")
	for iter := 0; iter < int(lenArrayA); iter++ {
		fmt.Printf("%d элемент: \n", iter+1)
		fmt.Scan(&arrayA[iter])
	}

	fmt.Println("\nВведите массива <B>, упорядоченный по возрастанию:: ")
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

	//переменная ii отвечает за передвижение в массиве А
	//переменная jj отвечает за передвижение в массиве B
	//переменная kk отвечает за передвижение в массиве C
	ii, jj, kk := 0, 0, 0

	//в это цикле записываем в массив С, из обоих массивов, пока в обоих
	//одновременно есть еще какие-то элементы
	//так как элементы упорядоченны то можно сравнивать в лоб
	for ii < len(arrayA) && jj < len(arrayB) {
		if arrayA[ii] < arrayB[jj] {
			(*arrayC)[kk] = arrayA[ii]
			ii++ //увеличиваем счетчик передвижения массива А
		} else {
			(*arrayC)[kk] = arrayB[jj]
			jj++ //увеличиваем счетчик передвижения массива В
		}

		//эта переменная увеличивается независимо от результата сравнений
		kk++ //увеличиваем счетчик передвижения массива С
	}

	//так же у нас может быть такая ситуация когда или массивы разных размеров или
	//просто после сравнения остались переменные только одного массива, но нам
	//нужно обработать и их

	// Если остались элементы в массиве A:
	for ii < len(arrayA) {
		(*arrayC)[kk] = arrayA[ii]
		ii++
		kk++
	}

	// Если остались элементы в массиве B:
	for jj < len(arrayB) {
		(*arrayC)[kk] = arrayB[jj]
		jj++
		kk++
	}

}
