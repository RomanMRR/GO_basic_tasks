package main

/*


Разработать программу, которая будет последовательно отправлять значения в канал, а с другой стороны канала — читать.
По истечению N секунд программа должна завершаться.




*/

import (
	"fmt"
	"sync"
	"time"
)

// Будет читать из канала
func ReadFromChannel(channel chan string, wg *sync.WaitGroup) {
	for data := range channel {
		fmt.Println("I get ", data)
	}
	fmt.Println("I finish")
	wg.Done()
}

// Усыпляет программы
func SleepProgram(channel chan string, finish chan bool, n time.Duration, wg *sync.WaitGroup) {
	time.Sleep(n * time.Second)
	finish <- true //Даёт команду главному потоку о завершении программы
	close(channel) //Даёт команду потоку чтения из канала о завершении работы
	fmt.Println("I finish")
	wg.Done()
}

func main() {
	var n time.Duration             //Количество секунд работы программы
	channel := make(chan string, 1) //Канал для записи
	finish := make(chan bool)       //Канал для сигнала об завершение программы
	var wg sync.WaitGroup           //Для синхронизации потоков
	wg.Add(2)                       //Будет ещё два потока

	fmt.Print("Seconds: ")
	fmt.Scan(&n)
	go ReadFromChannel(channel, &wg)
	go SleepProgram(channel, finish, n, &wg)

	// Отправляем данные в канал...
ForLabel:
	for {
		select {
		case <-finish: //Пока не получим сигнал о завершении работы
			break ForLabel
		default:
			channel <- "data"
			time.Sleep(30 * time.Millisecond)
		}
	}

	close(finish)
	wg.Wait() //Ждём, пока все потоки завершат работу
	fmt.Println("It's all")

}
