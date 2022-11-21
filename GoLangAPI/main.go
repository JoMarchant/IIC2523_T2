package main

import (
	"github.com/gin-gonic/gin"
	"product/api/models"
	"log"
)

func main() {
	err := models.ConnectDatabase()
	checkErr(err)
	r := gin.Default()

	// API v1
	router := r.Group("/products")
	{
		router.POST("/", createProduct)
		router.GET("/", getProducts)
		router.GET("/:id", getProduct)
		router.PATCH("/:id", updateProduct)
		router.DELETE("/:id", deleteProduct)
	}

	// By default it serves on :8080
	r.Run()
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func createProduct(c *gin.Context) {
	c.JSON(200, gin.H{"message": "A new Record Created!"})
}

func getProducts(c *gin.Context) {
	products, err := models.GetProducts()
	checkErr(err)

	if products == nil {
		c.JSON(404, gin.H{"message": "No Record Found!"})
	} else {
		c.JSON(200, gin.H{"data": products})
	}
}

func getProduct(c *gin.Context) {
	c.JSON(200, gin.H{"message": "All Records"})
}

func updateProduct(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Record Updated!"})
}
func deleteProduct(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Record Deleted!"})
}
