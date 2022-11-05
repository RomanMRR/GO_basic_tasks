package main

/* Дана последовательность температурных колебаний: -25.4, -27.0 13.0, 19.0, 15.5, 24.5, -21.0, 32.5.
Объединить данные значения в группы с шагом в 10 градусов. Последовательность в подмножноствах не важна.


Пример: -20:{-25.0, -27.0, -21.0}, 10:{13.0, 19.0, 15.5}, 20: {24.5}, etc.
*/

import "fmt"

func main() {

	temperatures := [8]float32{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	groups := make(map[int][]float32)

	for _, val := range temperatures {
		//распределяем по группам в зависимости от значения температуры
		switch {
		case val <= -20:
			groups[-20] = append(groups[-20], val)
		case val >= 10 && val < 20:
			groups[10] = append(groups[10], val)
		case val >= 20 && val < 30:
			groups[20] = append(groups[20], val)

		}
	}

	for key, value := range groups {
		fmt.Printf("%d: %v\n", key, value)
	}
}
