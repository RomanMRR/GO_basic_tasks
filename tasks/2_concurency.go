package main

/*
Написать программу, которая конкурентно рассчитает значение квадратов чисел взятых из массива
(2,4,6,8,10) и выведет их квадраты в stdout.


*/

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup //Перпеменная для работы с группами горутин
	numbers := [5]int{2, 4, 6, 8, 10}

	for _, val := range numbers {
		wg.Add(1) //Увеличиваем внутрениий счётчик на 1
		go func(number int) {
			fmt.Printf("Квадрат числа %d равен %d\n", number, number*number)
			wg.Done() //Уменьшаем счётчик на 1
		}(val)
	}

	wg.Wait() //Ожидание, пока счётчик не станет равен 0
}
