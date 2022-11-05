package main

/* Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество. */

import "fmt"

func main() {

	strings := [5]string{"cat", "cat", "dog", "cat", "tree"}

	//Создаём хэш-таблицу, к которой ключём будет строка, а значением - переменная типа bool
	set := make(map[string]bool)
	//Инициализауем хэш-таблицу, и так как ключ может встретиться только один раз, то тем самым получаем множество
	for _, val := range strings {
		set[val] = true
	}

	for key := range set {
		fmt.Println(key)
	}
}
