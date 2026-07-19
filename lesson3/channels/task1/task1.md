## Задание: Анализ и исправление кода с гонками данных

### Описание задачи
1. Внимательно изучить код.
2. Найти все ошибки, описать их в комментариях прямо в коде.
3. Исправить код, обеспечив корректную работу.

```golang
package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func main() {
	alreadyStored := make(map[int]struct{})
	//аллокация памяти
	capacity := 1000
	doubles := make([]int, 0, capacity)

	for i := 0; i < capacity; i++ {
		doubles = append(doubles, rand.Intn(10))
	}
	//заполняем рандомами 1000 элементов

	uniqueIDs := make(chan int, capacity)
	wg := sync.WaitGroup{}

	for i := 0; i < capacity; i++ {
		//Это идиома, для явного обновления счётчика
		//Новая переменная на итерацию
		//Работала строка до 1.22 версии
		i := i
		
		//запускаем wait
		wg.Add(1)
		go func() {
			//дефер вызовится после анонимной ф-ии
			defer wg.Done()
			//мапа по рандомным значениям
			if _, ok := alreadyStored[doubles[i]]; !ok {
				//не важно значение, важно наличие
				alreadyStored[doubles[i]] = struct{}{}
				uniqueIDs <- doubles[i]
			}
		}()
		//Сделали wait
	}
	wg.Wait()

	for val := range uniqueIDs {
		fmt.Println(val)
	}
	//Вывод адреса?
	fmt.Println(uniqueIDs)

	//Канал не заркыт
}
```