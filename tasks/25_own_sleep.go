package main

/* Реализовать собственную функцию sleep. */

import (
	"fmt"
	"time"
)

// СОбственная реализация Sleep
func Sleep(duration time.Duration) {
	<-time.After(duration) //Данная функция ждёт столько времени, сколько ей указали
}

func main() {
	fmt.Println("Сейчас ", time.Now().Second(), " секунд")
	Sleep(2 * time.Second)
	fmt.Println("А сейчас ", time.Now().Second(), " секунд")

}
