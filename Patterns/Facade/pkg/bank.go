package pkg

import (
	"errors"
	"fmt"
	"time"
)

type Bank struct {
	Name  string
	Cards []Card
}

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
			println(fmt.Sprintf("[Банк] Остаток положительный! Составляет %s рублей", card.Balance))
			return nil
		}
	}

	return nil
}
