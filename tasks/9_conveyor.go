package main

/* Разработать конвейер чисел.
Даны два канала: в первый пишутся числа (x) из массива,
во второй — результат операции x*2, после чего данные из второго канала должны выводиться в stdout. */

import (
	"fmt"
)

// Функция для приёма данных из массива и запуска функция для операции x*2
func MulWorker(numsCh chan []int, mul2Ch chan int) {
	for nums := range numsCh {
		mulBy2(nums, mul2Ch)

	}
}

// Функция для вычисления квадрата числа и суммы получившихся значений
func mulBy2(nums []int, mul2Ch chan int) {
	//Получаем данные из первого канала и умножаем значения массива на 2
	for _, n := range nums {
		mul2Ch <- 2 * n
	}
	close(mul2Ch)
}

func main() {
	numbers := [5]int{2, 4, 6, 8, 10}
	numsCh := make(chan []int) //Создаём канал для записи в него массива
	mul2Ch := make(chan int)   //Создаём переменную для записи операции x*2

	go MulWorker(numsCh, mul2Ch)
	numsCh <- numbers[:] //Записываем в канал массив
	for i := range mul2Ch {
		fmt.Println(i)
	}
}
