# Задание: Анализ кода на Go

Это задание направлено на глубокое понимание работы срезов (interface), их модификации и передачи в функциях Go.  
**Ваша задача:** Определить вывод каждой из предложенных программ и зафиксировать ответы **в сообщениях коммитов** с кратким объяснением логики.

1.
```go
package main  
  
import (  
    "fmt"  
)  
  
type MyError struct {  
    data string  
}  
  
func (m *MyError) Error() string {  
    return m.data  
}  
func foo(i int) error {  
    var err *MyError  
    if i > 5 {  
       err = &MyError{data: "i>5"}  
    }  
    return err  
}  
func main() {  
    err := foo(4)  
    if err != nil {  
       fmt.Println("oops")  
    } else {  
       fmt.Println("ok")  
    }  
}
//oops
//////////////////////////////////////////////////////
2.
package main  
  
import (  
    "fmt"  
)  
  
type errorString struct {  
    s string  
}  
  
func (e errorString) Error() string {  
    return e.s  
}  
  
func checkErr(err error) {  
    fmt.Println(err == nil)  
}  
  
func main() {  
    var e1 error  
    checkErr(e1)  
    //true

    var e *errorString  
    checkErr(e)  
    //false

    e = &errorString{}  
    checkErr(e)  
    //false
    //создаёт структуру с нулевым значением 
    
    e = nil  
    checkErr(e)  
    //false
}
////////////////////////////////////////////////////
3.
package main

import "fmt"

type CustomError struct {
	message string
}

func (e *CustomError) Error() string {
	return e.message
}

func returnError(flag bool) error {
	if flag {
		return &CustomError{"Что-то пошло не так"}
	}
	var err *CustomError
	return err
}

func main() {
	err1 := returnError(true)
	err2 := returnError(false)

	fmt.Println("err1 == nil:", err1 == nil) //false
	fmt.Println("err2 == nil:", err2 == nil) //false
    //Когда интерфейс == nil? пока ему не присвоено или не объялено ни одно значение
}
