package main

/*
Реализовать постоянную запись данных в канал (главный поток).
Реализовать набор из N воркеров, которые читают произвольные данные из канала и выводят в stdout.
Необходима возможность выбора количества воркеров при старте.

Программа должна завершаться по нажатию Ctrl+C. Выбрать и обосновать способ завершения работы всех воркеров.



*/

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

// Функция для приёма данных из основного потока
func Worker(id int, channel chan string, wg *sync.WaitGroup) {
	for data := range channel { //Горутина работает, пока открыт канал
		fmt.Printf("I am %d and I get %s\n", id, data)
	}

	fmt.Printf("I am %d and I've finished work\n", id)
	wg.Done()
}

func main() {
	groupWorkers := sync.WaitGroup{}        //Служит для создания группы горутин
	mainChannel := make(chan string)        //Основной поток
	var N int                               //Для приёма количества воркеров
	signalToStop := make(chan os.Signal, 1) //Создаём канал для приёма сигналов, указываем буфер размером один
	//ПОчему 1??? В документации сказано, что для таких каналов следует указывать буфер
	//И если канал ожидает только одно какое-то значение сигнала, то 1 достаточно
	signal.Notify(signalToStop, syscall.SIGINT) //Региструруем созданный канал на получение сигнала Ctrl+C, SIGINT как раз и нужен для этого

	fmt.Print("Введите количество воркеров: ")
	fmt.Scanf("%d\n", &N)

	for i := 0; i < N; i++ {
		groupWorkers.Add(1)
		go Worker(i, mainChannel, &groupWorkers)
	}

forLoop: //Для выхода из цикла ставим метку
	for {
		select {
		case <-signalToStop: //Получили Ctrl+C
			fmt.Println("We must stop!!!")
			close(mainChannel) //Закрываем канал, чтобы горутины также закрылись
			break forLoop      //Выходим из цикла

		default: //Отправляем в канал данные
			fmt.Println("Send data")
			mainChannel <- "some data"
			time.Sleep(1 * time.Second)
		}
	}

	groupWorkers.Wait() //Ожидаем, чтобы все горутины завершили работу
}
