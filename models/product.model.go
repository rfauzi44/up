package models

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/rfauzi44/up-echo/db"
	"github.com/rfauzi44/up-echo/libs"
)

type Product struct {
	IDProduct   string       `json:"id_product"`
	NameProduct string       `json:"name_product"`
	Price       int          `json:"price"`
	ReviewData  []ReviewData `json:"review_data,omitempty"`
}

type ReviewData struct {
	Username   string `json:"username"`
	Gender     string `json:"gender"`
	Skintype   string `json:"skintype"`
	Skincolor  string `json:"skincolor"`
	DescReview string `json:"desc_review"`
	LikeCount  int    `json:"like_count"`
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
	res.Data = map[string]string{
		"last_inserted_id": id_product,
	}

	return res, nil
}

func GetProductById(id_product string) (libs.Response, error) {
	var objProduct Product
	var res libs.Response

	conn := db.Connect()

	sqlStatement := "SELECT product.id_product, product.name_product, product.price, member.username, member.gender, member.skintype, member.skincolor, review_product.desc_review, review_product.like_count FROM product INNER JOIN review_product ON product.id_product = review_product.id_product INNER JOIN member ON review_product.id_member = member.id_member WHERE product.id_product = ?"

	rows, err := conn.Query(sqlStatement, id_product)

	if err != nil {
		return res, err
	}
	defer rows.Close()

	for rows.Next() {
		var reviewData ReviewData
		err = rows.Scan(
			&objProduct.IDProduct,
			&objProduct.NameProduct,
			&objProduct.Price,
			&reviewData.Username,
			&reviewData.Gender,
			&reviewData.Skintype,
			&reviewData.Skincolor,
			&reviewData.DescReview,
			&reviewData.LikeCount,
		)
		if err != nil {
			return res, err
		}
		objProduct.ReviewData = append(objProduct.ReviewData, reviewData)

	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = objProduct

	return res, nil

}

func GetAllProduct() (libs.Response, error) {
	var obj Product
	var objArr []Product
	var res libs.Response

	conn := db.Connect()

	sqlStatement := "SELECT * FROM product"

	rows, err := conn.Query(sqlStatement)

	if err != nil {
		return res, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(
			&obj.IDProduct,
			&obj.NameProduct,
			&obj.Price,
		)
		if err != nil {
			return res, err
		}
		objArr = append(objArr, obj)
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = objArr

	return res, nil

}
