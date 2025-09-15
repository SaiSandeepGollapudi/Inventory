package dataservice

import (
	"INVENTORY/model"
	"database/sql"
)

func CreateProduct(db *sql.DB, Product model.Product) error {
	query := `INSERT INTO products1(id, name, quantity) VALUES (?, ?, ?)`
	_, err := db.Exec(query, Product.Id, Product.Name, Product.Quantity)
	if err != nil {
		return err
	}
	return nil
}
