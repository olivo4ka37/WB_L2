package pkg

import (
	"errors"
	"fmt"
	"time"
)

type Shop struct {
	Name     string
	Products []Product
}

func (shop Shop) Sell(user User, product string) error {
	println("[Магазин] запрос к пользователю, для получения остатка по карте")
	time.Sleep(time.Second * 1)
	err := user.Card.CheckBalance()
	if err != nil {
		return err
	}
	fmt.Printf("[Магазин] Проверка, может ли пользователь купить товар! \n", user.Name)
	time.Sleep(time.Millisecond * 300)
	for _, prod := range shop.Products {
		if prod.Name != product {
			continue
		}

		if prod.Price > user.GetBalance() {
			return errors.New("[Магазин] Недостаточно средств для покупки товара!")
		}
		fmt.Printf("[Магазин] Товар [%s] - куплен!\n", prod.Name)
	}

	return nil
}
