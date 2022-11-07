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

func main() {
	var MyMap sync.Map     //Используем особую структуру - потокобезопасный Map
	var wg sync.WaitGroup  //Для синхронизации потоков
	var sb strings.Builder //Реализует самый эффективный способ конкантенации

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			sb.WriteString("Key ")
			sb.WriteString(strconv.Itoa(i)) //Преобразуем число в string и соединяем с Key
			MyMap.Store(sb.String(), i)     //Записываем результат в map
			sb.Reset()                      //Сбрасываем получившуюся строку
			wg.Done()
		}(i)

	}

	wg.Wait()

	MyMap.Range(func(key, value any) bool { //Итерируемся по map
		fmt.Println("Key is ", key, "value is ", value) //В этом методе передаётся функция, которая выполняется для
		//каждого элемента map
		return true //Возвращаем true, если будет false, то Range завершится
	})
}
