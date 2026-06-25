package main

import "fmt"

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
func GetAge(name string) int {
	// Реализуйте получение возраста
	return m[name]
}

// Удаление записи
func DeletePerson(name string) {
	// Реализуйте удаление
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

	fmt.Println("Возраст Anna:", GetAge("Anna"))

	DeletePerson("Boris")

	fmt.Println("После удаления Boris:")
	PrintAll()
}
