package main

import (
	"fmt"
	"sort"
	"strings"
)

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
	keys := make([]string, 0, len(freqMap))
	for word := range freqMap {
		keys = append(keys, word)
	}

	//Внутри функции обрабатывается правило (func(i,j int) bool{}).
	// Также внутри есть умение сортировать.
	//
	sort.Slice(keys, func(i, j int) bool {
		return freqMap[keys[i]] > freqMap[keys[j]]
	})

	//Проходимся по отсортированному срезу и печатаем слова из карты
	for _, word := range keys {
		fmt.Printf("%s: %d\n", word, freqMap[word])
	}
}

func main() {
	freqMap := make(map[string]int)
	text := "Golang is great and and and and golang is fast fast fast"
	freqMap = WordFrequency(text)
	PrintWordFrequency(freqMap)
}


//Привести к одному регистру (нижний)