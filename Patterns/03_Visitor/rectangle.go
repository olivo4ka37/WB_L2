package main

// Rectangle структура данных описывающая характеристики прямоугольника
type Rectangle struct {
	l int
	b int
}

// accept функция реализующая посетителя для типа Rectangle
func (t *Rectangle) accept(v Visitor) {
	v.visitForrectangle(t)
}

// getType функция вывода типа данных Rectangle
func (t *Rectangle) getType() string {
	return "rectangle"
}
