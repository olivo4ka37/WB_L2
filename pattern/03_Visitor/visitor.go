package main

// Visitor интерфейс паттерна Visitor для разных типов данных
type Visitor interface {
	visitForSquare(*Square)
	visitForCircle(*Circle)
	visitForrectangle(*Rectangle)
}
