package pkg

// User структура имитирующая структуру пользователя магазина
type User struct {
	Name string
	Card *Card
}

// GetBalance возвращает баланс карты пользователя
func (user User) GetBalance() float64 {
	return user.Card.Balance
}
