package main

/* Реализовать все возможные способы остановки выполнения горутины.
 */

import (
	"context"
	"fmt"
	"time"
)

// Будет горутиной, которую будем останавливать
func Worker(ctx context.Context) {
	for {
		select {
		case <-ctx.Done(): //Через одну секунду после запуска горутины, она завершится
			fmt.Println("Finish!")
			return
		default:
			fmt.Println("I am working!")
		}
	}
}

func main() { //Создаём канал, который будем служить признаком остановки горутины
	ctx, _ := context.WithTimeout(context.Background(), 1*time.Second) //Создаём контекст продолжительностью 3 секунды

	go Worker(ctx) //Создаём горутину

	time.Sleep(1 * time.Second) //горутина работает 1 секунду
	fmt.Println("Exit!")

	time.Sleep(1 * time.Second) //Для проверки, что на экран не выводится работа горутины
}
