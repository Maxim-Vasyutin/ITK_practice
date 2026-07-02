package main

type Stacker interface {
	Push(v int)
	Pop() int
}

type stack struct {
	data []int
}

func (s *stack) Push(v int) {
	//append сам расширит базовый массив, если не хватит cap
	s.data = append(s.data, v)
}

func (s *stack) Pop() int {
	if len(s.data) == 0 {
		panic("стек пуст")
	}

	//LIFO: берём последний элемент и укорачиваем слайс на 1
	v := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return v
}

func New() *stack {
	return &stack{}
}

func main() {}