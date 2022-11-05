package main

/* Разработать программу, которая проверяет, что все символы в строке уникальные (true — если уникальные, false etc).
 Функция проверки должна быть регистронезависимой.

Например:
abcd — true
abCdefAaf — false
	aabcd — false

*/

import (
	"fmt"
	"strings"
)

// Используем пустую структура для меньшего потребления памяти
type void struct {
}

// Функция для проверки уникальности символов
func CheckUnique(str []rune) bool {
	symbols := make(map[rune]void)
	str = []rune(strings.ToLower(string(str[:]))) //Используем руны, чтобы поддерживался Unicode

	for _, i := range str {
		_, ok := symbols[i]
		if ok { //Если такой символ был, значит в строке есть неуникальные символы
			return false
		} else { //Иначе запоминаем символ
			symbols[i] = void{}
		}
	}

	return true
}

func main() {
	string1 := "abcd"
	string2 := "abCdefAaf"
	string3 := "aabcd"
	string4 := "Как дела?"

	fmt.Printf("%s - %t\n", string1, CheckUnique([]rune(string1)))
	fmt.Printf("%s - %t\n", string2, CheckUnique([]rune(string2)))
	fmt.Printf("%s - %t\n", string3, CheckUnique([]rune(string3)))
	fmt.Printf("%s - %t\n", string4, CheckUnique([]rune(string4)))
}
