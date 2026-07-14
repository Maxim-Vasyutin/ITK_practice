package main

import (
	"errors"
	"fmt"
	"math/rand/v2"
)

var (
	ErrNotFound  = errors.New("ресурс не найден")
	TimeoutError = errors.New("таймаут операции")
)

type MyError struct {
	msg string
}

func (e MyError) Error() string {
	return e.msg
}

func ProcessError(err error) string {
	if errors.Is(err, TimeoutError) {
		return fmt.Sprintln("Требуется повторная проверка")
	}
	if errors.Is(err, ErrNotFound) {
		return fmt.Sprintln("Ресурс не найден")
	}
	return fmt.Sprintln("неизвестная ошибка")
}

func SimulateRequest() error {
	x := rand.IntN(100)

	if x < 50 {
		return fmt.Errorf("запрос не выполнен: %w", TimeoutError)
	} else if x >= 50 && x < 80 {
		return fmt.Errorf("ошибка: %w", ErrNotFound)
	} else if x >= 80 {
		return fmt.Errorf("неизвестная ошибка")
	}
	return nil
}

func main() {
	err := SimulateRequest()
	fmt.Println(ProcessError(err))
}
