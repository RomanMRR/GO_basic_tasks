package main

/*
К каким негативным последствиям может привести данный фрагмент кода,
и как это исправить? Приведите корректный пример реализации.


var justString string
func someFunc() {
  v := createHugeString(1 << 10)
  justString = v[:100]
}

func main() {
  someFunc()
}


*/

import (
	"fmt"
	"strings"
)

func countSymbols(str *string) int {
	count := 0
	for range *str {
		count += 1
	}

	return count
}

func createHugeString(n int) string {
	var str string

	for i := 0; i < n; i++ {
		str += "F"
	}

	return str
}

var justString string

func someFunc() {
	//Создаём большую строку, длинной 2^10 степени, так как каждый сдвиг влево это умножение на 2
	v := createHugeString(1 << 10)
	//Получаем срез получившейся строки
	// justString = v[:100] //Если мы будем просто брать срез, то строка
	//будет указывать на большую строку
	//И если изменить justString (в unsafe к примеру)
	//то изменится и большая строка
	justString = strings.Clone(v[:100]) //Поэтому копируем строку в новую область памяти

	fmt.Println(justString)
}

func main() {
	someFunc()
}
