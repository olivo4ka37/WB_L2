package main

import "fmt"

// MiddleCoordinates структура данных центральной точки координат фигуры
type MiddleCoordinates struct {
	x int
	y int
}

// visitForSquare Считает координаты центральной точкий для типа данных Square
func (a *MiddleCoordinates) visitForSquare(s *Square) {
	// Calculate middle point coordinates for square.
	// Then assign in to the x and y instance variable.
	fmt.Println("Calculating middle point coordinates for square")
}

// visitForCircle Считает координаты центральной точкий для типа данных Circle
func (a *MiddleCoordinates) visitForCircle(c *Circle) {
	fmt.Println("Calculating middle point coordinates for circle")
}

// visitForrectangle Считает координаты центральной точкий для типа данных Rectangle
func (a *MiddleCoordinates) visitForrectangle(t *Rectangle) {
	fmt.Println("Calculating middle point coordinates for rectangle")
}
