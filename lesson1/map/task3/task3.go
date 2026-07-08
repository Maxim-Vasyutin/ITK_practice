package main

import (
	"fmt"
)

// FilterByValue возвращает новую map, содержащую только элементы,
// значения которых присутствуют в allowedValues.
func FilterByValue(m map[int]string, allowedValues []string) map[int]string {
	// Преобразовать allowedValues в set для быстрой проверки
	// Создать новую map и заполнить её подходящими элементами
	allowedSet := make(map[string]struct{})
	local_map := make(map[int]string)

	for _, word := range allowedValues {
		allowedSet[word] = struct{}{}
	}

	for i, word := range m {
		if _, ok := allowedSet[word]; ok {
			local_map[i] = word
		}
	}
	return local_map
}

// InvertMap меняет ключи и значения местами.
// Если значения исходной map не уникальны, возвращает ошибку.
func InvertMap(m map[string]int) (map[int]string, error) {
	// Проверять уникальность значений
	// При обнаружении дубликата вернуть ошибку с описанием конфликта
	invertMap := make(map[int]string)

	for word, num := range m {
		if _, ok := invertMap[num]; ok {
			err := fmt.Errorf("Мапа уже имеет значение: %s", word)
			return nil, err
		}

		invertMap[num] = word

	}
	return invertMap, nil
}

func main() {
	// ==========================================
	// 1. ТЕСТ ДЛЯ ФУНКЦИИ FilterByValue
	// ==========================================
	fmt.Println("--- Тест FilterByValue ---")

	sourceMap := map[int]string{
		1: "apple",
		2: "banana",
		3: "orange",
		4: "pear",
	}
	allowed := []string{"apple", "orange", "grape"}

	// Вызываем вашу функцию фильтрации
	filteredResult := FilterByValue(sourceMap, allowed)

	// Ожидаем увидеть только apple (1) и orange (3)
	fmt.Printf("Исходная мапа: %v\n", sourceMap)
	fmt.Printf("Разрешенные слова: %v\n", allowed)
	fmt.Printf("Результат фильтрации: %v\n\n", filteredResult)

	// ==========================================
	// 2. ТЕСТ ДЛЯ ФУНКЦИИ InvertMap (Успешный случай)
	// ==========================================
	fmt.Println("--- Тест InvertMap (Успешный) ---")

	validMap := map[string]int{
		"Алексей": 1,
		"Мария":   2,
		"Иван":    3,
	}

	invertedResult, err := InvertMap(validMap)
	if err != nil {
		fmt.Printf("Ошибка: %v\n", err)
	} else {
		// Ожидаем увидеть мапу, где ключи — числа, а значения — имена
		fmt.Printf("Исходная мапа: %v\n", validMap)
		fmt.Printf("Инвертированная мапа: %v\n\n", invertedResult)
	}

	// ==========================================
	// 3. ТЕСТ ДЛЯ ФУНКЦИИ InvertMap (С дубликатом)
	// ==========================================
	fmt.Println("--- Тест InvertMap (С дубликатом) ---")

	badMap := map[string]int{
		"Алексей":  1,
		"Мария":    2,
		"Дубликат": 1, // Число 1 повторяется! В вашей функции num станет ключом invertMap
	}

	failedResult, err := InvertMap(badMap)
	if err != nil {
		// Ожидаем, что функция поймает дубликат и вернет созданную вами ошибку
		fmt.Println("Успешно поймали ошибку дубликата!")
		fmt.Printf("Текст ошибки: %v\n", err)
		fmt.Printf("Результат (должен быть nil): %v\n", failedResult)
	} else {
		fmt.Printf("Ошибка не сработала, результат: %v\n", failedResult)
	}
}


//Нейминг поправить (без)