package main

import "fmt"

func SafeDivide(a, b int) (num int) {
	defer func() {
		r := recover()
		if r != nil {
			num = 0
		}
	}()

	if b == 0 {
		panic("деление на ноль")
	}

	num = a / b

	return num
}

func main() {
	b := SafeDivide(10, 0)
	a := SafeDivide(10, 2)

	fmt.Println(b)
	fmt.Println(a)
}
