package routes

import (
	"ecommerce/delivery/controllers/auth"
	"ecommerce/delivery/controllers/product"
	"ecommerce/delivery/controllers/user"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RegisterPath(e *echo.Echo, uc *user.UsersController, pc *product.ProductController, ac *auth.AuthController) {
	e.Pre(middleware.RemoveTrailingSlash())
	e.GET("/products", pc.GetAllProductCtrl())

	e.Logger.Fatal(e.Start(":8000"))

}
