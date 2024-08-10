package repo

import (
	"fmt"
	"math"
	"math/rand"
	"sportsstore/models"
)

func (repo *MemoryRepo) Seed() {
	repo.categories = make([]models.Category, 3)
	for i := 0; i < 3; i++ {
		catName := fmt.Sprintf("Category_%v", i+1)
		repo.categories[i] = models.Category{ID: i + 1, CategoryName: catName}
	}
	for i := 0; i < 20; i++ {
		name := fmt.Sprintf("Product_%v", i+1)
		price := rand.Float64() * float64(rand.Intn(500))
		cat := &repo.categories[rand.Intn(len(repo.categories))]
		repo.products = append(repo.products, models.Product{
			ID:   i + 1,
			Name: name, Price: price,
			Description: fmt.Sprintf("%v (%v)", name, cat.CategoryName),
			Category:    cat,
		})
	}
}
func (repo *MemoryRepo) GetProductPage(page, pageSize int) ([]models.Product, int) {
    return getPage(repo.products, page, pageSize), len(repo.products)
}
func getPage(src []models.Product, page, pageSize int) []models.Product {
    start := (page -1) * pageSize
    if page > 0 && len(src) > start {
        end := (int)(math.Min((float64)(len(src)), (float64)(start + pageSize)))
        return src[start : end]
    }
    return []models.Product{}
}