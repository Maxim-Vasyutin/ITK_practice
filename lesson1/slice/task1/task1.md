Это задание направлено на глубокое понимание работы срезов (slices), их модификации и передачи в функциях Go.  
**Ваша задача:** Определить вывод каждой из предложенных программ и зафиксировать ответы **в сообщениях коммитов** с кратким объяснением логики.
### 1.
```go 
package main

import "fmt"

type account struct {
	value int
}

func main() {
	s1 := make([]account, 0, 2) //срез структур с длинной 0 и вместимостью 2
	s1 = append(s1, account{})  //s1[0] - пустая структура.
	s2 := append(s1, account{}) // Создаётся ещё срез, такой же вместиомсти. Второе значение - пустая структура.
	acc := &s2[0]               //acc = первое значение s2. по указателю
	acc.value = 100             //acc = 100

	fmt.Println(s1, s2)


	//Вывод:
	// 100    //Второе значение не было инициализированно
	// 100 0  //2 среза делять общий массив (пока вместимость позволяет)

	s1 = append(s1, account{}) //s1{0,0}
	acc.value += 100           //acc = 200
	fmt.Println(s1, s2)


	//Вывод:
	// 200 0
	// 200 0
}

```
-----
2.
```go
package main

import "fmt"

func main() {
	slice := make([]string, 0, 5)
	slice = append(slice, "0", "1", "2", "3")
	fmt.Println(slice, len(slice), cap(slice)) // [0, 1, 2, 3], 4, 5
	addToSlice1(slice)
	fmt.Println(slice, len(slice), cap(slice)) //[0, 1, 2, one], 4, 5
	addToSlice2(slice)
	fmt.Println(slice, len(slice), cap(slice)) // [0, 1, 2, one, two], 5, 5
}

func addToSlice1(slice []string) {
	slice = append(slice[1:3], "one")
}

// Так, как 3ий элемент не включительно - он затирается

func addToSlice2(slice []string) {
	slice = append(slice, "two")
}
//Тут ошибся. Разобрал с ИИ, что происходит изменение среза внутри функции, а базовый массив не изменяется

```
---
3.
```go
package main

import "fmt"

func main() {
	a1 := make([]int, 0, 10)
	a1 = append(a1, []int{1, 2, 3, 4, 5}...)
	a2 := append(a1, 6)
	// 1,2,3,4,5,6
	a3 := append(a1, 7)
	// 1,2,3,4,5,7
	fmt.Println(a1, a2, a3)
	// Вывод:
	// 1,2,3,4,5
	// 1,2,3,4,5,7
	// 1,2,3,4,5,7

	// Так произойдёт, потому что происходит изменение одного и того же среза
	// а2 = [а1{1,2,3,4,5}, 6]
	// а3 = [а1{1,2,3,4,5}, 7]
}

```
---
4.
```go
package main

import "fmt"

func main() {
	a := []int{1, 2, 3}
	b := a[:2]
	// Так как берём в срез не включительно 2ой элемент, после изменений среза b, срез a тоже поменяется
	b = append(b, 4)
	fmt.Println(b) //1,2,4
	fmt.Println(a) //1,2,4
}

```
-----
5.
```go
package main

import "fmt"

func main() {
	arr := []int{1, 2, 3}
	src := arr[:1] //src = [1]

	foo(src)
	fmt.Println(src) // [1]
	fmt.Println(arr) // [1,5,3]
}

func foo(src []int) {
	src = append(src, 5)
	//[1,5]
}

//Меяется базовый массив.
//Внутри foo не пересоздаётся
```
-----
6.
```go
package main

import "fmt"

func main() {
	arr := [5]int{1, 2, 3, 4, 5} // len = 5, cap = 5
	bar := arr[1:3]
	bar = append(bar, 10, 11, 12, 13)
	fmt.Println(arr, bar) //
}

//bar
//len = 2, cap = 5 ->
//-> bar = 2, 3, 10, 11, 12, 13 (len = 6, cap = 10)
//Базовый массив пересоздался, потому что мы вышли за cap

//Вывод:
//[1,2,3,4,5] [2,3,10,11,12,13]

```
-----
7.
```go
package main

import "fmt"

func main() {
	a := []string{"a", "b", "c"}   // len = 3, cap = 3
	b := a[1:2]                    // b = [b]
	fmt.Println(b, cap(b), len(b)) // [b], cap = 2, len = 1
	b[0] = "q"                     // b = [q]
	fmt.Println(a)                 // a = [q, b, c]
}

// Увидел, что ошибка, не совсем понял почему.

//Видимо, слайс всегда указывает на базовый массив, пока базовый не пересоздастся
//b[0] это a[1]
```
---
8.
```go
package main

import (
	"fmt"
)

func main() {
	nums := make([]int, 1, 3)
	fmt.Println(nums) // 0
	appendSlice(nums, 1)
	fmt.Println(nums) // 0. В ф-ии срез изменился локально
	copySlice(nums, []int{2, 3})
	fmt.Println(nums) // 0
	mutateSlice(nums, 1, 4)
	fmt.Println(nums) // 0, 4 (будет паника а не значения)
}

func appendSlice(sl []int, val int) {
	sl = append(sl, val)
	//sl = 0, 1
}

func copySlice(sl, cp []int) {
	copy(sl, cp)
	//Не знал, что копирование происходит по минимальной длине среза.
}

func mutateSlice(sl []int, idx, val int) {
	sl[idx] = val
	//Ошибся.
	//У нас доступен только первый элемент исходного массива (индекс = 0)
	//Поэтому передача значения индекса 1 означает выход за пределы
}

```
---
9.
```go
package main

import (
	"fmt"
)

func main() {
	slice := make([]int, 3, 4)
	appendingSlice(slice[:1])
	fmt.Println(slice) // 0 1 0
}

func appendingSlice(slice []int) {
	slice = append(slice, 1)
}
```