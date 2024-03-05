package main

// Shape интерфейс определения типа данных
type Shape interface {
	getType() string
	accept(Visitor)
}
