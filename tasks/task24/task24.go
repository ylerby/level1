package main

import (
	"fmt"
	"math"
)

/*
Разработать программу нахождения расстояния между двумя точками,
которые представлены в виде структуры Point
с инкапсулированными параметрами x,y и конструктором.
*/

// Point - структура содержащая координаты точки типа float64
type Point struct {
	x,
	y float64
}

// New - "конструктор" для типа Point. Создает новый объект типа Point
func New(newX, newY float64) Point {
	return Point{
		x: newX,
		y: newY,
	}
}

// distance - метод структуры Point, принимает объект типа Point (вторая точка),
// возвращает расстояние между точками
func (p Point) distance(secondPoint Point) float64 {
	xDifference := p.x - secondPoint.x
	yDifference := p.y - secondPoint.y

	// расстояние между точками вычисляется как √((x1 - x2)ˆ2 + (y1 - y2)ˆ2)
	distance := math.Sqrt(math.Pow(xDifference, 2) + math.Pow(yDifference, 2))
	return distance
}

func main() {
	firstPoint := New(14.4, 64.2)
	secondPoint := New(-54.4, 194.2)

	result := firstPoint.distance(secondPoint)
	fmt.Printf("Расстояние между точками: %f\n", result)
}
