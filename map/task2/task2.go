package main

import "strings"

// WordFrequency принимает строку текста и возвращает map с частотой слов.
func WordFrequency(text string) map[string]int {
	// TODO: Реализуйте функцию.
	map_count := make(map[string]int)
	str := strings.Fields(text)

	for _, count := range str {
		map_count[count]++
	}
	return map_count
}

// PrintWordFrequency выводит частотный анализ слов, отсортированный по убыванию частоты.
func PrintWordFrequency(freqMap map[string]int) {
	// TODO: Реализуйте функцию.
	sort_map := make(map[string]int)
	for value := range freqMap {
		sort_map[value] = freqMap[value]
	}
}

func main() {

	text := "golang is great and golang is fast"

}
