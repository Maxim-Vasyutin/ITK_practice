## Задание

### Цель
1. Создать интерфейс `PaymentProcessor` с методом `ProcessPayment(amount float64) error`.
2. Реализовать интерфейс для трех Банков:
    - **Sberbank**
    - **Tbank**
    - **Alfabank**

### Требования
- Каждый провайдер должен иметь уникальный идентификатор (например, `APIKey`).
- Метод `ProcessPayment` должен:
    - Возвращать `nil`, если сумма платежа положительная.
    - Возвращать ошибку `ErrInvalidAmount`, если сумма ≤ 0.
    - Возвращать ошибку `ErrProviderUnavailable`, если провайдер недоступен (заглушка). Сделать рандомный шанс, что банк недоступен.

    
package main

import "errors"

// Общие ошибки

var (
	ErrInvalidAmount       = errors.New("некорректная сумма платежа")
	ErrProviderUnavailable = errors.New("провайдер недоступен")
)

// PaymentProcessor - интерфейс обработки платежей

type PaymentProcessor interface {
	ProcessPayment(amount float64) error
}