package main

// OnCommand структура содержащая в себе команду включения девайса
type OnCommand struct {
	device Device
}

// execute включает девайс
func (c *OnCommand) execute() {
	c.device.on()
}
