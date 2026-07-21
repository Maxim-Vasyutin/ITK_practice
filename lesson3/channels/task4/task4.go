package main

import (
	"fmt"
	"sync"
	"time"
)

func mergeChannels(channels ...<-chan int) <-chan int {
	out := make(chan int)
	var wg sync.WaitGroup

	wg.Add(len(channels))
	//Первый цикл будет разбивать множество на подмножества (отдельные каналы)
	for _, ch := range channels {
		//Дальше делим разные каналы на горутины
		go func(ch <-chan int) {
			//Вытаскиваем значения из канала
			for v := range ch {
				out <- v
			}
			wg.Done()
		}(ch)
	}

	//фоном запускаем ожидание и закрытие (если все горутины завершились)
	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func main() {
	a := make(chan int)
	b := make(chan int)
	c := make(chan int)

	go func() {
		defer close(a)
		a <- 1000
		time.Sleep(3 * time.Second)
		a <- 2000

	}()
	go func() {
		defer close(b)
		b <- 11
		b <- 11
		b <- 11
	}()

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		close(c)
	}()

	out := mergeChannels(a, b, c)
	for v := range out {
		fmt.Println(v)
	}
}
