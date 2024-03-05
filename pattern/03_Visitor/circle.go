package main

// Circle структура данных описывающая характеристики круга
type Circle struct {
	radius int
}

// accept функция реализующая посетителя для типа Circle
func (c *Circle) accept(v Visitor) {
	v.visitForCircle(c)
}

// getType функция вывода типа данных Circle
func (c *Circle) getType() string {
	return "Circle"
}
