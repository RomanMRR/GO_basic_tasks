package main

/* Разработать программу нахождения расстояния между двумя точками,
которые представлены в виде структуры Point с инкапсулированными параметрами x,y и конструктором.
*/

import (
	"fmt"
	"math"
)

// Стркуктура с координатами точки
type Point struct {
	x, y float64
}

// Конструктор
func NewPoint(x, y float64) *Point {
	return &Point{
		x: x,
		y: y,
	}
}

// Функция для нахождения расстояния между двумя точками по формуле
func GetDistance(point1, point2 *Point) float64 {
	return math.Sqrt(math.Pow(point1.x-point2.x, 2) + math.Pow(point1.y-point2.y, 2))
}

func main() {
	point1 := NewPoint(2, 5)
	point2 := NewPoint(2, 8)

	fmt.Println("Расстояние равно ", GetDistance(point1, point2))
}
