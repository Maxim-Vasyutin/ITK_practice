package main

import (
	"errors"
	"fmt"
)

type MyError struct {
	msg string
}

func (e MyError) Error() string {
	return e.msg
}

func handle() error {
	return MyError{msg: "ошибка имеется"}
}

func simpleError() error {
	return errors.New("это простая ошибка через пакет errors")
}

func formatError() error {
	str := "через пакет fmt"
	return fmt.Errorf("это ошибка с форматированием %s", str)
}

func main() {
	fmt.Println(handle())
	fmt.Println(simpleError())
	fmt.Println(formatError())
}
