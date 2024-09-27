package tests

import (
	"bookstore-api/controller"
	"bookstore-api/services"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestGetBooks(t *testing.T) {
	router := gin.Default()
	router.GET("/books", controller.GetBooks)

	req, _ := http.NewRequest("GET", "/books", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code 200, got %d", w.Code)
	}
}

func BenchmarkGetBooks(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := services.GetBooks()
		if err != nil {
			b.Errorf("Error fetching books: %v", err)
		}
	}
}
