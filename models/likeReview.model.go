package models

import (
	"github.com/rfauzi44/up-echo/db"
	"github.com/rfauzi44/up-echo/libs"
	"net/http"
)

type LikeReview struct {
	IDReview string `json:"id_review"`
	IDMember string `json:"id_member"`
}

func AddLikeReview(id_review, id_member string) (libs.Response, error) {
	var res libs.Response

	conn := db.Connect()

	sqlStatement := "INSERT like_review (id_review, id_member) VALUES (?, ?)"

	stmt, err := conn.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	_, err = stmt.Exec(id_review, id_member)
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Success"

	return res, nil
}

func DislikeReview(id_review, id_member string) (libs.Response, error) {
	var res libs.Response

	conn := db.Connect()

	sqlStatement := "DELETE FROM like_review WHERE id_review=? AND id_member=?"

	stmt, err := conn.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	_, err = stmt.Exec(id_review, id_member)
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Success"

	return res, nil
}
