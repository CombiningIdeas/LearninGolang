package tasks

import "fmt"

//Variant11
//Вариант 11
//Даны массивы натуральных чисел А и В, упорядоченные по возрастанию.
//Получить упорядоченный по возрастанию массив С, содержащий все элементы
//массива А, которых нет в В и все элементы массива В, которых нет в А.

func main() {

	var lenArrayA, lenArrayB uint32

	fmt.Println("Введите длину массива <A>: ")
	fmt.Scan(&lenArrayA)
	fmt.Println("Введите длину массива <B>: ")
	fmt.Scan(&lenArrayB)

	var arrayA []uint32 = make([]uint32, lenArrayA)
	var arrayB []uint32 = make([]uint32, lenArrayB)
	var arrayC []uint32

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

	//поскольку нам нужны не повторяющиеся элементы обоих массивов, то мы можем продолжать
	//итерации по циклу до тех пор, пока в хотя бы одном из массивов есть элементы:
	for ii < len(arrayA) || jj < len(arrayB) {
		if ii < len(arrayA) && (jj >= len(arrayB) || arrayA[ii] < arrayB[jj]) {
			//это по сути случай когда счетчик масива В стал равен(или больше) длинны массива В, но другой еще нет
			//то тогда да элементы в массиве А, будут уникальными, или же когда элементы массива А еще не закончились,
			//но текущий элемент массива А, меньше текущего элемента массива В, тогда запишем элемент из массива А, и
			//увеличим его счетчик, не увеличивая счетчик массива В, так как следующие элементы могут быть равны к примеру.
			*arrayC = append(*arrayC, arrayA[ii])
			ii++
		} else if jj < len(arrayB) && (ii >= len(arrayA) || arrayB[jj] < arrayA[ii]) {
			//тут тоже самое как и в первом if, только наоборот.
			*arrayC = append(*arrayC, arrayB[jj])
			jj++
		} else { //в ином случае, если элементы совпадают, то мы просто увеличиваем счетчик
			ii++
			jj++
		}
	}
}
