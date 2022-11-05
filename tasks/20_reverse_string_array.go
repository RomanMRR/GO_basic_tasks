/*
Разработать программу, которая переворачивает слова в строке.
Пример: «snow dog sun — sun dog snow».
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Функция для переворачивания слов в строке
func ReverseStringArray(str string) string {

	arratString := strings.Split(str, " ")
	l := 0
	r := len(arratString) - 1
	for l < r { //Идём одновременно и с конца и с начала
		arratString[l], arratString[r] = arratString[r], arratString[l] //Меняем элементы
		l++
		r--
	}

	return strings.Join(arratString, " ")
}

func main() {
	var str string
	reader := bufio.NewReader(os.Stdin) //Для чтения с консоли

	str, _ = reader.ReadString('\n')        //Читаем пока не нажали на Enter
	str = strings.Replace(str, "\n", "", 1) //Заменяем символ новой строки на пустоту, так как этот символ на не нужен
	result_string := ReverseStringArray(str)
	fmt.Println(result_string)
}
