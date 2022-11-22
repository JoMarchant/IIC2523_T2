package main

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"log"
	"product/api/models"
	"strconv"
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
	if err != nil && err != sql.ErrNoRows {
		log.Print(err)
	}
}

func createProduct(c *gin.Context) {
	result := models.CreateProduct(c)

	if result == nil {
		c.JSON(200, gin.H{"message": "Record Created!"})
	} else {
		c.JSON(500, gin.H{"message": "Error!"})
		log.Print(result)
	}
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
	id_param := c.Param("id")
	id, err := strconv.ParseInt(id_param, 10, 32)

	checkErr(err)

	product, err := models.GetProduct(int(id))
	checkErr(err)

	if err != nil || product.Created_at == "" {
		c.JSON(404, gin.H{"message": "No Record Found!"})
	} else {
		c.JSON(200, gin.H{"data": product})
	}
}

func updateProduct(c *gin.Context) {
	id_param := c.Param("id")
	id, err := strconv.ParseInt(id_param, 10, 32)
	checkErr(err)
	result := models.UpdateProduct(int(id), c)

	if result == nil {
		c.JSON(200, gin.H{"message": "Record Updated!"})
	} else if result.Error() == "ERROR: Se esperaba una fila afectada" {
		c.JSON(404, gin.H{"message": "No Record Found!"})
	} else {
		c.JSON(500, gin.H{"message": "Error!"})
		log.Print(result)
	}
}
func deleteProduct(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Record Deleted!"})
}
