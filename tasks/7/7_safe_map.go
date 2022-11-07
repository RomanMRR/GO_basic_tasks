package main

/*
Реализовать конкурентную запись данных в map.
*/

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
)

// Стркуктура с мьютексом внутри и map, что позволит реализовать безопасную запись данных в map
type SafeMap struct {
	mx sync.Mutex //По-умолчанию unlock
	m  map[string]int
}

// Конструктор
func NewSafeMap() *SafeMap {
	return &SafeMap{
		m: make(map[string]int),
	}
}

// Функция для конкурентной записи данных в map
func (c *SafeMap) Write(key string, value int) {
	c.mx.Lock() //Блокируем операцию на чтения и запись для других потоков
	c.m[key] = value
	c.mx.Unlock() //Разблокируем
}

func main() {
	MyMap := NewSafeMap()
	var wg sync.WaitGroup  //Для синхронизации потоков
	var sb strings.Builder //Реализует самый эффективный способ конкантенации

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			sb.WriteString("Key ")
			sb.WriteString(strconv.Itoa(i)) //Преобразуем число в string и соединяем с Key
			MyMap.Write(sb.String(), i)     //Записываем результат в map
			sb.Reset()                      //Сбрасываем получившуюся строку
			wg.Done()
		}(i)

	}

	wg.Wait()

	for key, value := range MyMap.m {
		fmt.Printf("Key is %s, value is %d\n", key, value)
	}
}
