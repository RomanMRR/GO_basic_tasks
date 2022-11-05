package main

/* Реализовать пересечение двух неупорядоченных множеств. */

import "fmt"

type void struct{}

// Функция для создания множества из слайса строк
func makeSet(strings []string) map[string]void {
	var member void

	//Создаём хэш-таблицу, к которой ключём будет строка, а значением - пустая структура
	set := make(map[string]void)
	//Инициализауем хэш-таблицу, и так как ключ может встретиться только один раз, то тем самым получаем множество
	//Приэтом в значениях лежит пустая структура, а потому такой способ занимает меньше памяти, чем использования bool
	for _, val := range strings {
		set[val] = member
	}

	return set
}

func main() {

	// Создаём пустую структура, которая почти не занимает память
	var member void

	strings1 := [5]string{"cat", "cat", "dog", "cat", "tree"}
	strings2 := [5]string{"cat", "cat", "dog1", "cat3", "tree"}

	set_strings1 := makeSet(strings1[:])
	set_strings2 := makeSet(strings2[:])
	intersection := make(map[string]void) //В данное множество будем класть одинаковые элементы

	for key := range set_strings1 {
		_, ok := set_strings2[key] //Если элемент из первого множества есть во втором
		if ok {
			intersection[key] = member //То мы добавляем элемент в третье множество
		}
	}

	for key := range intersection {
		fmt.Println(key)
	}
}
