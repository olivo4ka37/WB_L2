package pkg

import "time"

// Card структура карты банка, Name - Имя карты
// Balance баланс карты
// Bank банк в котором была выпущена карта
type Card struct {
	Name    string
	Balance float64
	Bank    *Bank
}

// CheckBalance функция реализующая проверку баланса карты, возвращает ошибку если баланс меньше или равен 0
func (card Card) CheckBalance() error {
	println("[Карта] Запрос в банк для проверки остатка")
	time.Sleep(time.Second * 1)
	card.Bank.CheckBalance(card.Name)
	return card.Bank.CheckBalance(card.Name)
}
