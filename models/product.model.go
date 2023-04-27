package models

import (
	"database/sql"
	"net/http"

	"github.com/google/uuid"
	"github.com/rfauzi44/up-echo/db"
	"github.com/rfauzi44/up-echo/libs"
)

type Product struct {
	IDProduct   string `json:"id_product"`
	NameProduct string `json:"name_product"`
	Price       int    `json:"price"`
}

func AddProduct(name_product string, price int) (libs.Response, error) {
	var res libs.Response

	uuid := uuid.New()
	id_product := uuid.String()

	conn := db.Connect()

	sqlStatement := "INSERT product (id_product, name_product, price) VALUES (?, ?, ?)"

	stmt, err := conn.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	_, err = stmt.Exec(id_product, name_product, price)
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Success"

	return res, nil
}

func GetProductById(id_product string) (libs.Response, error) {
	var obj Product
	var res libs.Response

	conn := db.Connect()

	sqlStatement := "SELECT * FROM product WHERE id_product = ?"

	row := conn.QueryRow(sqlStatement, id_product)

	err := row.Scan(&obj.IDProduct, &obj.NameProduct, &obj.Price)
	if err != nil {
		if err == sql.ErrNoRows {
			res.Status = http.StatusNotFound
			res.Message = "Product not found"
		} else {
			res.Status = http.StatusInternalServerError
			res.Message = "Internal server error"
		}
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = obj
	return res, nil
}
