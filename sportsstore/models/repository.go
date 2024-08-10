package models

type Repository interface {
	GetProduct(id int) Product
	GetProducts() []Product
	GetProductPage(page, pageSize int) (products []Product, totalAvailable int)
	GetCategories() []Category
	SaveProduct(*Product)
	SaveCategory(*Category)
	GetProductPageCategory(categoryId int, page, pageSize int) (products []Product,
		totalAvailable int)
	GetOrder(id int) Order
	GetOrders() []Order
	SaveOrder(*Order)
	SetOrderShipped(*Order)
	Seed()
	Init()
}
