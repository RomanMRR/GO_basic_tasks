package main

/* Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0. */

import "fmt"

// Устанавливает нужный бит числа в единицу
func setBit(number int64, k int) int64 {
	mask := int64(1) << k  //Создаём маску, в которой на нужном месте стоит 1
	number = number | mask //Устанавливаем нужный бит в 1, так как если на k месте будет 0, то он благодаря маске превратится в 1
	return number
}

func clearBit(number int64, k int) int64 {
	mask := ^(int64(1) << k) //Создаём маску, в котором на нужном месте стоит 0
	number &= mask           //Устанавливаем нужный бит в 0, так как если на k месте будет стоять 1, но благодаря операции AND он станет 0
	return number
}

func main() {
	var number int64
	var k int

	fmt.Scan(&number, &k)

	number = setBit(number, k)
	fmt.Printf("Число после установки %d бита: %d\n", k, number)
	number = clearBit(number, k)
	fmt.Printf("Число после сброса %d бита: %d\n", k, number)
}
