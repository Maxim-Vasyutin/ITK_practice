package main

import (
	"fmt"
)

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	go generator(naturals)
	go obrabotchik(naturals, squares)

	for read := range squares {
		fmt.Println(read)
		//range всё доделает до завершения осн.горутины, потому что range
		//будет работать до конца цепочки значений переменной squares,
		//которая завершается после закрытия канала squares
	}
}

func obrabotchik(naturals <-chan int, squares chan<- int) {
	//Конструкция x := <- naturals, но считывает все числа из канала
	for x := range naturals {
		x = x * x
		squares <- x
	}
	close(squares)
}

func generator(naturals chan<- int) {
	for i := 0; i < 10; i++ {
		naturals <- i
	}
	//Не удаляет канал, а показывает, что в него больше не придут данные
	//Закрываем, чтобы явно показать, что данных больше не будет
	close(naturals)
}
