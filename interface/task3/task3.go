package main

import (
	"errors"
	"fmt"
	"strings"
)

var (
	ErrUnsupported = errors.New("обновление недоступно")
)

type Device interface {
	UpdateOS(string) error
	GetInfo() string
}

type Smartphone struct {
	OSVersion string
	Model     string
}

type Laptop struct {
	OSVersion string
	Model     string
}

type Smartwatch struct {
	OSVersion string
	Model     string
}

func (sp *Smartphone) UpdateOS(version string) error {
	if version >= "12.0" {
		return ErrUnsupported
	}
	sp.OSVersion = version
	return nil
}
func (lpt *Laptop) UpdateOS(version string) error {
	if !strings.HasPrefix(version, "Windows") {
		return ErrUnsupported
	}
	lpt.OSVersion = version
	return nil
}
func (sw *Smartwatch) UpdateOS(version string) error {
	if len(version) < 5 {
		return ErrUnsupported
	}
	sw.OSVersion = version
	return nil
}

func (sp *Smartphone) GetInfo() string {
	return fmt.Sprintf("Модель: %s, OC: %s", sp.Model, sp.OSVersion)
}
func (lpt *Laptop) GetInfo() string {
	return fmt.Sprintf("Модель: %s, OC: %s", lpt.Model, lpt.OSVersion)
}
func (sw *Smartwatch) GetInfo() string {
	return fmt.Sprintf("Модель: %s, OC: %s", sw.Model, sw.OSVersion)
}

//ИИ написал тест
//Ошибки были завязаны на неточностях и на невнимательном чтении задания
func main() {
	// Хелпер: пробуем обновить и печатаем результат
	tryUpdate := func(d Device, version string) {
		fmt.Printf("До:  %s\n", d.GetInfo())
		err := d.UpdateOS(version)
		if err != nil {
			fmt.Printf("Обновление на %q → ошибка: %v\n", version, err)
		} else {
			fmt.Printf("Обновление на %q → ок. После: %s\n", version, d.GetInfo())
		}
		fmt.Println("---")
	}

	fmt.Println("=== СМАРТФОН (нельзя, если текущая ОС >= 12.0) ===")
	// 11.0 младше 12.0 → обновление должно пройти
	tryUpdate(&Smartphone{Model: "iPhone", OSVersion: "11.0"}, "11.5")
	// 12.0 → ровно граница, обновлять нельзя
	tryUpdate(&Smartphone{Model: "iPhone", OSVersion: "12.0"}, "13.0")
	// ЛОВУШКА: 9.0 семантически МЛАДШЕ 12.0 → обновлять МОЖНО.
	// Но "9.0" >= "12.0" лексикографически = true. Если у тебя
	// наивное строковое сравнение — здесь будет ложная блокировка.
	tryUpdate(&Smartphone{Model: "iPhone", OSVersion: "9.0"}, "10.0")

	fmt.Println("=== НОУТБУК (только префикс Windows) ===")
	tryUpdate(&Laptop{Model: "ThinkPad", OSVersion: "Windows 10"}, "Windows 11")
	tryUpdate(&Laptop{Model: "ThinkPad", OSVersion: "Windows 10"}, "Linux Mint")
	// edge: регистр. HasPrefix чувствителен к регистру — "windows" ≠ "Windows"
	tryUpdate(&Laptop{Model: "ThinkPad", OSVersion: "Windows 10"}, "windows 11")

	fmt.Println("=== УМНЫЕ ЧАСЫ (нельзя, если новая версия короче 5 символов) ===")
	tryUpdate(&Smartwatch{Model: "Watch", OSVersion: "1.0.0"}, "1.0")    // 3 символа → нельзя
	tryUpdate(&Smartwatch{Model: "Watch", OSVersion: "1.0.0"}, "10.0.5") // 6 символов → можно
	tryUpdate(&Smartwatch{Model: "Watch", OSVersion: "1.0.0"}, "12.34")  // ровно 5 → "короче 5" не выполнено, можно

	fmt.Println("=== ПОЛИМОРФИЗМ: все устройства через интерфейс ===")
	devices := []Device{
		&Smartphone{Model: "Pixel", OSVersion: "10.0"},
		&Laptop{Model: "MacBook", OSVersion: "Windows 11"},
		&Smartwatch{Model: "Galaxy Watch", OSVersion: "5.0.0"},
	}
	for _, d := range devices {
		fmt.Println(d.GetInfo())
	}
}
