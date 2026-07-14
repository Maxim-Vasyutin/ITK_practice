package main

import (
	"errors"
	"fmt"
)

var ErrNotVer = errors.New("недопустимый возраст")

type MyError struct {
	Code int
	Msg  string
}

func (e MyError) Error() string {
	return e.Msg
}

func SimpleError() error {
	return errors.New("простая ошибка")
}

func FormattedError(age int) error {
	return fmt.Errorf("ошибка: возраст %d. %w", age, ErrNotVer)
}

func StructError() error {
	return MyError{Code: 404, Msg: "не найдено"}
}
