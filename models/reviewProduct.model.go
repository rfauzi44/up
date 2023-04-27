package models

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/rfauzi44/up-echo/db"
	"github.com/rfauzi44/up-echo/libs"
)

type ReviewProduct struct {
	IDReview   string `json:"id_review"`
	IDProduct  string `json:"id_product"`
	IDMember   string `json:"id_member"`
	DescReview string `json:"desc_review"`
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

	return res, nil
}
