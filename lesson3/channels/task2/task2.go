package main

import (
	"time"
)

func worker() chan int {
	ch := make(chan int)
	go func() {
		time.Sleep(3 * time.Second)
		ch <- 42
	}()
	return ch
}
func main() {
	timeStart := time.Now()
	work1 := worker()
	work2 := worker()

	_, _ = <-work1, <-work2
	println(int(time.Since(timeStart).Seconds()))
}

//В первоначальном варинанте функция выполнялась 6 секунд,
// потому что строка _, _ = <-worker(), <-worker(), будет делать вызовы поочерёдно,
// ожидая завершения прошлого вызова

//в нынешнем варианте - сначла мы запустили воркеры, а потом начинаем ждать их выполнения
//!Чтение из небуферезированного канала останавливает текущую горутину!