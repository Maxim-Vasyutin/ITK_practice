Этот проект демонстрирует различные способы работы со слайсами в Go, включая очистку, обнуление и особенности внутренней структуры.

**Ваша задача:** Определить вывод каждого случая и зафиксировать ответы **в сообщениях коммитов** с кратким объяснением логики.
package main


```go
import (
	"fmt"
	"unsafe"
)

func main() {
	//1
	first := []int{1, 2, 3, 4, 5}
	first = nil
	fmt.Println("first:", first, ":", len(first), ":", cap(first))
	//nil: 0 : 0 
	//Отпускаем привязку к массиву

	//2
	second := []int{1, 2, 3, 4, 5}
	second = second[:0]
	fmt.Println("second:", second, ":", len(second), ":", cap(second))
	//0 : 0 : 5
	//переиспользуем срез

	//3
	third := []int{1, 2, 3, 4, 5}
	clear(third)
	fmt.Println("third:", third, ":", len(third), ":", cap(third))
	//clear - обнуляет значения элементов
	//[0,0,0,0,0] : 5 : 5

	//4
	fourth := []int{1, 2, 3, 4, 5}
	clear(fourth[1:3])
	fmt.Println("fourth:", fourth, ":", len(fourth), ":", cap(fourth))
	//[1,0,0,4,5] : 5 : 5

	//5
	slice := make([]int, 3, 6)
	array := [3]int(slice[:3])
	//Тут снимается копия со среза. Массив - это значение
	slice[0] = 10

	fmt.Println("slice = ", slice, len(slice), cap(slice))
	fmt.Println("array =", array, len(array), cap(array))
	//slice = [10, 0, 0], 3, 6
	//array = [0, 0, 0], 3, 3

	//6 В каких случаях Slice пустой или нулевой
	//1
	var data []string
	fmt.Println("var data []string:")
	fmt.Printf("\tempty=%t nil=%t size=%d data=%p\n", len(data) == 0, data == nil, unsafe.Sizeof(data), unsafe.SliceData(data))
	//true, false, 24, ненулевой адрес

	//про Sizeof: Не знал, что 3 поля будут равны по 8 байт
	//Sizeof считает размер типа, а не наполнения

	//про SliceData - показывает куда смотрит заголовок. Поэтому значение - это указатель (%p)

	//2
	data = []string(nil)
	fmt.Println("data = []string(nil):")
	fmt.Printf("\tempty=%t nil=%t size=%d data=%p\n", len(data) == 0, data == nil, unsafe.Sizeof(data), unsafe.SliceData(data))
	//true, true, 24, нулевой адрес	

	//3
	data = []string{}
	fmt.Println("data = []string{}")
	fmt.Printf("\tempty=%t nil=%t size=%d data=%p\n", len(data) == 0, data == nil, unsafe.Sizeof(data), unsafe.SliceData(data))
	//true, false, 24, ненулевой адрес

	//4
	data = make([]string, 0)
	fmt.Println("data =make([]string,0)")
	fmt.Printf("\tempty=%t nil=%t size=%d data=%p\n", len(data) == 0, data == nil, unsafe.Sizeof(data), unsafe.SliceData(data))
	//true, false, 24, ненулевой адрес

	//5
	empty := struct{}{}
	fmt.Println("empty struct address ", unsafe.Pointer(&empty))
	//какой-то zerobase/nilpointer (заглушка)
}
```