package main

/*
	Реализовать паттерн «цепочка вызовов».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Chain-of-responsibility_pattern
*/

/*
Цепочка обязанностей - паттерн поведения объектов.
Он позволяет избежать привязки отправителя запроса к его получателю, давая шанс обработать запрос нескольким объектам.
Связывает объекты-получатели в цепочку и передает запрос вдоль этой цепочки, пока его не обработают.

Применимость паттерна:
- когда есть более одного объекта, способного обработать запрос,
причем настоящий обработчик заранее неизвестен и должен быть найден автоматически;
- когда необходимо отправить запрос одному из нескольких объектов, не указывая явно, какому именно;
- когда есть набор объектов, способных обработать запрос, должен задаваться динамически.

Плюсы и минусы:
- ослабление связанности;
- дополнительная гибкость при распределении обязанностей между объектами;
- получение не гарантировано.

Примеры использования:
- в веб приложении цепочка обязанностей может быть использована, для обработки заказа в интернет маркет-плейсе проверка наличия товара на складе,
проверка статуса платежа, обработка доставки и тд.
*/

import "fmt"

// Handler интерфейс обработчиков
type Handler interface {
	Handle(byte)
}

// ConcreteHandlerA структура содержащая некий обработчик A
type ConcreteHandlerA struct {
	next Handler
}

// Handle обработчик структуры A
func (h ConcreteHandlerA) Handle(msg byte) {
	if msg == 'A' {
		fmt.Println("The request was handled by ConcreteHandlerA")
	} else if h.next != nil {
		h.next.Handle(msg)
	}
}

// ConcreteHandlerB структура содержащая некий обработчик B
type ConcreteHandlerB struct {
	next Handler
}

// Handle обработчик структуры B
func (h ConcreteHandlerB) Handle(msg byte) {
	if msg == 'B' {
		fmt.Println("The request was handled by ConcreteHandlerB")
	} else if h.next != nil {
		h.next.Handle(msg)
	}
}

// ConcreteHandlerC структура содержащая некий обработчик C
type ConcreteHandlerC struct {
	next Handler
}

// Handle обработчик структуры 	C
func (h ConcreteHandlerC) Handle(msg byte) {
	if msg == 'C' {
		fmt.Println("The request was handled by ConcreteHandlerC")
	} else if h.next != nil {
		h.next.Handle(msg)
	}
}

func main() {
	// Инициализация интерфейса в котором реализованна цепочка обязанностей
	handler := ConcreteHandlerA{
		next: ConcreteHandlerB{
			next: ConcreteHandlerC{},
		},
	}

	handler.Handle('A')
	handler.Handle('B')
	handler.Handle('C')
}
