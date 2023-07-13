package main

import "fmt"

type Vehicle interface {
	Start()
	Stop()
	Accelerate()
	Brake()
}

type Car struct {
}

func (c Car) Start() {
	fmt.Println("Car started")
}

func (c Car) Stop() {
	fmt.Println("Car stopped")
}

func (c Car) Accelerate() {
	fmt.Println("Car accelerated")
}

func (c Car) Brake() {
	fmt.Println("Car braked")
}

type Motorcycle struct {
}

func (m Motorcycle) Start() {
	fmt.Println("Motorcycle started")
}

func (m Motorcycle) Stop() {
	fmt.Println("Motorcycle stopped")
}

func (m Motorcycle) Accelerate() {
	fmt.Println("Motorcycle accelerated")
}

func (m Motorcycle) Brake() {
	fmt.Println("Motorcycle braked")
}

func PerformDrive(vehicle Vehicle) {
	vehicle.Start()
	vehicle.Accelerate()
	vehicle.Brake()
	vehicle.Stop()
}

func main() {
	car := Car{}
	motorcycle := Motorcycle{}

	PerformDrive(car)
	PerformDrive(motorcycle)
}
