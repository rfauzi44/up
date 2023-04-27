package models

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/rfauzi44/up-echo/db"
	"github.com/rfauzi44/up-echo/libs"
)

type Member struct {
	IDMember  string `json:"id_member"`
	Username  string `json:"username"`
	Gender    string `json:"gender"`
	Skintype  string `json:"skintype"`
	Skincolor string `json:"skincolor"`
}

func GetAllMember() (libs.Response, error) {
	var obj Member
	var objArr []Member
	var res libs.Response

	conn := db.Connect()

	sqlStatement := "SELECT * FROM member"

	rows, err := conn.Query(sqlStatement)

	if err != nil {
		return res, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&obj.IDMember, &obj.Username, &obj.Gender, &obj.Skintype, &obj.Skincolor)
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

func AddMember(username, gender, skintype, skincolor string) (libs.Response, error) {
	var res libs.Response

	uuid := uuid.New()
	id_member := uuid.String()

	conn := db.Connect()

	sqlStatement := "INSERT member (id_member, username, gender, skintype, skincolor) VALUES (?, ?, ?, ?, ?)"

	stmt, err := conn.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	_, err = stmt.Exec(id_member, username, gender, skintype, skincolor)
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Success"

	return res, nil
}

func UpdateMember(id_member, username, gender, skintype, skincolor string) (libs.Response, error) {
	var res libs.Response

	conn := db.Connect()

	sqlStatement := "UPDATE member SET username = ?, gender = ?, skintype = ?, skincolor = ? WHERE id_member = ?"

	stmt, err := conn.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	_, err = stmt.Exec(username, gender, skintype, skincolor, id_member)
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Success"

	return res, nil
}

func DeleteMember(id_member string) (libs.Response, error) {
	var res libs.Response

	conn := db.Connect()

	sqlStatement := "DELETE FROM member WHERE id_member = ?"

	stmt, err := conn.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	_, err = stmt.Exec(id_member)
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Success"

	return res, nil
}
