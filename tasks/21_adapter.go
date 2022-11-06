package main

import (
	"fmt"
	"strconv"
)

type PowerNumber interface { //Будет интерфейс
	PowerTwo() int
}

type Number struct { //Будет структуру
	data int //Содержащая число
}

func (c *Number) PowerTwo() int { //Эта структура удовлетворяет интерфейсу выше
	//И может возводить своё число во вторую степень
	return c.data * c.data
}

type String struct { //Но есть другая структура
	data string
}

func (c *String) ConvertToNumber() int { //Она интерфейсу не удовлетворяет,
	//но может конвертировать строку в число
	number, _ := strconv.Atoi(c.data)
	return number
}

type StringAdapter struct { //Будет адаптер, который будет работать с неудовлетворяющей интерфусу структурой
	objectString *String
}

func (c *StringAdapter) PowerTwo() int { //Адаптер уже удовлетворяет интерфейсу
	number := c.objectString.ConvertToNumber() //Он адаптирует данные структуры String так, чтобы можно было их
	//использовать с интерфейсом
	return number * number
}

func DoSomething(something PowerNumber) { //Эта функция работает с интерфейсом
	fmt.Printf("Получившееся число во второй степени это %d\n", something.PowerTwo())
}

func main() {
	number := &Number{data: 5} //Создаём структуру

	DoSomething(number) //И работает с ней, так как она удовлетворяет интерфейсу
	//функция успешно выполнит работу

	str := &String{data: "6"}      //Однако эта структура с функцией работть не сможет
	str_adapter := &StringAdapter{ //Поэтому заносим её в адаптер
		objectString: str,
	}

	DoSomething(str_adapter) //И работает с адаптером
}
