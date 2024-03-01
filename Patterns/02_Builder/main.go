package main

import (
	"fmt"
)

/*
	Реализовать паттерн «строитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Builder_pattern
*/

/*
Строитель - паттерн, порождающий объекты.
Он отделяет конструирование сложного объекта от его представления,
так что в результате одного и того же процесса конструирования могут получаться разные представления.

Применимость паттерна:
- алгоритм создания сложного объекта не должен зависеть от того,
из каких частей состоит объект и как они стыкуются между собой;
- процесс конструирования должен обеспечивать различные представления конструируемого объекта.

Плюсы:
- позволяет изменять внутреннее представление продукта;
- изолирует код, реализующий конструирование и представление;
- дает более тонкий контроль над процессом конструирования.

Минусы:

- Усложнение архитектуры кода при надобности управлении классов паттерна;

Вывод: Данный паттерн нужно реализовывать по мере его необходимости, для создания сложных объектов.

Примеры использования:
- приложение для конвертирования из формата RTF из библиотеки ЕТ++ [WGM88]
*/

// Pointers Тип содержащий две координаты типа float64, для графика осей x y
type Pointers struct {
	X float64
	Y float64
}

// PointersBuilderI Интерфейс для реализации 02_Builder паттерна для типа Pointers
type PointersBuilderI interface {
	SetX(value float64) PointersBuilderI
	SetY(value float64) PointersBuilderI

	Build() Pointers
}

// pointersBuilder Структура, которая будет в себе содержать инкапсулированную реализацию билдера для структуры Pointers
type pointersBuilder struct {
	x float64
	y float64
}

// SetX Устанавливает значение в координату x структуры pointersBuilder
func (p *pointersBuilder) SetX(value float64) PointersBuilderI {
	p.x = value

	return p
}

// SetY Устанавливает значение в координату y структуры pointersBuilder
func (p *pointersBuilder) SetY(value float64) PointersBuilderI {
	p.y = value

	return p
}

// Build Возвращает структуру Pointers со значениями x и y структуры pointersBuilder
func (p pointersBuilder) Build() Pointers {
	return Pointers{
		X: p.x,
		Y: p.y,
	}
}

// NewPointersBuilder создает новую структуру pointersBuilder
func NewPointersBuilder() pointersBuilder {
	return pointersBuilder{}
}

func main() {
	// Создаем новую переменную типа pointersBuilder
	pointer1 := NewPointersBuilder()
	pointer2 := NewPointersBuilder()

	// При помощи конструктора в билдере задаем нужные нам координаты соответственно
	pointer1.SetX(1).SetY(1).Build()
	pointer2.SetX(4).SetY(5).Build()

	fmt.Println(pointer1, pointer2)
}
