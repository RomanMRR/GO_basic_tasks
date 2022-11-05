package main

/* Разработать программу, которая перемножает, делит, складывает,
вычитает две числовых переменных a,b, значение которых > 2^20. */

import (
	"fmt"
	"math/big"
)

func main() {

	a := new(big.Int)
	b := new(big.Int)
	sum := new(big.Int)
	mul := new(big.Int)
	div := new(big.Int)
	diff := new(big.Int)

	var a_input, b_input string

	fmt.Println("Введите два больших числа:")
	fmt.Scanf("%s\n", &a_input)
	fmt.Scanf("%s\n", &b_input)

	a.SetString(a_input, 10)
	b.SetString(b_input, 10)

	fmt.Println("Результат сложения:", sum.Add(a, b))
	fmt.Println("Результат умножения:", mul.Mul(a, b))
	fmt.Println("Результат вычитания:", div.Sub(a, b))
	fmt.Println("Результат деления:", diff.Div(a, b))

}
