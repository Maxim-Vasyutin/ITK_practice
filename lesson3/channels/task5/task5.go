package main

import (
	"fmt"
	"os"
	"strings"
)

// Структура здесь для того, чтобы отправлять пачку информации
// о состоянии работы над файлом
type result struct {
	nameFile string
	count    int
	err      error
}

// file1 - 13
// file2 - 21
// file3 - 5
func main() {
	//Канал типа. То есть, будет передавать инфу про конкретный тип
	results := make(chan result)
	files := []string{"file1.txt", "file2.txt", "file3.txt"}

	//разделяет срез на пути
	for _, path := range files {
		go func(path string) {
			n, err := countWords(path)
			results <- result{
				nameFile: path,
				count:    n,
				err:      err,
			}
		}(path)
	}

	var total int
	for i := 0; i < len(files); i++ {
		r := <-results
		if r.err != nil {
			fmt.Printf("ошибка: %v при работе с файлом: %s\n", r.err, r.nameFile)
		} else {
			fmt.Printf("Всё нормально с файлом %s. Кол-во слов: %d\n", r.nameFile, r.count)
			total += r.count
		}
	}
	fmt.Println("\nВсего слов: ", total)
}

func countWords(path string) (int, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return 0, err
	}

	text := string(data)
	words := strings.Fields(text)
	count := len(words)

	return count, nil
}
