package main

import (
	"errors"
	"fmt"
	"math/rand"
)

// Общие ошибки

var (
	ErrInvalidAmount       = errors.New("некорректная сумма платежа")
	ErrProviderUnavailable = errors.New("провайдер недоступен")
)

// PaymentProcessor - интерфейс обработки платежей

type PaymentProcessor interface {
	ProcessPayment(amount float64) error
}

type Sberbank struct {
	APIKey string
}

type Tbank struct {
	APIKey string
}

type Alfabank struct {
	APIKey string
}

func (s Sberbank) ProcessPayment(amount float64) error {
	if amount <= 0 {
		return ErrInvalidAmount
	}

	if rand.Float64() < 0.3 {
		return ErrProviderUnavailable
	}

	fmt.Printf("PP: %s. \nHave a %f\n", s.APIKey, amount)
	return nil
}

func (t Tbank) ProcessPayment(amount float64) error {
	if amount <= 0 {
		return ErrInvalidAmount
	}

	if rand.Float64() < 0.3 {
		return ErrProviderUnavailable
	}

	fmt.Printf("PP: %s. \nHave a %f\n", t.APIKey, amount)
	return nil
}

func (a Alfabank) ProcessPayment(amount float64) error {
	if amount <= 0 {
		return ErrInvalidAmount
	}

	if rand.Float64() < 0.9 {
		return ErrProviderUnavailable
	}

	fmt.Printf("PP: %s. \nHave a %f\n", a.APIKey, amount)
	return nil
}

func main() {

	sber := Sberbank{"sber_api_123"}
	tbank := Tbank{"tbank_api_345"}
	alfa := Alfabank{"alfa_api_567"}

	/*
		err := sber.ProcessPayment(500)
		if err != nil {
			fmt.Println("Sber: ", err)
		}

		err = tbank.ProcessPayment(-200)
		if err != nil {
			fmt.Println("Tbank: ", err)
		}

		err = alfa.ProcessPayment(500)
		if err != nil {
			fmt.Println("Alfa: ", err)
		}
	*/

	/*
		Создаю срез структур (sber, tbank, alfa), которые реализовывают
		интерфейс (PaymentProcessor).
		Реализовывают через метод ProcessPayment.
		p - вызов структуры, чья очередь сейчас выпала.

		Одна операция проходит через все банки из среза.
	*/
	processors := []PaymentProcessor{sber, tbank, alfa}

	for _, p := range processors {
		if err := p.ProcessPayment(500); err != nil {
			fmt.Println("ошибка:", err)
			continue
		}
		fmt.Println("платёж положительный")
	}
}
