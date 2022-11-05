package main

/* Удалить i-ый элемент из слайса. */

import "fmt"

func main() {
	s := []int{1, 2, 3, 4, 5}
	i := 2

	if len(s) != 0 && i < len(s)-1 { // защищаемся от паники
		s = append(s[:i], s[i+1:]...) //соединяем две части слайса, исключая i-ый элемент
	}
	fmt.Println(s)
}
