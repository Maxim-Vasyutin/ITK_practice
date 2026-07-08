package main

import (
	"errors"
	"fmt"
)

var (
	ErrEngineAlreadyRunning = errors.New("двигатель уже работает")
	ErrEngineOff            = errors.New("двигатель не запущен")
	ErrLowBattery           = errors.New("низкий заряд батареи")
)

type Vehicle interface {
	StartEngine() error
	StopEngine() error
	GetInfo() string
}

// ////////////////////////////////
type Car struct {
	brand    string
	engineOn bool
}

func (c Car) Honk() string {
	return "Beep beep!"
}

func (c *Car) StartEngine() error {
	if c.engineOn {
		return ErrEngineAlreadyRunning
	}
	c.engineOn = true
	return nil
}

func (c *Car) StopEngine() error {
	if !c.engineOn {
		return ErrEngineOff
	}
	c.engineOn = false
	return nil
}

func (c *Car) GetInfo() string {
	return fmt.Sprintf("Машина %s\nДвигатель запущен: %v", c.brand, c.engineOn)
}

func (c Car) GetEngineStatus() bool {
	return c.engineOn
}

// //////////////////////////////
type Truck struct {
	Car
	cargoCapacity float64
}

func (t Truck) Honk() string {
	return "Honk Honk!"
}

func (t *Truck) StartEngine() error {
	//если правда, что он запущен, то возвращаем ошибку запуска
	if t.engineOn {
		return ErrEngineAlreadyRunning
	}
	t.engineOn = true
	return nil
}

func (t *Truck) StopEngine() error {

	if !t.engineOn {
		return ErrEngineOff
	}
	t.engineOn = false
	return nil
}
func (t *Truck) GetInfo() string {
	return fmt.Sprintf("Грузовик %s\nДвигатель запущен: %v"+
		"\nГрузоподъёмность: %f", t.brand, t.engineOn, t.cargoCapacity)
}

func (t Truck) GetEngineStatus() bool {
	return t.engineOn
}

func (t Truck) GetCargoCapacity() float64 {
	return t.cargoCapacity
}

// //////////////////////////////////
type ElectricCar struct {
	Car
	batteryLevel int
}

func (ec *ElectricCar) StartEngine() error {
	if ec.engineOn {
		return ErrEngineAlreadyRunning
	}

	if ec.batteryLevel < 5 {
		return ErrLowBattery
	}
	ec.engineOn = true
	return nil
}

func (ec *ElectricCar) StopEngine() error {
	if !ec.engineOn {
		return ErrEngineOff
	}
	ec.engineOn = false
	return nil
}

func (ec *ElectricCar) GetInfo() string {
	return fmt.Sprintf("Машина %s\nДвигатель запущен: %v"+
		"\nУровень заряда батареи: %d", ec.brand, ec.engineOn, ec.batteryLevel)
}

func (ec ElectricCar) GetEngineStatus() bool {
	return ec.engineOn
}
func (ec ElectricCar) GetBatteryLevel() int {
	return ec.batteryLevel
}

//////////////////////////

//сделать встраиваение, а не дублирование (методы start stop)