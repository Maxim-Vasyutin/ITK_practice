package main

// RemoveUnordered удаляет элемент по индексу без сохранения порядка.
// Если индекс выходит за границы слайса, возвращает исходный слайс.
func RemoveUnordered[T any](s []T, i int) []T {
	// реализовать
	if i >= len(s) || i < 0 {
		return s
	}

	s[i] = s[len(s)-1]
	s = s[:len(s)-1]
	return s
}

// RemoveOrdered удаляет элемент по индексу с сохранением порядка.
// Если индекс выходит за границы слайса, возвращает исходный слайс.
func RemoveOrdered[T any](s []T, i int) []T {
	// реализовать
	if i >= len(s) || i < 0 {
		return s
	}

	for j := i; j < len(s)-1; j++ {
		s[j] = s[j+1]
	}
	s = s[:len(s)-1]
	return s
}

// RemoveAllByValue удаляет все вхождения указанного значения.
func RemoveAllByValue[T comparable](s []T, value T) []T {
	// реализовать
	var count int
	for _, v := range s {
		if v != value {
			s[count] = v
			count++
		}
	}
	s = s[:count]
	return s
	//В цикле происходит перезапись оригинального слайса
	//А count переопределяет длину уже переработанного слайса
	//И в конце ф-ии у нас в начале слайса будут нужные нам значения, а count нужен для границы этих значений
}

// RemoveDuplicates оставляет только уникальные элементы (сохраняет порядок).
func RemoveDuplicates[T comparable](s []T) []T {
	// реализовать
	//[1,2,3,1,2]
	//unique {
	//[1] = 1
	//[2] = 1
	//[3] = 1
	//}
	//перезаписываю слайс на месте, count - граница уникальных значений
	unique := make(map[T]bool)
	var count int
	for _, v := range s {
		if unique[v] {
			continue //дубликат - пропускаем
		}
		unique[v] = true
		s[count] = v
		count++
	}
	return s[:count]
}

// RemoveIf удаляет элементы, удовлетворяющие условию predicate.
func RemoveIf[T any](s []T, predicate func(T) bool) []T {
	// реализовать
		var count int
	for _, v := range s {
		if !predicate(v) {
			s[count] = v
			count++
		}
	}
	return s[:count]
	return s
}

// RemoveOrderedWithNil удаляет элемент по индексу (для слайса указателей),
// обнуляя удаляемый элемент для предотвращения утечек памяти.
func RemoveOrderedWithNil[T any](s []*T, i int) []*T {
	//реализовать
	if i >= len(s) || i < 0 {
		return s
	}

	for j := i; j < len(s)-1; j++ {
		s[j] = s[j+1]
	}
	//После сдвига последний элемент дублируется (он же теперь на позиции len-2)
	//Если его не занулить - базовый массив продолжит держать указатель
	//и GC не сможет освободить объект, хотя из слайса мы его "удалили"
	s[len(s)-1] = nil
	s = s[:len(s)-1]
	return s
}

// ShrinkCapacity сокращает вместимость слайса, если она превышает
// удвоенную длину после удаления элементов.
func ShrinkCapacity[T any](s []T) []T {
	//реализовать
	if cap(s) <= 2*len(s) {
		return s
	}
	//Единственный способ уменьшить cap - создать новый слайс и скопировать
	//s[:len(s):len(s)] не подойдёт
	//потому что, прошлый массив всё равно остаётся в памяти
	shrunk := make([]T, len(s))
	copy(shrunk, s)
	return shrunk
}

func main() {
	//реализовать
		//1. RemoveUnordered - на место удалённого встаёт последний
	a := []int{1, 2, 3, 4, 5}
	a = RemoveUnordered(a, 1)
	fmt.Println("RemoveUnordered(i=1):", a) // [1 5 3 4]

	//2. RemoveOrdered - порядок сохраняется
	b := []int{1, 2, 3, 4, 5}
	b = RemoveOrdered(b, 1)
	fmt.Println("RemoveOrdered(i=1):", b) // [1 3 4 5]

	//выход за границы - вернётся исходный слайс
	b = RemoveOrdered(b, 10)
	fmt.Println("RemoveOrdered(i=10):", b) // [1 3 4 5]

	//3. RemoveAllByValue
	c := []int{1, 2, 1, 3, 1, 4}
	c = RemoveAllByValue(c, 1)
	fmt.Println("RemoveAllByValue(1):", c) // [2 3 4]

	//4. RemoveDuplicates
	d := []int{1, 2, 3, 1, 2, 4}
	d = RemoveDuplicates(d)
	fmt.Println("RemoveDuplicates:", d) // [1 2 3 4]

	//5. RemoveIf - удаляю чётные
	e := []int{1, 2, 3, 4, 5, 6}
	e = RemoveIf(e, func(v int) bool { return v%2 == 0 })
	fmt.Println("RemoveIf(чётные):", e) // [1 3 5]

	//6. RemoveOrderedWithNil - слайс указателей
	x, y, z := 10, 20, 30
	p := []*int{&x, &y, &z}
	p = RemoveOrderedWithNil(p, 1)
	fmt.Println("RemoveOrderedWithNil(i=1):", *p[0], *p[1]) // 10 30

	//7. ShrinkCapacity
	f := make([]int, 100)
	f = f[:10] //len=10, cap=100 -> cap больше чем 2*len
	fmt.Println("до Shrink: len =", len(f), "cap =", cap(f))
	f = ShrinkCapacity(f)
	fmt.Println("после Shrink: len =", len(f), "cap =", cap(f)) // 10 10
}
