package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rfauzi44/up/models"
)

func AddLikeReview(c echo.Context) error {
	var likeReview models.LikeReview
	err := c.Bind(&likeReview)
	if err != nil {
		return err
	}
	result, err := models.AddLikeReview(likeReview.IDReview, likeReview.IDMember)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	return c.JSON(http.StatusOK, result)
}

func DislikeReview(c echo.Context) error {

	var likeReview models.LikeReview

	err := c.Bind(&likeReview)
	if err != nil {
		return err
	}
	result, err := models.DislikeReview(likeReview.IDReview, likeReview.IDMember)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	return c.JSON(http.StatusOK, result)
}
