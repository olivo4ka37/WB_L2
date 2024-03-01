package main

import "fmt"

// Tv структура имитирующая телевизор
// имеет состояние isRunning (Включен или Выключен)
type Tv struct {
	isRunning bool
}

// on включает телевизор
func (t *Tv) on() {
	t.isRunning = true
	fmt.Println("Turning tv on")
}

// off выключает телевизор
func (t *Tv) off() {
	t.isRunning = false
	fmt.Println("Turning tv off")
}
