package main

/* Реализовать все возможные способы остановки выполнения горутины.
 */

import (
	"fmt"
	"time"
)

// Будет горутиной, которую будем останавливать
func Worker(channel chan bool) {
	for {
		select {
		case <-channel: //Если из канала пришло какое-то значение, то останавливаем горутину
			fmt.Println("Exit")
			return
		default: //Просто будем писать в stdout, что-нибудь, имитируя работу горутины
			fmt.Println("Waiting")
			time.Sleep(30 * time.Microsecond)
		}
	}
}

func main() {
	channel := make(chan bool) //Создаём канал, который будем служить признаком остановки горутины

	go Worker(channel)          //Создаём горутину
	time.Sleep(3 * time.Second) //Имитируем работу горутины
	channel <- true             //Записываем в канал признак остановки

}
