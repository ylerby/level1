package main

import "fmt"

// Driver представляет водителя, который заправляет автомобиль.
type Driver struct {
}

// InsertGasIntoCar заливает топливо в автомобиль.
func (d Driver) InsertGasIntoCar(c Car) {
	fmt.Println("Водитель заливает топливо в автомобиль")
	c.InsertGas()
}

// Car представляет интерфейс машины.
type Car interface {
	InsertGas()
}

// GasolineCar представляет машину на обычном бензине.
type GasolineCar struct {
}

// InsertGas метод "наполнения бензином" GasolineCar.
func (g *GasolineCar) InsertGas() {
	fmt.Println("Наполнение бензином GasolineCar")
}

// DieselCar представляет машину на дизельном бензине.
type DieselCar struct {
}

// insertDiesel метод "наполнения дизелем" DieselCar.
func (d *DieselCar) insertDiesel() {
	fmt.Println("Наполнение бензином DieselCar")
}

// DieselCarAdapter представляет адаптер для машины с DieselCar.
type DieselCarAdapter struct {
	DieselCarMachine *DieselCar
}

// InsertGas - метод "наполнения бензином" DieselCar.
func (d *DieselCarAdapter) InsertGas() {
	fmt.Println("Вместо обычного бензина заливаем дизельное топливо")
	d.DieselCarMachine.insertDiesel()
}

func main() {
	driver := &Driver{}
	gasolineCar := &GasolineCar{}

	driver.InsertGasIntoCar(gasolineCar)

	DieselCarMachine := &DieselCar{}
	DieselCarAdapter := &DieselCarAdapter{DieselCarMachine: DieselCarMachine}

	driver.InsertGasIntoCar(DieselCarAdapter)
}
