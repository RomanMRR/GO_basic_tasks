package main

/* Поменять местами два числа без создания временной переменной. */

import (
	"fmt"
)

func main() {
	var a, b int

	a = 5
	b = 6

	fmt.Printf("Было a=%d, b=%d\n", a, b)
	a, b = b, a
	fmt.Printf("Стало a=%d, b=%d\n", a, b)

}
