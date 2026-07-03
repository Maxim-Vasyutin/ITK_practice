package main

type Stacker interface {
	Push(v int)
	Pop() int
}

type stack struct {
	//...
	data []int
}

func (s *stack) Push(v int) {

	s.data = append(s.data, v)
}

func (s *stack) Pop() int {
	if len(s.data) == 0 {
		panic("stack is empty")
	}
	last := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]

	//Обнулять элемент нужно только если элементы ссылочные
	//Поэтому можно просто срезать последний (верхний)

	return last
}

func New() *stack {
	
	return &stack{}
}
