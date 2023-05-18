package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rfauzi44/up/models"
)

func AddProduct(c echo.Context) error {
	var product models.Product
	err := c.Bind(&product)
	if err != nil {
		return err
	}
	result, err := models.AddProduct(product.NameProduct, product.Price)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	return c.JSON(http.StatusOK, result)
}

func GetProductById(c echo.Context) error {

	id_product := c.Param("id_product")
	result, err := models.GetProductById(id_product)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	return c.JSON(http.StatusOK, result)
}

func GetAllProduct(c echo.Context) error {
	result, err := models.GetAllProduct()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)

}
