package main

/*
Разработать программу, которая в рантайме способна определить тип переменной:
 int, string, bool, channel из переменной типа interface{}.

*/

import (
	"fmt"
)

func TypeDetecter(a interface{}) {
	//Здесб используем специальную конструкцию - переключатель типов
	//Switch здесь сравнивает по типу, а не по значению
	switch a.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case bool:
		fmt.Println("bool")
	default:
		fmt.Println("chan")
	}
}

func main() {
	a := 1
	TypeDetecter(a)
	b := false
	TypeDetecter(b)
	c := "String"
	TypeDetecter(c)
	var channel chan string
	TypeDetecter(channel)
}
