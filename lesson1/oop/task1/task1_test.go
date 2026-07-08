package main

import (
	"errors"
	"testing"
)

func TestCar_StartStopEngine(t *testing.T) {
	car := Car{brand: "BMW"}

	err := car.StartEngine()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if !car.GetEngineStatus() {
		t.Fatal("engine should be ON")
	}

	err = car.StopEngine()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if car.GetEngineStatus() {
		t.Fatal("engine should be OFF")
	}
}

func TestCar_InvalidTransitions(t *testing.T) {
	car := Car{brand: "BMW"}

	err := car.StopEngine()
	if !errors.Is(err, ErrEngineOff) {
		t.Fatalf("expected ErrEngineOff, got %v", err)
	}

	_ = car.StartEngine()
	err = car.StartEngine()

	if !errors.Is(err, ErrEngineAlreadyRunning) {
		t.Fatalf("expected ErrEngineAlreadyRunning, got %v", err)
	}
}

func TestTruck_Behavior(t *testing.T) {
	truck := Truck{
		Car:           Car{brand: "Volvo"},
		cargoCapacity: 10.5,
	}

	err := truck.StartEngine()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if !truck.GetEngineStatus() {
		t.Fatal("truck engine should be ON")
	}

	if truck.Honk() != "Honk Honk!" {
		t.Fatal("unexpected honk sound")
	}

	if truck.GetCargoCapacity() != 10.5 {
		t.Fatal("wrong cargo capacity")
	}
}

func TestElectricCar_BatteryConstraint(t *testing.T) {
	ec := ElectricCar{
		Car:          Car{brand: "Tesla"},
		batteryLevel: 3,
	}

	err := ec.StartEngine()

	if !errors.Is(err, ErrLowBattery) {
		t.Fatalf("expected ErrLowBattery, got %v", err)
	}

	if ec.GetEngineStatus() {
		t.Fatal("engine should NOT start with low battery")
	}
}

func TestElectricCar_BoundaryBattery(t *testing.T) {
	ec := ElectricCar{
		Car:          Car{brand: "Tesla"},
		batteryLevel: 5,
	}

	err := ec.StartEngine()

	if err != nil {
		t.Fatalf("battery=5 should be valid, got error: %v", err)
	}

	if !ec.GetEngineStatus() {
		t.Fatal("engine should be ON")
	}
}

func TestElectricCar_StartStop(t *testing.T) {
	ec := ElectricCar{
		Car:          Car{brand: "Tesla"},
		batteryLevel: 80,
	}

	err := ec.StartEngine()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if !ec.GetEngineStatus() {
		t.Fatal("engine should be ON")
	}

	err = ec.StopEngine()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if ec.GetEngineStatus() {
		t.Fatal("engine should be OFF")
	}
}

func TestVehicle_Polymorphism(t *testing.T) {
	vehicles := []Vehicle{
		&Car{brand: "BMW"},
		&Truck{Car: Car{brand: "Volvo"}, cargoCapacity: 12},
		&ElectricCar{Car: Car{brand: "Tesla"}, batteryLevel: 80},
	}

	for _, v := range vehicles {
		err := v.StartEngine()
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if v.GetInfo() == "" {
			t.Fatal("GetInfo should not be empty")
		}
	}
}
