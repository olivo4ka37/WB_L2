package pkg

import "time"

type Card struct {
	Name    string
	Balance float64
	Bank    *Bank
}

func (card Card) CheckBalance() error {
	println("[Карта] Запрос в банк для проверки остатка")
	time.Sleep(time.Second * 1)
	card.Bank.CheckBalance(card.Name)
	return card.Bank.CheckBalance(card.Name)
}
