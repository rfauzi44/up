package routers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rfauzi44/up-echo/controllers"
)

func New() *echo.Echo {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		link := "<a href='https://documenter.getpostman.com/view/25042327/2s93eSYvCx'>here</a>"
		response := "Hello World! This is up-echo API. You can check Postman Documentation " + link
		return c.HTML(http.StatusOK, response)
	})

	//Member
	e.GET("/member", controllers.GetAllMember)
	e.POST("/member", controllers.AddMember)
	e.PUT("/member/:id_member", controllers.UpdateMember)
	e.DELETE("/member/:id_member", controllers.DeleteMember)

	//Product
	e.POST("/product", controllers.AddProduct)
	e.GET("/product/:id_product", controllers.GetProductById)
	e.GET("/product", controllers.GetAllProduct)

	//ReviewProduct
	e.POST("/product/review", controllers.AddReviewProduct)
	e.GET("/product/review", controllers.GetAllReview)

	//LikeReview
	e.POST("/product/review/like", controllers.AddLikeReview)
	e.DELETE("/product/review/like", controllers.DislikeReview)

	return e
}
