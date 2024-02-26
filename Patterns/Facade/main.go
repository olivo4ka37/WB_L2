package main

import (
	"WB_L2/Patterns/Facade/pkg"
	"fmt"
)

// Инициализация банка, двух карт, двух пользователей, одного продукта и магазина с этим продуктом.
var (
	bank = pkg.Bank{
		Name:  "Банк",
		Cards: []pkg.Card{},
	}

	card1 = pkg.Card{
		Name:    "CARD-1",
		Balance: 200,
		Bank:    &bank,
	}

	card2 = pkg.Card{
		Name:    "CARD-2",
		Balance: 5,
		Bank:    &bank,
	}

	user = pkg.User{
		Name: "Андрей",
		Card: &card1,
	}

	user2 = pkg.User{
		Name: "Андриана",
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
	// В данном примере функция структуры магазина Sell является неким фасадом,
	// за которым происходит углубленный вызов другой бизнес логики.
	// В этом примере это легко проследить так как каждая функция пишет о своём вызове,
	// указывая структуру к которой она принадлежит.

	println("[Банк] Выпуск карт")

	// Добавление карт в слайс банка
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
