package main

/*
Дана структура Human (с произвольным набором полей и методов).
Реализовать встраивание методов в структуре Action от родительской структуры Human (аналог наследования).


*/

import "fmt"

type Human struct {
	Name string
	Year int
	Post string
}

type Action struct {
	Human
}

// String возвращает информацию о человеке.
func (human Human) String() string {
	return fmt.Sprintf("Имя: %s, Год рождения: %d, Должность: %s", human.Name, human.Year, human.Post)
}

// Print выводит информацию о человеке
func (human Human) Print() {
	fmt.Println(human.String())
}

func main() {
	action := Action{Human{"Максим", 1987, "Go-разработчик"}}
	action.Print()
}
