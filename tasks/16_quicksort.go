package main

/* Реализовать быструю сортировку массива (quicksort) встроенными методами языка.
 */

import "fmt"

func quicksort(array []int, left int, right int) []int {
	//Создаем указатели на левую и правую части массива
	l := left
	r := right

	//Берём значение центральной ячейки, относительно которой будем сортировать
	pivot := array[(left+right)/2]

	//Цикл, начинающий саму сортировку
	for l < r {

		//Ищем значения больше 'центра'
		for array[r] > pivot {
			r--
		}

		//Ищем значения меньше 'центра'
		for array[l] < pivot {
			l++
		}

		//Когда найдём первые числа, не удовлетворяющие верхним условиям
		if l <= r {
			//Мы меняем их местами
			array[r], array[l] = array[l], array[r]
			l++
			r--
		}
	}

	//Пройдя весь массив мы сортируем его левую часть
	if r > left {
		quicksort(array, left, r)
	}

	//И правую часть
	if l < right {
		quicksort(array, l, right)
	}

	return array
}

func main() {
	array := [10]int{5, 6, 3, 2, 7, 8, 9, 12, 13, 10}
	quicksort(array[:], 0, len(array)-1)
	fmt.Println(array)
}
