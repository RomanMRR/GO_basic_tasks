package main

/* Реализовать все возможные способы остановки выполнения горутины.
 */

import (
	"fmt"
	"time"
)

// Будет горутиной, которую будем останавливать
func Worker(channel chan bool) {
	for range channel { //Как только канал закроется, горутина завершит работу
		fmt.Println("I am working!")
	}
	for {
		fmt.Println("Am I working?")
	}

}

func main() {
	channel := make(chan bool) //Создаём канал, который будем служить признаком остановки горутины

	go Worker(channel)
	channel <- true             //Создаём горутину
	time.Sleep(1 * time.Second) //Имитируем работу горутины
	//Записываем в канал признак остановки
	close(channel)

}
