```go

1.# Задание: Анализ кода на Go  
   
**Ваша задача:** Определить вывод программы и зафиксировать ответы **в сообщениях коммитов** с кратким объяснением логики.

package main  
  
import (  
    "fmt"  
)  
  
func main() {  
    fmt.Println("start")  
    for i := 1; i < 4; i++ {  
       defer fmt.Println(i)  
    }  
    fmt.Println("end")  
}

//Вызов будет копиться до конца родительской функции, а потом вызывается в обратном порядке
/*Вывод:
start
end
3
2
1
*/
////////////////////////////////////////////////////////////
2.
# Задание: Анализ кода на Go  
  
package main  
  
import "fmt"  
  
func main() {  
    value := 123  
    defer fmt.Println(value)  
    changeValue(&value)  
}  
func changeValue(value *int) {  
    *value = 456  
}

//defer добавился в очередь вызовов со старым значением
//Вывод:
//123
////////////////////////////////////////////////////////////

3.
package main

import (
	"errors"
	"fmt"
)

func main() {
    println("Case 1")
    case1()
    println()
    println()

    println("Case 2")
    case2()
    println()
    println()

    println("Case 3")
    case3()
    println()
    println()

}

func case1() {
    helperWithDefer := func(isError bool) error {
        var retVal error

        defer func() {
            retVal = errors.New("Extra error")
        }()

        if isError {
            retVal = errors.New("Default error")
        }

        return retVal
    }

    helperWithoutDefer := func(isError bool) error {
        var retVal error

        if isError {
            retVal = errors.New("Default error")
        }

        return retVal
    }

    fmt.Println("\twithout:")
    fmt.Println(helperWithoutDefer(false))
    fmt.Println(helperWithoutDefer(true))
    fmt.Println("\twith:")
    fmt.Println(helperWithDefer(false))
    fmt.Println(helperWithDefer(true))

    /*
    Вывод1:
        without
    false
    Defaullt error
        with
    Extra error
    Default error
    Extra error
    */

    //return сначала записывает результат в переменную возврата, потом уже выполняет defer
    //Там ещё что-то про возвращаемое значение. Оно не конкретное, мы возвращаем тип
    /*
    Вывод2:
    without
    nil
    Defaullt error
        with
    nil
    Defaullt error
    */
}


func case2() {
    helperWithDefer := func(isError bool) (retVal error) {
        defer func() {
            retVal = errors.New("Extra error")
        }()

        if isError {
            retVal = errors.New("Default error")
        }

        return
    }

    helperWithoutDefer := func(isError bool) (retVal error) {
        if isError {
            retVal = errors.New("Default error")
        }

        return
    }

    fmt.Println("\twithout:")
    fmt.Println(helperWithoutDefer(false))
    fmt.Println(helperWithoutDefer(true))
    fmt.Println("\twith:")
    fmt.Println(helperWithDefer(false))
    fmt.Println(helperWithDefer(true))


    /*
        without:
    nil
    Default error
        with:
    Extra error
    Extra error
    */

}


func case3() {
    helperWithDefer := func(isError bool) (retVal error) {
        defer func() {
            retVal = errors.New("First Error")
        }()

        defer func() {
            retVal = errors.New("Second Error")
        }()

        if isError {
            retVal = errors.New("Default error")
        }

        return
    }

    helperWithoutDefer := func(isError bool) (retVal error) {
        if isError {
            retVal = errors.New("Default error")
        }

        return
    }

    fmt.Println("\twithout:")
    fmt.Println(helperWithoutDefer(false))
    fmt.Println(helperWithoutDefer(true))
    fmt.Println("\twith:")
    fmt.Println(helperWithDefer(false))
    fmt.Println(helperWithDefer(true))

    /*
        without:
    nil
    Default error
        with:
    First Error
    First Error
    */
    //defer записывается в очередь LIFO
}