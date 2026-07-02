Этот проект демонстрирует различные способы работы со слайсами в Go, включая очистку, обнуление и особенности внутренней структуры.

**Ваша задача:** Определить вывод каждого случая и зафиксировать ответы **в сообщениях коммитов** с кратким объяснением логики.
package main


```go
package main

import (
	"fmt"
	"unsafe"
)

func main() {
	//1
	first := []int{1, 2, 3, 4, 5}
	first = nil
	fmt.Println("first:", first, ":", len(first), ":", cap(first))
	//Вывод:
	// first: [] : 0 : 0

	//nil обнуляет весь заголовок слайса: указатель, len и cap
	//Связь с базовым массивом полностью теряется

	//2
	second := []int{1, 2, 3, 4, 5}
	second = second[:0]
	fmt.Println("second:", second, ":", len(second), ":", cap(second))
	//Вывод:
	// second: [] : 0 : 5

	//Срез до нуля обнуляет только len. cap остаётся 5, потому что
	//указатель на базовый массив никуда не делся - данные всё ещё там
	//Через second[:5] их можно вернуть

	//3
	third := []int{1, 2, 3, 4, 5}
	clear(third)
	fmt.Println("third:", third, ":", len(third), ":", cap(third))
	//Вывод:
	// third: [0 0 0 0 0] : 5 : 5

	//clear не удаляет элементы, а записывает zero value

	//4
	fourth := []int{1, 2, 3, 4, 5}
	clear(fourth[1:3])
	fmt.Println("fourth:", fourth, ":", len(fourth), ":", cap(fourth))
	//Вывод:
	// fourth: [1 0 0 4 5] : 5 : 5

	//fourth[1:3] - это тот же базовый массив. clear трогает только это окно,
	//а результат видим через fourth, потому что массив общий

	//5
	slice := make([]int, 3, 6)
	array := [3]int(slice[:3])
	slice[0] = 10

	fmt.Println("slice = ", slice, len(slice), cap(slice))
	fmt.Println("array =", array, len(array), cap(array))
	//Вывод:
	// slice =  [10 0 0] 3 6
	// array = [0 0 0] 3 3

	//Конвертация [3]int(slice[:3]) - это КОПИРОВАНИЕ данных в новый
	//массив, а не создание ссылки. Поэтому slice[0] = 10 после
	//конвертации на array уже не влияет
	//cap у массива всегда равен len - у массива нет запаса,
	//это фиксированный блок памяти



	//6 В каких случаях Slice пустой или нулевой
	//1
	var data []string
	fmt.Println("var data []string:")
	fmt.Printf("\tempty=%t nil=%t size=%d data=%p\n", len(data) == 0, data == nil, unsafe.Sizeof(data), unsafe.SliceData(data))
	//Вывод:
	// empty=true nil=true size=24 data=0x0

	//nil слайс: заголовок есть, но указатель на данные нулевой

	//2
	data = []string(nil)
	fmt.Println("data = []string(nil):")
	fmt.Printf("\tempty=%t nil=%t size=%d data=%p\n", len(data) == 0, data == nil, unsafe.Sizeof(data), unsafe.SliceData(data))
	//Вывод:
	// empty=true nil=true size=24 data=0x0

	//Явная конвертация nil - тот же nil слайс, что и в прошлой таске

	//3
	data = []string{}
	fmt.Println("data = []string{}")
	fmt.Printf("\tempty=%t nil=%t size=%d data=%p\n", len(data) == 0, data == nil, unsafe.Sizeof(data), unsafe.SliceData(data))
	//Вывод:
	//empty=true nil=false size=24 data=0x...

	//Пустой слайс: len=0, но указатель уже не нулевой

	//4
	data = make([]string, 0)
	fmt.Println("data =make([]string,0)")
	fmt.Printf("\tempty=%t nil=%t size=%d data=%p\n", len(data) == 0, data == nil, unsafe.Sizeof(data), unsafe.SliceData(data))
	//Вывод:
	// empty=true nil=false size=24 data=0x...
	

	empty := struct{}{}
	fmt.Println("empty struct address ", unsafe.Pointer(&empty))
	//!!!
	//все аллокации нулевого размера
	//в Go указывают на один спец. адрес в рантайме (zerobase),
	//чтобы не тратить память на пустоту
```