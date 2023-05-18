package models

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/rfauzi44/up/db"
	"github.com/rfauzi44/up/libs"
)

type ReviewProduct struct {
	IDReview   string `json:"id_review"`
	IDProduct  string `json:"id_product"`
	IDMember   string `json:"id_member"`
	DescReview string `json:"desc_review"`
	LikeCount  int    `json:"like_count"`
}

func AddReviewProduct(id_product, id_member, desc_review string) (libs.Response, error) {
	var res libs.Response

	uuid := uuid.New()
	id_review := uuid.String()

	conn := db.Connect()

	sqlStatement := "INSERT review_product (id_review, id_product, id_member, desc_review ) VALUES (?, ?, ?, ?)"

	stmt, err := conn.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	_, err = stmt.Exec(id_review, id_product, id_member, desc_review)
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]string{
		"last_inserted_id": id_review,
	}

	return res, nil
}

func GetAllReview() (libs.Response, error) {
	var obj ReviewProduct
	var objArr []ReviewProduct
	var res libs.Response

	conn := db.Connect()

	sqlStatement := "SELECT * FROM review_product"

	rows, err := conn.Query(sqlStatement)

	if err != nil {
		return res, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&obj.IDReview, &obj.IDProduct, &obj.IDMember, &obj.DescReview, &obj.LikeCount)
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
