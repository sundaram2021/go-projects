package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
)

type MenuItem struct {
	Item   string  `json:"item"`
	Recipe string  `json:"recipe"`
	Price  float64 `json:"price"`
}

type Order struct {
	Order []string `json:"orders"` // Correct struct field name and JSON tag
}

var total float64 // Use float64 to match the Price type

func GetRecipes(c *gin.Context) {
	var items []MenuItem

	data, err := ioutil.ReadFile("recipes.json")
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}

	if err = json.Unmarshal(data, &items); err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, items)
}

func CreateOrder(c *gin.Context) {
	var orders Order

	if err := c.ShouldBindJSON(&orders); err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"Error": err.Error()})
		return
	}

	var items []MenuItem

	data, err := ioutil.ReadFile("recipes.json")
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}

	if err = json.Unmarshal(data, &items); err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}

	// Create a map to store item prices for quick lookup
	priceMap := make(map[string]float64)
	for _, item := range items {
		priceMap[item.Item] = item.Price
	}

	// Calculate the total price
	total = 0 // Reset total to 0 before calculation
	for _, orderItem := range orders.Order {
		if price, exists := priceMap[orderItem]; exists {
			total += price
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"total": total,
	})
}

func UpdateOrder(c *gin.Context) {
	CreateOrder(c)
}

func GetOrder(c *gin.Context){
	var orders *Order
	c.JSON(http.StatusOK, gin.H{
		"orders": orders,
	})
}

func DeleteOrder(c *gin.Context) {
	var orders *Order

	if orders == nil {
		orders = &Order{} // Properly initialize the pointer
	}
	// Get the item name from the URL parameter
	itemToDelete, err := url.QueryUnescape(c.Param("item"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid item name"})
		return
	}

	var index = -1

	for i, orderItem := range orders.Order {
		if orderItem == itemToDelete {
			index = i
			break
		}
	}
	var items []MenuItem


	orders.Order = append(orders.Order[:index], orders.Order[index+1:]...)

	// Create a map to store item prices for quick lookup
	priceMap := make(map[string]float64)
	for _, item := range items {
		priceMap[item.Item] = item.Price
	}

	// Calculate the total price
	// total = 0 // Reset total to 0 before calculation
	for _, orderItem := range orders.Order {
		if price, exists := priceMap[orderItem]; exists {
			total += price
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"total": total,
	})
}

func main() {
	routes := gin.Default()

	routes.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "welcome to recipes api"})
	})

	routes.GET("/recipes", GetRecipes)
	routes.POST("/orders", CreateOrder)
	routes.PUT("/orders", UpdateOrder)
	routes.DELETE("/orders/:item", DeleteOrder)

	fmt.Println("server is running on port 8080...")
	routes.Run(":8080")
}
