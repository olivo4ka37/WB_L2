package main

import (
	"fmt"
	"strconv"
)

/*
	Реализовать паттерн «стратегия».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Strategy_pattern
*/

/*
Стратегия - паттерн поведения объектов.
Он определяет семейство алгоритмов, инкапсулирует каждый из них и делает их взаимозаменяемыми.
Стратегия позволяет изменять алгоритмы независимо от клиентов, которые ими пользуются.

Применимость паттерна:
- когда имеется много родственных классов, отличающихся только поведением;
- когда нужно иметь несколько разных вариантов алгоритма;
- когда в алгоритме содержатся данные, о которых клиент не должен "знать";
- когда в классе определено много поведений, что представлено разветвленными условными операторами.

Плюсы:
- семейства родственных алгоритмов;
- альтернатива порождению подклассов;
- с помощью стратегий можно избавиться от условных операторов;
- выбор реализации;
- обмен информацией между стратегией и контекстом;

Минусы
- клиенты должны "знать" о различных стратегиях;
- увеличение числа объектов.

Примеры использования:
- Библиотеки ЕТ++ [WGM88] и Interviews используют стратегии для инкапсуляции алгоритмов разбиения на строки.
*/

// StrategySort Интерфейс в котором будут храниться все наши стратегии
type StrategySort interface {
	Sort([]int)
}

// BubbleSort структура методом которой является сортировка пузырьком
type BubbleSort struct {
}

// Sort сортировка пузырьком
func (s *BubbleSort) Sort(a []int) {
	size := len(a)
	if size < 2 {
		return
	}
	for i := 0; i < size; i++ {
		for j := size - 1; j >= i+1; j-- {
			if a[j] < a[j-1] {
				a[j], a[j-1] = a[j-1], a[j]
			}
		}
	}
}

// InsertionSort структура методом которой является сортировка вставками
type InsertionSort struct {
}

// Sort сортировка вставками
func (s *InsertionSort) Sort(a []int) {
	size := len(a)
	if size < 2 {
		return
	}
	for i := 1; i < size; i++ {
		var j int
		var buff = a[i]
		for j = i - 1; j >= 0; j-- {
			if a[j] < buff {
				break
			}
			a[j+1] = a[j]
		}
		a[j+1] = buff
	}
}

// Context структура инкапсулирующая стратегию сортировки
type Context struct {
	strategy StrategySort
}

// Algorithm функция реализующая выбор стратегии
func (c *Context) Algorithm(a StrategySort) {
	c.strategy = a
}

// Sort вызывает сортировку
func (c *Context) Sort(s []int) {
	c.strategy.Sort(s)
}

func main() {
	// Инициализируем два массива
	data1 := []int{8, 2, 6, 7, 1, 3, 9, 5, 4}
	data2 := []int{8, 2, 6, 7, 1, 3, 9, 5, 4}

	// Проверяем работу сортировки пузырьком
	ctx := new(Context)
	ctx.Algorithm(&BubbleSort{})
	ctx.Sort(data1)

	var result1 string
	for _, val := range data1 {
		result1 += strconv.Itoa(val) + ","
	}
	fmt.Println(result1)

	// Проверяем работу сортировки вставками
	ctx.Algorithm(&InsertionSort{})
	ctx.Sort(data2)

	var result2 string
	for _, val := range data2 {
		result2 += strconv.Itoa(val) + ","
	}
	fmt.Println(result2)
}
