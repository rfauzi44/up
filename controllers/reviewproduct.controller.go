package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rfauzi44/up-echo/models"
)

func AddReviewProduct(c echo.Context) error {
	var reviewproduct models.ReviewProduct
	err := c.Bind(&reviewproduct)
	if err != nil {
		return err
	}
	result, err := models.AddReviewProduct(reviewproduct.IDProduct, reviewproduct.IDMember, reviewproduct.DescReview)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	return c.JSON(http.StatusOK, result)
}
