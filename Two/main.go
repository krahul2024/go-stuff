package main

import "fmt"

func main() {
	car1 := Car{
		Make:       "Toyota",
		Model:      "Corolla",
		Year:       2022,
		Price:      25000,
		Horsepower: 140,
	}

	bike1 := Bike{
		Make:         "Honda",
		Model:        "CBR600RR",
		Year:         2021,
		Price:        12000,
		EngineSizeCC: 600,
		IsElectric:   false,
	}

	bus1 := Bus{
		Make:              "Mercedes-Benz",
		Model:             "Sprinter",
		Year:              2020,
		Price:             60000,
		PassengerCapacity: 12,
		FuelType:          "Diesel",
	}

	vehicles := []Vehicle{car1, bike1, bus1}
	for i := range vehicles {
		vehicles[i].details()
	}
}

// interface for common actions on a vehicle struct
type Vehicle interface {
	make()
	price() float32
	details()
}

type Car struct {
	Make       string
	Model      string
	Year       int
	Price      float32
	Horsepower int
}

func (car Car) make() {
	fmt.Println(car.Make)
}

func (car Car) price() float32 {
	return car.Price
}

func (car Car) details() {
	fmt.Printf("Car: %s %s, Year: %d, Price: $%.2f, Horsepower: %d\n", car.Make, car.Model, car.Year, car.price(), car.Horsepower)
}

type Bike struct {
	Make         string
	Model        string
	Year         int
	Price        float32
	EngineSizeCC int
	IsElectric   bool
}

func (bike Bike) make() {
	fmt.Println(bike.Make)
}

func (bike Bike) price() float32 {
	return bike.Price
}

func (bike Bike) details() {
	fmt.Printf("Bike: %s %s, Year: %d, Price: $%.2f, Engine Size: %dcc, Electric: %t\n", bike.Make, bike.Model, bike.Year, bike.price(), bike.EngineSizeCC, bike.IsElectric)
}

type Bus struct {
	Make              string
	Model             string
	Year              int
	Price             float32
	PassengerCapacity int
	FuelType          string
}

func (bus Bus) make() {
	fmt.Println(bus.Make)
}

func (bus Bus) price() float32 {
	return bus.Price
}

func (bus Bus) details() {
	fmt.Printf("Bus: %s %s, Year: %d, Price: $%.2f, Passenger Capacity: %d, Fuel Type: %s\n", bus.Make, bus.Model, bus.Year, bus.price(), bus.PassengerCapacity, bus.FuelType)
}
