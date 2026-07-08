Это задание направлено на понимание работы срезов, функций и передачи данных в Go.  
**Ваша задача:**
1. **Проанализировать вывод программ** и объяснить поведение кода.
2. **Исправить код** так, чтобы достигался корректный результат (в некоторых случаях требуется несколько решений)


### 1.// Версия 1.21
```go
package main

import (
	"fmt"
)

func main() {
	var numbers []*int
	for _, value := range []int{10, 20, 30, 40} {
		//(фикс)
		//v := value 
		//numbers = append(numbers, &v)
		numbers = append(numbers, &value)
	}
	for _, number := range numbers {
		fmt.Println("d", *number)
	}
}
```
----
### 2.
```go
package main

import (
	"fmt"
	"strings"
)

func chengeSlice(arr []string) {
	arr[0] = "Goodbye"
}

func appendSomeData(arr []string) []string {
	arr = append(arr, "!")
	return arr //
}

func main() {
	someSlice := []string{"Hello", "World"}
	chengeSlice(someSlice)
	someSlice = appendSomeData(someSlice)
	fmt.Println(strings.Join(someSlice, ""))
}
//без return в appendSomeData новый пересозданный массив не попадал наружу
```
----
### 3.
```go
package main

import "fmt"

func test(testSlice []string) []string {
	testSlice = append(testSlice, "Пока")
	return testSlice
}
func main() {
	testSlice := make([]string, 0, 3)
	testSlice = append(testSlice, "Привет")
	testSlice = append(testSlice, "Привет")

	testSlice = test(testSlice)
	fmt.Println(testSlice)
	
}
//добавил явный возврат среза из test
//В main len 3 не читался
//testSlice = testSlice[:3]
```
----
### 4.
```go
package main

import "fmt"

func main() {
	first := []int{10, 20, 30, 40}
	second := make([]*int, len(first))
	for i, v := range first {
		//
		second[i] = &v
	}
	fmt.Println(*second[0], *second[1])
}
//в 0 и в 1 находить одно адресное пространство
```
----
### 5.
```go
package main

import (
	"fmt"
)

func main() {
	slice := make([]string, 3, 4)
	fmt.Println(slice)

	appendSlice(slice)
	slice = slice[:4]
	fmt.Println(slice)

	mutareSlice(slice)
	fmt.Println(slice)
}

func appendSlice(slice []string) {
	slice = append(slice, "privet")
}
func mutareSlice(slice []string) {
	slice[0] = "vasya"
}
//main не видел изменения в слайсе через append
//безопаснее сделать через стандартный return


//сделать возврат значения через указатель