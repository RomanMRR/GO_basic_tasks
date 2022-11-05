package main

/* Реализовать все возможные способы остановки выполнения горутины.
 */

import (
	"fmt"
	"time"
)

// Будет горутиной, которую будем останавливать
func Worker(channel chan string) {
	for {
		value, ok := <-channel
		fmt.Println(value)
		//Если канал закрыт, то завершаем горутину
		if !ok {
			fmt.Println("I've done")
			return
		}
		fmt.Println("Am I Working?")
	}
}

func main() {
	channel := make(chan string) //Создаём канал, при закрытии котого горутина будет останавливаться

	go Worker(channel) //Создаём горутину
	channel <- "Working!"
	channel <- "Still Working!"
	close(channel)
	time.Sleep(50 * time.Microsecond) //Небольшая задержка для проверки, что горутина остановлена
}
