package main

import (
	"fmt"
)

func Level1() {
	defer func() {
		r := recover()

		if r != nil {
			fmt.Println("Паника обработана на уровне 1: ", r)
		}
	}()
	Level2()
}

func Level2() {
	defer func() {
		fmt.Println("Завершаем Level2")
	}()
	Level3()
}

func Level3() {
	panic("ошибка в Level3")
}

func main() {
	Level1()
}
