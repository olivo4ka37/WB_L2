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
- в ЕТ++ паттерн цепочка обязанностей применяется для обработки запросов на обновление графического изображения.
*/

import "fmt"

type Handler interface {
	Handle(byte)
}

type ConcreteHandlerA struct {
	next Handler
}

func (h ConcreteHandlerA) Handle(msg byte) {
	if msg == 'A' {
		fmt.Println("The request was handled by ConcreteHandlerA")
	} else if h.next != nil {
		h.next.Handle(msg)
	}
}

type ConcreteHandlerB struct {
	next Handler
}

func (h ConcreteHandlerB) Handle(msg byte) {
	if msg == 'B' {
		fmt.Println("The request was handled by ConcreteHandlerB")
	} else if h.next != nil {
		h.next.Handle(msg)
	}
}

type ConcreteHandlerC struct {
	next Handler
}

func (h ConcreteHandlerC) Handle(msg byte) {
	if msg == 'C' {
		fmt.Println("The request was handled by ConcreteHandlerC")
	} else if h.next != nil {
		h.next.Handle(msg)
	}
}

func main() {
	handler := ConcreteHandlerA{
		next: ConcreteHandlerB{
			next: ConcreteHandlerC{},
		},
	}

	handler.Handle('C')
}
