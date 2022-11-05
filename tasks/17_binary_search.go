package main

/* Реализовать бинарный поиск встроенными методами языка. */

import "fmt"

func leftBinSearch(array []int, left int, right int, need int) int {
	for left < right {
		m := (right + left) / 2
		if need < array[m] {
			right = m
		} else if need > array[m] {
			left = m + 1
		} else {
			return m
		}
	}
	return left
}

func main() {
	array := [13]int{2, 3, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	fmt.Println(leftBinSearch(array[:], 0, len(array)-1, 9))
}
