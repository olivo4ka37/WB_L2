package main

// Button структура кнопки, которая будет выполнять команду
type Button struct {
	command Command
}

// press функция при вызове которой, вызывается команда
func (b *Button) press() {
	b.command.execute()
}
