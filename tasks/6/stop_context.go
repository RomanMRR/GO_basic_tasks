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
		case <-ctx.Done():
			fmt.Println("Finish!")
			return
		default:
			fmt.Println("I am working!")
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background()) //Создаём контекст с функцией отмены

	go Worker(ctx) //Создаём горутину

	time.Sleep(3 * time.Second) //Имитируем работу
	cancel()
	time.Sleep(1 * time.Second) //Для проверки, что на экран не выводится работа горутины
	fmt.Println("Exit!")
}
