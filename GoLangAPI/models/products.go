package models

import (
	"errors"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/gin-gonic/gin"
)

var DB *sql.DB

func ConnectDatabase() error {
	db, err := sql.Open("sqlite3", "./sqlite.db")
	if err != nil {
		return err
	}

	DB = db
	return nil
}

type Product struct {
	Id         int    `json:"id"`
	Name     string `json:"name"`
	Description      string `json:"description"`
	Price    int `json:"price"`
	Exp_date      string `json:"exp_date"`
	Created_at string `json:"created_at"`
}

func GetProducts() ([]Product, error) {

	rows, err := DB.Query("SELECT * FROM products")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	Products := make([]Product, 0)

	for rows.Next() {
		singleProduct := Product{}
		err = rows.Scan(
			&singleProduct.Id,
			&singleProduct.Name,
			&singleProduct.Description,
			&singleProduct.Price,
			&singleProduct.Exp_date,
			&singleProduct.Created_at)

		if err != nil {
			return nil, err
		}

		Products = append(Products, singleProduct)
	}

	err = rows.Err()

	if err != nil {
		return nil, err
	}

	return Products, nil
}

func CreateProduct(c *gin.Context) error {
	stmt, err := DB.Prepare("INSERT INTO products (name, description, price, exp_date) VALUES (?, ?, ?, ?)")

	if err != nil {
		return err
	}

	defer stmt.Close()

	var productBody Product

	c.BindJSON(&productBody)

	r, err := stmt.Exec(productBody.Name, productBody.Description, productBody.Price, productBody.Exp_date)

	if err != nil {
		return err
	}

	if i, err := r.RowsAffected(); err != nil || i != 1 {
		return errors.New("ERROR: Se esperaba una fila afectada")
	}

	return nil
}
