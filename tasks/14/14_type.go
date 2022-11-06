package main

/*
Разработать программу, которая в рантайме способна определить тип переменной:
 int, string, bool, channel из переменной типа interface{}.

*/

import (
	"fmt"
	"reflect"
)

func TypeDetecter(a interface{}) {
	//Используем специльный метод для определния типа переменной
	fmt.Println(reflect.ValueOf(a).Type())
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
