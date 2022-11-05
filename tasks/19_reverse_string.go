package main

/* Разработать программу, которая переворачивает подаваемую на ход строку (например: «главрыба — абырвалг»). Символы могут быть unicode.
 */

import (
	"fmt"
	"strings"
)

// Функция для переворачивания строки
func ReverseString(str []rune) string {
	var sb strings.Builder //Реализует эффективный способ конкантенации

	right := len(str) - 1 //Будем идти с конца строки

	for right >= 0 { //ПОка не дошли до конца
		sb.WriteString(string(str[right])) //Получаем символы строки с конца
		right--
	}

	return sb.String()
}

func main() {
	var str string

	fmt.Scanln(&str)

	fmt.Println(ReverseString([]rune(str)))
}
