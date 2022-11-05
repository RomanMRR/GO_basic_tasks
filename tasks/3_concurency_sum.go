package main

/*
Дана последовательность чисел: 2,4,6,8,10.
Найти сумму их квадратов(22+32+42….) с использованием конкурентных вычислений.


*/

import (
	"fmt"
)

// Функция для приёма данных из массива и запуска функция для подсчёта суммы
func SumWorker(numsCh chan []int, sumCh chan int) {
	for nums := range numsCh {
		sumCh <- sum(nums)

	}
}

// Функция для вычисления квадрата числа и суммы получившихся значений
func sum(nums []int) int {
	s := 0
	for _, n := range nums {
		s += n * n

	}

	return s
}

func main() {
	numbers := [5]int{2, 4, 6, 8, 10}
	numsCh := make(chan []int) //Создаём канал для записи в него массива
	sumCh := make(chan int)    //Создаём переменную для записи итоговой суммы

	go SumWorker(numsCh, sumCh)
	numsCh <- numbers[:] //Записываем в канал массив
	fmt.Println(<-sumCh)
}
