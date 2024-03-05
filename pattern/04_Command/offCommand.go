package main

// OffCommand структура содержащая в себе команду выключения девайса
type OffCommand struct {
	device Device
}

// execute выключает девайс
func (c *OffCommand) execute() {
	c.device.off()
}
