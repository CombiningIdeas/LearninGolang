package tasks

import "fmt"

//Variant3
//Вариант 3
//Дано:
//А – массив натуральных чисел, в котором нет одинаковых элементов;
//В – массив натуральных чисел, в котором нет одинаковых элементов.
//Получить массив С, содержащий все элементы массива А, которых нет в В.

func main() {

	var lenArrayA, lenArrayB, lenArrayC uint32

	fmt.Println("Введите длину массива <A>: ")
	fmt.Scan(&lenArrayA)
	fmt.Println("Введите длину массива <B>: ")
	fmt.Scan(&lenArrayB)

	lenArrayC = lenArrayA //длина нашего массива в этой задачке не может быть
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

	var lengthCounter int
	*lenArrayC = 0 //обнуляем длину массива С

	//тут нужен алгоритм который проходил бы по 2 циклам и по итогу сравнивал каждый
	//элемент массива A с каждым элементом массива B, если ни одно значение из массива B,
	//не было равно какому-либо значению из массива А, то тогда это самое значение
	//из массива А, запишем в массив С. Для этого будем использовать счетчик, который будем
	//увеличивать в случае если не совпадают и обнулять если совпадают и если будут не совпадать
	//все значения в массиве B, то тогда счетчик будет равен длине массива B, и вот тогда то
	//и запишем это значение массива А, в массив C, и увеличим длину самого массива С на 1.
	for ii := 0; ii < int(lenArrayA); ii++ {
		lengthCounter = 0 //на каждом новом элемента массива А, мы обнуляем счетчик.

		for jj := 0; jj < int(lenArrayB); jj++ {
			if arrayA[ii] != arrayB[jj] {
				lengthCounter++ //длина счетчика не совпадений одного значения массива В,
				// со всеми значениями массива А
			} else {
				lengthCounter = 0 // если же все же значение есть в массиве B, то сразу обнуляем
				// счетчик и переходим на следующее значение массива А
				jj = int(lenArrayB) //записываем максимальное значение чтобы бессмысленно
				// не выполнять оставшиеся итерации
			}

			if lengthCounter == int(lenArrayB) {
				(*arrayC)[int(*lenArrayC)] = arrayA[ii] //записываем результат в наш динамический массив С
				(*lenArrayC)++                          //а теперь увеличиваем новую длину массива С
			}
		}
	}

}
