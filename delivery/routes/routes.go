package routes

import (
	"ecommerce/delivery/controllers/auth"
	"ecommerce/delivery/controllers/category"
	"ecommerce/delivery/controllers/order"
	"ecommerce/delivery/controllers/product"
	"ecommerce/delivery/controllers/shoppingcart"
	"ecommerce/delivery/controllers/user"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RegisterPath(e *echo.Echo, uc *user.UsersController, pc *product.ProductController, sc *shoppingcart.ShoppingCartController, cc *category.CategoryController, oc *order.OrderController, ac *auth.AuthController) {
	e.Pre(middleware.RemoveTrailingSlash())
	auth := e.Group("")
	auth.Use(middleware.JWT([]byte("RAHASIA")))
	e.GET("/products", pc.GetAllProductCtrl())
	e.POST("/products", pc.CreateProductControllers())
	e.PUT("/products/:id", pc.UpdateProductCtrl())
	e.DELETE("/products/:id", pc.DeleteProductCtrl())
	auth.POST("/shoppingcart", sc.CreateShoppingCartCtrl())
	auth.PUT("/shoppingcart/:id", sc.UpdateShoppingCartCtrl())
	auth.DELETE("/shoppingcart/:id", sc.DeleteShoppingCartCtrl())
	e.GET("/products", pc.GetAllProductCtrl())
	e.GET("/productss/", pc.FilterProductCtrl())
	e.GET("/products/:id", pc.GetProductCtrl())
	auth.POST("/products", pc.CreateProductControllers())
	auth.PUT("/products/:id", pc.UpdateProductCtrl())
	auth.DELETE("/products/:id", pc.DeleteProductCtrl())
	e.POST("/users/register", uc.RegisterUserCtrl())
	e.POST("/users/login", ac.LoginAuthCtrl())
	auth.GET("/users", uc.GetUserCtrl())
	auth.PUT("/users", uc.UpdateUserCtrl())
	auth.DELETE("/users", uc.DeleteUserCtrl())
	e.GET("/category", cc.GetAllCategoryCtrl())
	auth.POST("/category", cc.PostCategoryCtrl())
	auth.PUT("/category/:id", cc.UpdateCategoryCtrl())
	auth.DELETE("/category/:id", cc.DeleteCategoryCtrl())

	auth.POST("/order", oc.CreateOrderCtrl())
	auth.PUT("/order/:id", oc.)
	auth.DELETE("/order/:id", cc.DeleteCategoryCtrl())

	e.Logger.Fatal(e.Start(":8000"))

}
