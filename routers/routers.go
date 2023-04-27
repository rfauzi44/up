package routers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rfauzi44/up-echo/controllers"
)

func New() *echo.Echo {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.GET("/member", controllers.GetAllMember)
	e.POST("/member", controllers.AddMember)
	e.PUT("/member/:id_member", controllers.UpdateMember)
	e.DELETE("/member/:id_member", controllers.DeleteMember)

	e.POST("/product", controllers.AddProduct)
	e.GET("/product/:id_product", controllers.GetProductById)

	e.POST("/product/review", controllers.AddReviewProduct)

	e.POST("/product/review/like", controllers.AddLikeReview)
	e.DELETE("/product/review/like", controllers.DislikeReview)

	return e
}
