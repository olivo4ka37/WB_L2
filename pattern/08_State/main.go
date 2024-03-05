package main

import "fmt"

/*
	Реализовать паттерн «состояние».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/State_pattern
*/

/*
Состояние - паттерн поведения объектов.
Он позволяет объекту варьировать свое поведение в зависимости от внутреннего состояния.
Извне создается впечатление, что изменился класс объекта.

Применимость паттерна:
- когда поведение объекта зависит от его состояния и должно изменяться во время выполнения;
- когда в коде операций встречаются состоящие из многих ветвей условные операторы,
в которых выбор ветви зависит от состояния.

Плюсы:
- локализует зависящее от состояния поведение и делит его на части, соответствующие состояниям;
- делает явными переходы между состояниями;
- объекты состояния можно разделять.

Примеры использования:
- В графических редакторах позволяет клиентам легко определять новые виды инструментов.
*/

// FreezingMode интерфейс реализующий режим заморозки в холодильнике
type FreezingMode interface {
	Freeze() string
}

// Fridge структура имитирующая холодильник
type Fridge struct {
	state FreezingMode
}

// Freeze функция структуры Fridge возвращающая текущее состояние заморозки
func (a *Fridge) Freeze() string {
	return a.state.Freeze()
}

// SetState установка состояния заморозки в холодильник
func (a *Fridge) SetState(state FreezingMode) {
	a.state = state
}

// NewFridge конструктор холодильника, возвращает тип холодильник и устанавливает значение обычной заморозки
func NewFridge() *Fridge {
	return &Fridge{state: &DefaultFreezingMode{}}
}

// DefaultFreezingMode структура состояния обычной заморозки
type DefaultFreezingMode struct {
}

// Freeze функция структуры DefaultFreezingMode реализующая обычный режим заморозки
func (a *DefaultFreezingMode) Freeze() string {
	return "Freeze"
}

// SuperFreezingMode структура состояния супер заморозки
type SuperFreezingMode struct {
}

// Freeze функция структуры SuperFreezingMode реализующая супер режим заморозки
func (a *SuperFreezingMode) Freeze() string {
	return "Super freeze"
}

func main() {
	// Инициализация холодильника через конструктор
	fridge := NewFridge()

	// Проверка работы переключений состояний
	fmt.Println(fridge.Freeze())
	fridge.SetState(&SuperFreezingMode{})
	fmt.Println(fridge.Freeze())
}
