package main

import (
	"fmt"
	"log"
)

/*
	Реализовать паттерн «фабричный метод».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Factory_method_pattern
*/

/*
Фабричный метод - паттерн, порождающий классы.
Он определяет интерфейс для создания объекта, но оставляет подклассам решение о том,
какой класс инстанцировать. Фабричный метод позволяет классу делегировать инстанцирование подклассам.

Применимость паттерна:
- когда классу заранее неизвестно, объекты каких классов ему нужно создавать;
- когда класс спроектирован так, чтобы объекты, которые он создает, специфицировались подклассами;
- когда класс делегирует свои обязанности одному из нескольких вспомогательных
подклассов, и вы планируете локализовать знание о том, какой класс принимает эти обязанности на себя.

Плюсы:
- предоставляет подклассам операции-зацепки (hooks);
- соединяет параллельные иерархии.

Примеры использования:
- Фабричные методы в изобилии встречаются в инструментальных библиотеках и каркасах.
*/

// action условные действия которые могут быть вызваны.
type action string

const (
	A action = "A"
	B action = "B"
	C action = "C"
)

// Creator обеспечивает интерфейс фабрики.
type Creator interface {
	CreateProduct(action string) Product // Factory Method
}

// Product обеспечивает интерфейс продукта.
// Все продукты созданные фабрикой должны возвращать один интерфейс.
type Product interface {
	Use() string
}

// ConcreteCreator реализует интерфейс креатора.
type ConcreteCreator struct{}

// NewCreator это ConcreteCreator конструктор.
func NewCreator() Creator {
	return &ConcreteCreator{}
}

// CreateProduct фабричный метод.
func (p *ConcreteCreator) CreateProduct(action string) Product {
	var product Product

	switch action {
	case `A`:
		product = &ConcreteProductA{string(action)}
	case `B`:
		product = &ConcreteProductB{string(action)}
	case `C`:
		product = &ConcreteProductC{string(action)}
	default:
		log.Fatalln("Неизвестное действие")
	}

	return product
}

// ConcreteProductA реализует продукт "A".
type ConcreteProductA struct {
	action string
}

// Use возвращает действие продукта.
func (p *ConcreteProductA) Use() string {
	return p.action
}

// ConcreteProductB реализует продукт "B".
type ConcreteProductB struct {
	action string
}

// Use возвращает действие продукта.
func (p *ConcreteProductB) Use() string {
	return p.action
}

// ConcreteProductC реализует продукт "C".
type ConcreteProductC struct {
	action string
}

// Use возвращает действие продукта.
func (p *ConcreteProductC) Use() string {
	return p.action
}

func main() {
	// Тестируем работу фабрики
	x := NewCreator()
	a := x.CreateProduct("A")
	fmt.Println(a)
	b := x.CreateProduct("B")
	fmt.Println(b)
	c := x.CreateProduct("C")
	fmt.Println(c)
	v := x.CreateProduct("V")
	fmt.Println(v)
}
