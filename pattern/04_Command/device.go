package main

// Device интерфейс девайса который можем включать и выключать
type Device interface {
	on()
	off()
}
