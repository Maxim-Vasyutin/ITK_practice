package main

import (
	"encoding/json"
	"fmt"
	"sync"
	"time"
)

type User struct {
	Name string
}

type item struct {
	valueItem interface{}
	dieTime   time.Time
}

//В valueItem лежит интерфейс, чтобы принимать любое значение (any)
//В dieTime время смерти (время создания + время жизни)

type Cache struct {
	v  map[string]item
	mu sync.RWMutex
}

//Мьютекс ставиться над полями, которые защищает

func NewCache() *Cache {
	return &Cache{v: make(map[string]item)}
}

// База
// Добавляет значение с указанным TTL
func (c *Cache) Set(key string, value interface{}, ttl time.Duration) {
	//ttl - промежуток
	c.mu.Lock()
	defer c.mu.Unlock()
	//момент смерти
	c.v[key] = item{
		valueItem: value,
		dieTime:   time.Now().Add(ttl),
	}
}

// Возвращает значение (с проверкой TTL)
func (c *Cache) Get(key string) (interface{}, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	//проверка, есть ли значение в мапе по ключу
	if it, ok := c.v[key]; ok {
		//Время смерти раньше, чем сейчас?
		//То есть, запись протухла?
		if it.dieTime.Before(time.Now()) {
			return nil, false
		}
		//если нет, то есть, запись НЕ протухла - отправляем значение
		return it.valueItem, true
	}
	return nil, false
}

// Удаляет конкретный ключ
func (c *Cache) Delete(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	delete(c.v, key)
}

// Проверяет наличие непросроченного ключа
func (c *Cache) Exists(key string) bool {
	c.mu.RLock()
	defer c.mu.RUnlock()

	if it, ok := c.v[key]; ok {
		if it.dieTime.After(time.Now()) {
			return true
		}
	}

	return false
}

// Расширенные операции
// Полностью очищает кэш
func (c *Cache) Clear() {
	//Так как поля приватные и внешних ссылок нет, можно удалить полностью
	//(присвоить новую чистую мапу)
	c.v = map[string]item{}
}

// Сериализует данные в JSON
func (c *Cache) ToJSON() ([]byte, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	vereficationMap := make(map[string]interface{})

	for key, valueLocal := range c.v {
		if valueLocal.dieTime.After(time.Now()) {
			vereficationMap[key] = valueLocal.valueItem
		}
	}
	return json.Marshal(vereficationMap)
}

// Типизированное получение
func GetAs[T any](c *Cache, key string) (T, error) {
	// методы НЕ могут иметь собственных параметров типа (дженериков)
	// Поэтому передаём кэш обычным аргументом.

	//Тут без блокировок, потому что дальше ресурс будет блокироваться в Get()
	var zero T

	inter, ok := c.Get(key)

	//если значение отсутствует после проверки
	if !ok {
		return zero, fmt.Errorf("значения нет")
	} else {
		//если есть, переводим в тип T
		val, ok := inter.(T)
		//если получилось, то всё ок
		if ok {
			return val, nil
		} else {
			return zero, fmt.Errorf("несовпадение типа")
		}
	}
}

func main() {
	cache := NewCache()

	// --- Базовые операции ---
	cache.Set("user", User{Name: "Alice"}, time.Hour)
	cache.Set("temp_data", 42, time.Minute)

	v, ok := cache.Get("user")
	fmt.Println("Get(user):", v, ok) // {Alice} true

	v, ok = cache.Get("temp_data")
	fmt.Println("Get(temp_data):", v, ok) // 42 true

	v, ok = cache.Get("missing")
	fmt.Println("Get(missing):", v, ok) // <nil> false

	fmt.Println("Exists(user):", cache.Exists("user"))       // true
	fmt.Println("Exists(missing):", cache.Exists("missing")) // false

	// --- Типизированное получение ---
	u, err := GetAs[User](cache, "user")
	fmt.Println("GetAs[User](user):", u, err) // {Alice} <nil>
	fmt.Println("  u.Name напрямую:", u.Name) // Alice  <- уже типизировано!

	n, err := GetAs[int](cache, "temp_data")
	fmt.Println("GetAs[int](temp_data):", n, err) // 42 <nil>

	// промах по типу: в "user" лежит User, просим int
	bad, err := GetAs[int](cache, "user")
	fmt.Println("GetAs[int](user):", bad, err) // 0 несовпадение типа

	// несуществующий ключ
	miss, err := GetAs[User](cache, "missing")
	fmt.Println("GetAs[User](missing):", miss, err) // {} значения нет

	// --- JSON ---
	b, err := cache.ToJSON()
	fmt.Println("ToJSON:", string(b), err)

	// --- TTL: проверяем протухание ---
	cache.Set("fast", "я умру быстро", 100*time.Millisecond)
	fmt.Println("fast сразу:", cache.Exists("fast")) // true

	time.Sleep(200 * time.Millisecond)

	v, ok = cache.Get("fast")
	fmt.Println("fast после сна:", v, ok)                        // <nil> false
	fmt.Println("Exists(fast) после сна:", cache.Exists("fast")) // false

	// протухший ключ не должен попасть в JSON
	b, _ = cache.ToJSON()
	fmt.Println("ToJSON после протухания:", string(b))

	// --- Delete ---
	cache.Delete("user")
	fmt.Println("Exists(user) после Delete:", cache.Exists("user")) // false

	// --- Clear ---
	cache.Clear()
	b, _ = cache.ToJSON()
	fmt.Println("ToJSON после Clear:", string(b)) // {}
}
