Реализуйте структуру стека с использованием слайсов, удовлетворяющую следующему интерфейсу:

```go
type Stacker interface {
    Push(v int)
    Pop() int
}
```

### Требования к реализации

1. Операция Push(v int)
    Должна добавлять целочисленное значение v в стек.

2. Операция Pop() int Должна возвращать последний добавленный элемент, реализуя поведение LIFO (последним пришёл — первым ушёл).
    Если стек пуст, вызов метода Pop() должен приводить к панике.

3. Конструктор
    Реализуйте функцию New() *stack, возвращающую новый экземпляр стека.

4. Реализация должна находится в main.go
5. Реализация должна успешно проходить тесты. Для их запуска введите команду `go test ./...` в этой директории


```go
package main

import (
	"testing"
)

func TestStack_PushPop(t *testing.T) {
	s := New()

	s.Push(1)
	s.Push(2)
	s.Push(3)

	tests := []struct {
		expected int
	}{
		{3},
		{2},
		{1},
	}

	for _, tc := range tests {
		got := s.Pop()
		if got != tc.expected {
			t.Errorf("Pop() = %d; ожидалось %d", got, tc.expected)
		}
	}
}

func TestStack_PopEmpty(t *testing.T) {
	s := New()

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Ожидалась паника при попытке извлечь элемент из пустого стека")
		}
	}()

	s.Pop()
}
```

```go
package main

type Stacker interface {
	Push(v int)
	Pop() int
}

type stack struct {
	//...
}

func (s *stack) Push(v int) {
	panic("unimplemented")
}

func (s *stack) Pop() int {
	panic("unimplemented")
}

func New() *stack {
	return &stack{}
}
```

