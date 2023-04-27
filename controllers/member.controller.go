package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rfauzi44/up-echo/models"
)

func GetAllMember(c echo.Context) error {
	result, err := models.GetAllMember()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)

}

func AddMember(c echo.Context) error {
	var member models.Member
	err := c.Bind(&member)
	if err != nil {
		return err
	}
	result, err := models.AddMember(member.Username, member.Gender, member.Skintype, member.Skincolor)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	return c.JSON(http.StatusOK, result)
}

func UpdateMember(c echo.Context) error {

	id_member := c.Param("id_member")
	var member models.Member
	err := c.Bind(&member)
	if err != nil {
		return err
	}
	result, err := models.UpdateMember(id_member, member.Username, member.Gender, member.Skintype, member.Skincolor)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	return c.JSON(http.StatusOK, result)
}

func DeleteMember(c echo.Context) error {

	id_member := c.Param("id_member")
	result, err := models.DeleteMember(id_member)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	return c.JSON(http.StatusOK, result)
}
