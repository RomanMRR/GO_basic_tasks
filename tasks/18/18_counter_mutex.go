package main

/*
Реализовать структуру-счетчик,
 которая будет инкрементироваться в конкурентной среде.
 По завершению программа должна выводить итоговое значение счетчика.

*/

import (
	"fmt"
	"sync"
)

// Сама структура
type Counter struct {
	counter int
}

func main() {
	c := Counter{counter: 0} //Инициализируем структуру
	n := 200                 //Будем считать до n
	m := sync.Mutex{}        //Для обеспечение потокобезопасной записи
	wg := sync.WaitGroup{}   //ДЛя синхронизации потоков
	wg.Add(n)                //Столько будет горутин
	for i := 0; i < n; i++ {
		go func(i int) {
			m.Lock()    //Блокируем на чтение и запись из других потоков
			c.counter++ //Изменяем значение
			m.Unlock()  //Разблокируем
			wg.Done()
		}(i)
	}
	wg.Wait() //Ждём завершения всех горутин

	fmt.Println(c.counter) //Выводид значение n, если всё верно

}
