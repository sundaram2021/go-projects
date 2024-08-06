package users

import "github.com/sundaram2021/go-projects/ecommerce/cart"

type User struct {
	Name    string
	ID      int32
	Balance float64
	Orders  cart.Cart
}

func (u *User) AbleToBuyCartItem() bool {
	return u.Orders.Total <= u.Balance
}

func (u *User) isUserBalanceZero() bool {
	return u.Balance == 0
}

