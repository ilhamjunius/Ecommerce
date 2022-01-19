package routes

import (
	"ecommerce/constant"
	"ecommerce/delivery/controllers/auth"
	"ecommerce/delivery/controllers/category"
	"ecommerce/delivery/controllers/product"
	"ecommerce/delivery/controllers/shoppingcart"
	"ecommerce/delivery/controllers/user"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RegisterPath(e *echo.Echo, uc *user.UsersController, pc *product.ProductController, sc *shoppingcart.ShoppingCartController, cc *category.CategoryController, ac *auth.AuthController) {
	e.Pre(middleware.RemoveTrailingSlash())
	auth := e.Group("")
	auth.Use(middleware.JWT([]byte(constant.JWT_SECRET_KEY)))
	e.GET("/products", pc.GetAllProductCtrl())
	e.POST("/products", pc.CreateProductControllers())
	e.PUT("/products/:id", pc.UpdateProductCtrl())
	e.DELETE("/products/:id", pc.DeleteProductCtrl())
	e.POST("/shoppingcart", sc.CreateShoppingCartCtrl())
	e.PUT("/shoppingcart/:id", sc.UpdateShoppingCartCtrl())
	e.DELETE("/shoppingcart/:id", sc.DeleteShoppingCartCtrl())
	e.GET("/products", pc.GetAllProductCtrl())
	e.POST("/products", pc.CreateProductControllers())
	e.PUT("/products/:id", pc.UpdateProductCtrl())
	e.DELETE("/products/:id", pc.DeleteProductCtrl())
	e.POST("/users/register", uc.RegisterUserCtrl())
	e.POST("/users/login", ac.LoginAuthCtrl())
	auth.GET("/users", uc.GetUserCtrl())
	auth.PUT("/users", uc.UpdateUserCtrl())
	auth.DELETE("/users", uc.DeleteUserCtrl())
	e.POST("/category", cc.PostCategoryCtrl())
	e.PUT("/category/:id", cc.UpdateCategoryCtrl())
	e.DELETE("/category/:id", cc.DeleteCategoryCtrl())

	e.Logger.Fatal(e.Start(":8000"))

}
