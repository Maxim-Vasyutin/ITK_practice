package main

import (
	"fmt"
	"sync"
)

var m map[string]int

func init() {
	// Инициализируйте map
	m = make(map[string]int)
}

// Добавление записей
func AddPerson(name string, age int) {
	// Реализуйте добавление записи
	m[name] = age
}

// Получение возраста
func GetAge(name string) (int, bool) {
	// Реализуйте получение возраста
	if _, ok := m[name]; ok {

		return m[name], ok
	}
	return m[name], false
}

// Удаление записи
func DeletePerson(name string) {
	// Реализуйте удаление
	var mu sync.Mutex

	mu.Lock()
	defer mu.Unlock()

	delete(m, name)
}

// Вывод всех записей
func PrintAll() {
	// Реализуйте вывод всех записей
	for name, age := range m {
		fmt.Printf("%s: %d\n", name, age)
	}
}

func main() {
	AddPerson("Anna", 30)
	AddPerson("Boris", 45)
	AddPerson("Clara", 28)

	fmt.Println("Все записи:")
	PrintAll()

	age, ok := GetAge("Anna")
	if age == 0 && ok {
		fmt.Println("Возраст Anna больше 0")
	}
	fmt.Println("Возраст Anna:", age)

	DeletePerson("Boris")

	fmt.Println("После удаления Boris:")
	PrintAll()
}

//Можно возвращать не только int но и bool из GetAge
//Добавить идиому _,ok

//Добавление удаление записей - добавить мьютекс. гонка данных
