package main

import (
	"WB_L2/Patterns/Facade/pkg"
	"fmt"
)

var (
	bank = pkg.Bank{
		Name:  "Банк",
		Cards: []pkg.Card{},
	}

	card1 = pkg.Card{
		Name:    "CRD-1",
		Balance: 200,
		Bank:    &bank,
	}

	card2 = pkg.Card{
		Name:    "CRD-2",
		Balance: 5,
		Bank:    &bank,
	}

	user = pkg.User{
		Name: "Покупатель1",
		Card: &card1,
	}

	user2 = pkg.User{
		Name: "Покупатель2",
		Card: &card2,
	}

	prod = pkg.Product{
		Name:  "Сыр",
		Price: 150,
	}

	shop = pkg.Shop{
		Name: "SHOP",
		Products: []pkg.Product{
			prod,
		},
	}
)

func main() {
	println("[Банк] Выпуск карт")
	bank.Cards = append(bank.Cards, card1, card2)
	fmt.Printf("[%s]\n", user.Name)
	err := shop.Sell(user, prod.Name)
	if err != nil {
		println(err.Error())
		return
	}

	fmt.Printf("[%s]\n", user2.Name)
	err = shop.Sell(user2, prod.Name)
	if err != nil {
		println(err.Error())
		return
	}

}
