package main

import (
	"fmt"
)

// AreaCalculator Калькулятор площади фигуры
type AreaCalculator struct {
	area int
}

// visitForSquare считает площадь квадрата
func (a *AreaCalculator) visitForSquare(s *Square) {
	// Calculate area for square.
	// Then assign in to the area instance variable.
	fmt.Println("Calculating area for square")
}

// visitForCircle считает площадб круга
func (a *AreaCalculator) visitForCircle(s *Circle) {
	fmt.Println("Calculating area for circle")
}

// visitForrectangle считает площадь прямоугольника
func (a *AreaCalculator) visitForrectangle(s *Rectangle) {
	fmt.Println("Calculating area for rectangle")
}
