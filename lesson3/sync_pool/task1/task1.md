## Оптимизация обработки строк с sync.Pool

### Описание задачи
В высоконагруженном сервисе частые аллокации буферов для преобразования строк создают нагрузку на GC.
Цель — реализовать оптимизированную функцию `ProcessString` с использованием `sync.Pool`, чтобы переиспользовать буферы `[]byte`.

### Требования
1. Функция `ProcessString(s string) string` преобразует строку в верхний регистр.
2. Использование `sync.Pool` для буферов `[]byte`.
3. Потокобезопасность, отсутствие утечек памяти.
```go
func main() {
	examples := []string{
		"hello, world!",
		"gopher",
		"lorem ipsum dolor sit amet",
	}

	for _, s := range examples {
		processed := ProcessString(s)
		fmt.Printf("Original: %q\nProcessed: %q\n\n", s, processed)
	}
}
```