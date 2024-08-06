package cart

import (
	"github.com/sundaram2021/go-projects/ecommerce/products"
)

type Cart struct {
	Total float64
	Products products.Products
}


func (c *Cart) getItems() []products.Product {
	var items []products.Product

	for _, ele := range c.Products.Products {
		items = append(items, ele)
	}
	return items
}

func (c *Cart) getTotalPrice() float64 {
	var price float64

	for _, ele := range c.Products.Products {
		price += ele.Price
	}
	return price
}