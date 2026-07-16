package main

import "fmt"

func CausePanic() {
	panic("что-то пошло не так")
}

// Когда происходит паника, go начинает сматвать стэк вызовов вверх
func HandlePanic() {
	//аргументы вычисляются сразу
	defer func() {
		r := recover()
		if r != nil {
			fmt.Println("Паника перехвачена: ", r)
		}
	}()

	CausePanic()

}

func main() {
	HandlePanic()
}
