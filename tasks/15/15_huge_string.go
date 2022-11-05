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
		str += "狐"
	}

	return str
}

var justString string

func someFunc() {
	//Создаём большую строку, длинной 2^10 степени, так как каждый сдвиг влево это умножение на 2
	v := createHugeString(1 << 10)
	//Получаем срез получившейся строки
	justString = v[:100]

	// re, _ := regexp.Compile(`[^\w]`)

	justString = ""
	fmt.Println(v)
	//Но если сам символ занимает больше чем 1 байт (в нашем случае это иероглиф), то будет получено не 100 смволов, а меньше
	//Посчитаем количество символов
	count := countSymbols(&justString)
	fmt.Println("До исправления строка была:", justString, "символов было:", count)
	//Для решения будем использовать руны, так как они могут вместить в себя такие большие символы
	var runes = make([]rune, 1<<10)
	//Преобразуем результат работы функции в руны, так как функция возвращает строку, поэтому нужно преобразовывать
	runes = []rune(createHugeString(1 << 10))
	//Теперь мы имеем руны, из рун получаем 100 символов, а уже потом конвертируем их в строку, тем самым мы точно получим 100 символов
	justString = string(runes[:100])
	count = countSymbols(&justString)
	fmt.Println("После исправления получилась строка:", justString, "символов стало:", count)
}

func main() {
	someFunc()
}
