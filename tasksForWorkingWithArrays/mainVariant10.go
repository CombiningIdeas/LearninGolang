package tasks

import "fmt"

//Variant10
//Вариант 10
//Даны массивы натуральных чисел А и В, упорядоченные по возрастанию.
//Получить упорядоченный по возрастанию массив С, содержащий все элементы
//массива А, которых нет в В.

func main() {

	var lenArrayA, lenArrayB uint32

	fmt.Println("Введите длину массива <A>: ")
	fmt.Scan(&lenArrayA)
	fmt.Println("Введите длину массива <B>: ")
	fmt.Scan(&lenArrayB)

	var arrayA []uint32 = make([]uint32, lenArrayA)
	var arrayB []uint32 = make([]uint32, lenArrayB)

	//еще мы можем не задавать длину массиву в go, тогда потом придется добавлять элементы
	//через append в конец массива, длину массива можно узнать как в пайтоне через метод len()
	//так что можно по сути писать в любом стиле, можно и считывание элементов сделать через apend()
	//но я думаю что лучше оставить так в стиле С, так при считывании меньше кода надо будет писать.
	var arrayC []uint32 //

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

	workingWithArrays(arrayA, arrayB, &arrayC)

	//fmt.Println(lenArrayC) //тестирование длинны массива С
	fmt.Println("\n\nМассив <С>: ")
	for iter := 0; iter < len(arrayC); iter++ {
		fmt.Printf("%d элемент: %d \n", iter+1, arrayC[iter])
	}

}

func workingWithArrays(arrayA []uint32, arrayB []uint32, arrayC *[]uint32) {

	ii, jj := 0, 0 //это счетчики, которые нужны для прохода по массивам ii - для
	//массива arrayA, а jj - для массива arrayB.

	for ii < len(arrayA) {
		if jj < len(arrayB) && arrayA[ii] == arrayB[jj] {
			//Пропускаем этот элемент, увеличивая оба индекса (ii и jj),
			//поскольку нам нужны только те элементы arrayA, которых нет в arrayB.
			ii++
			jj++
		} else if jj < len(arrayB) && arrayA[ii] > arrayB[jj] {
			//если элемент массива arrayA стал больше текущего элемента массива B,
			//то мы не записываем текущий элемент массива А, поскольку массив В при увеличении
			//индекса может дать элемент равный текущему в массиве А
			jj++
		} else {
			//а вот тут мы уже записываем тут или кончился массив В, а значит массив А с
			//ним уже никак не совпадет, или элемент массива В, больше чем текущий элемент массива А,
			//и все прошлый элементы массива В, не были равны этому элементу массива А. И тут как не
			//увеличивай или не уменьшай массив В, ни один его элемент уже не станет равным текущему в массиве А.
			*arrayC = append(*arrayC, arrayA[ii])
			ii++
		}
	}

}
