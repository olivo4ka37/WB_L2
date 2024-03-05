package main

// Square тип данных описывающий характеристики квадрата
type Square struct {
	side int
}

// accept функция реализующая посетителя для типа Square
func (s *Square) accept(v Visitor) {
	v.visitForSquare(s)
}

// getType функция вывода типа данных Square
func (s *Square) getType() string {
	return "Square"
}
