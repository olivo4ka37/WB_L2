package pkg

import (
	"errors"
	"fmt"
	"time"
)

// Bank структура имитирующая работу банка
// имеет два поля Name - имя банка
// Cards - слайс типа Card - Все карты принадлежащие этому банку
type Bank struct {
	Name  string
	Cards []Card
}

// CheckBalance функция проверяющая баланс пользователя, если баланс меньше или равен 0 возвращает ошибку
func (bank Bank) CheckBalance(cardNumber string) error {
	println(fmt.Sprintf("[Банк] Получение остатка по карте %s", cardNumber))
	time.Sleep(time.Millisecond * 300)

	for _, card := range bank.Cards {
		if card.Name != cardNumber {
			continue
		}

		if card.Balance <= 0 {
			return errors.New("[Банк] Недостаточно средств!")
		} else {
			println(fmt.Sprintf("[Банк] Остаток положительный! Составляет %.2f рублей", card.Balance))
			return nil
		}
	}

	return nil
}
