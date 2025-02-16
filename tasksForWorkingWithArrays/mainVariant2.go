package tasks

import "fmt"

//Variant2
//Вариант 2
//Дано:
//А – массив натуральных чисел, в котором нет одинаковых элементов;
//В – массив натуральных чисел, в котором нет одинаковых элементов.
//Получить массив С, содержащий все такие элементы, которые есть и в массиве А и в
//массиве В.

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

	var repetitionCounter int = 0

	//здесь важно знать длину нового динамического массива
	//по сути нужно найти повторяющиеся элементы, и мы знаем что в каждом массиве
	//они уникальны(по условию) то писать проверку на то когда в каждом есть условно
	//повторяющиеся и в другом точно такие же и при проверки счетчик будет много раз
	//увеличен хотя должен быть всего 1, но опять таки у нас этой проблемы нету, так
	//что все хорошо, дополнительная проверка не потребуется
	for ii := 0; ii < int(lenArrayA); ii++ {
		for jj := 0; jj < int(lenArrayB); jj++ {
			if arrayA[ii] == arrayB[jj] {
				(*arrayC)[repetitionCounter] = arrayA[ii] //сразу записываем результат в наш динамический массив С
				repetitionCounter++                       //длина будущего массива С.
			}
		}
	}

	*lenArrayC = uint32(repetitionCounter) //делаем разыменование и приведение типов

}
