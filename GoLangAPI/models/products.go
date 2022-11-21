package models

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
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
	Fecha_exp      string `json:"fecha_exp"`
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
			&singleProduct.Fecha_exp,
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
