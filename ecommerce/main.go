package main

import (
	"fmt"
	"github.com/sundaram2021/go-projects/ecommerce/cart"
	"github.com/sundaram2021/go-projects/ecommerce/products"
	"github.com/sundaram2021/go-projects/ecommerce/users"
)

func main() {
	// Initialize products
	myProducts := products.Products{
		Products: []products.Product{
			{Id: 1, Name: "name1", Qty: 22, Price: 34.55},
			{Id: 2, Name: "name2", Qty: 2, Price: 345.15},
		},
	}

	// Initialize the cart
	myCart := cart.Cart{
		Total:    50.0, // Example total
		Products: myProducts,
	}

	// Initialize the user
	myUser := users.User{
		Name:    "John Doe",
		ID:      123,
		Balance: 100.0, // Example balance
		Orders:  myCart,
	}

	// Example usage
	fmt.Println("User:", myUser.Name)
	fmt.Println("Cart Total:", myUser.Orders.Total)
	for _, product := range myUser.Orders.Products.Products {
		fmt.Printf("Product ID: %d, Name: %s, Qty: %d, Price: %.2f\n", product.Id, product.Name, product.Qty, product.Price)
	}

	if myUser.AbleToBuyCartItem() {
		fmt.Println("User can buy the items in the cart.")
	} else {
		fmt.Println("User cannot buy the items in the cart.")
	}
}
