package main

/* Реализовать бинарный поиск встроенными методами языка. */

import "fmt"

func leftBinSearch(array []int, left int, right int, need int) interface{} { //ЗДесь указываем пустой интерфейс, так как может быть возвращено или число или указатель
	for left < right { //Идём слева направо
		m := (right + left) / 2 //Получаем индекс серидины той части, что ещё не рассматривали
		if need < array[m] {    //Если нужный элемент меньше,
			right = m //То надо идти в левую часть массива
		} else if need > array[m] {
			left = m + 1 //Иначе в правую часть
		} else {
			return m //Если нашли нужный элемент возвращаем nil
		}
	}
	return nil //Если не нашли, то возвращаем nil
}

func main() {
	array := [13]int{2, 3, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	need := 1000
	fmt.Println(leftBinSearch(array[:], 0, len(array)-1, need))
}
