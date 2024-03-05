package main

// Command интерфейс самой команды
type Command interface {
	execute()
}
