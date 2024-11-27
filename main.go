package main

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

type Product struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

var products = []Product{
	{ID: 1, Name: "iPhone 16", Price: 8500},
	{ID: 2, Name: "iPhone", Price: 4500},
}

func main() {
	router := gin.Default()

	// маршруты
	router.LoadHTMLGlob("templates/*")

	// главная
	router.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})

	// страница с карточками
	router.GET("/cards", func(c *gin.Context) {
		c.HTML(200, "cards.html", gin.H{
			"products": products,
		})
	})

	// все товары (апи)
	router.GET("/products", func(c *gin.Context) {
		c.JSON(200, products)
	})

	// добавление
	router.POST("/products", func(c *gin.Context) {
		var newProduct Product
		if err := c.ShouldBindJSON(&newProduct); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		newProduct.ID = len(products) + 1
		products = append(products, newProduct)
		c.JSON(201, newProduct)
	})

	// редактирование
	router.PUT("/products/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(400, gin.H{"error": "Invalid ID"})
			return
		}

		var updatedProduct Product
		if err := c.ShouldBindJSON(&updatedProduct); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		for i, p := range products {
			if p.ID == id {
				updatedProduct.ID = id
				products[i] = updatedProduct
				c.JSON(200, updatedProduct)
				return
			}
		}
		c.JSON(404, gin.H{"message": "Product not found"})
	})

	// удаление
	router.DELETE("/products/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(400, gin.H{"error": "Invalid ID"})
			return
		}

		for i, p := range products {
			if p.ID == id {
				products = append(products[:i], products[i+1:]...)
				c.JSON(200, gin.H{"message": "Product deleted"})
				return
			}
		}
		c.JSON(404, gin.H{"message": "Product not found"})
	})

	router.Run(":8080")
}
