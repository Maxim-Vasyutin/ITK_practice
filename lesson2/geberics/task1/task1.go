package main

type Stack[T any] struct {
	elements []T
}

func NewStack[T any]() *Stack[T] {

	return &Stack[T]{}
}

func (t *Stack[T]) Push(value T) {
	t.elements = append(t.elements, value)
}

func (t *Stack[T]) Pop() (T, bool) {
	var x T
	if len(t.elements) == 0 {
		return x, false
	}

	x = t.elements[len(t.elements)-1]
	t.elements = t.elements[:len(t.elements)-1]
	return x, true
}

func (t *Stack[T]) Peek() (T, bool) {
	var x T
	if len(t.elements) == 0 {
		return x, false
	}
	x = t.elements[len(t.elements)-1]
	return x, true
}

func (t *Stack[T]) IsEmpty() bool {
	return len(t.elements) == 0
}
