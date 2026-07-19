## JSON-кэш с `sync.Pool` и `map + RWMutex`

### **Описание**
Этот проект демонстрирует **потокобезопасный JSON-кэш** с поддержкой TTL и оптимизированной сериализацией.  
Используется `sync.Pool` для **эффективной работы с JSON**, а также `map + RWMutex` для **более быстрого доступа к данным**.

---

### **Основные возможности**
- **Хранение объектов в `map` (с TTL)**
- **Автоматическое удаление устаревших объектов**
- **Быстрая сериализация JSON с `sync.Pool`**
- **Использование `sync.RWMutex` для конкурентного доступа**

---

### **Методы**
#### **Базовые операции**
- `Set(key string, value interface{})` – **добавить объект в кэш**
- `Get(key string) (interface{}, bool)` – **получить объект по ключу**
- `Delete(key string)` – **удалить объект**
- `ToJSON() ([]byte, error)` – **сериализовать кэш в JSON**

---

### **Как это работает?**
- Все объекты хранятся в **`map[string]item`** (ключ → объект с TTL).
- `sync.Pool` позволяет **переиспользовать JSON-буферы**, снижая нагрузку на GC.
- Очистка устаревших данных выполняется **в отдельной горутине**.  

```go
func main() {
    cache := NewObjectCache(5 * time.Second)

    // Добавляем данные в кэш
    cache.Set("user:1", map[string]string{"name": "Alice", "role": "admin"})
    cache.Set("user:2", map[string]string{"name": "Bob", "role": "user"})

    // Получаем объект
    if user, found := cache.Get("user:1"); found {
        fmt.Println("Найден:", user)
    }

    // Выводим JSON
    jsonData, _ := cache.ToJSON()
    fmt.Println("Кэш в JSON:", string(jsonData))

    // Ждём истечения TTL и проверяем снова
    time.Sleep(6 * time.Second)
    _, found := cache.Get("user:1")
    fmt.Println("После TTL, user:1 найден?", found)
}
```
