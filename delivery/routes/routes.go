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

func RegisterPath(e *echo.Echo, uc *user.UsersController, pc *product.ProductController, sc *shoppingcart.ShoppingCartController, cc *category.CategoryController, ac *auth.AuthController, oc *order.OrderController) {
	e.Pre(middleware.RemoveTrailingSlash())
	auth := e.Group("")
	auth.Use(middleware.JWT([]byte("RAHASIA")))
	//Users
	e.POST("/users/register", uc.RegisterUserCtrl())
	e.POST("/users/login", ac.LoginAuthCtrl())
	auth.GET("/users", uc.GetUserCtrl())
	auth.PUT("/users", uc.UpdateUserCtrl())
	auth.DELETE("/users", uc.DeleteUserCtrl())
	//Category
	e.GET("/category", cc.GetAllCategoryCtrl())
	e.GET("/category/:id", cc.GetCategoryCtrl())
	auth.POST("/category", cc.PostCategoryCtrl())
	auth.PUT("/category/:id", cc.UpdateCategoryCtrl())
	auth.DELETE("/category/:id", cc.DeleteCategoryCtrl())
	//ShoppingCart
	auth.GET("/shoppingcart", sc.GetShppingCartCtrl())
	auth.POST("/shoppingcart", sc.CreateShoppingCartCtrl())
	auth.PUT("/shoppingcart/:id", sc.UpdateShoppingCartCtrl())
	auth.DELETE("/shoppingcart/:id", sc.DeleteShoppingCartCtrl())
	//Orders
	auth.GET("/orders", oc.GetAllOrdersCtrl())
	auth.GET("/orders/:id", oc.GetOrderCtrl())
	auth.POST("/orders", oc.CreateOrderCtrl())
	auth.PUT("/orders/cancel/:id", oc.CancelOrderCtrl())
	auth.POST("/orders/pay/:id", oc.PayOrderCtrl())
	auth.POST("/orders/check/:id", oc.CheckOrderCtrl())
	//Products
	e.GET("/products", pc.GetAllProductCtrl())
	e.GET("/pagination", pc.Pagination())
	e.GET("/products/:id", pc.GetProductCtrl())
	auth.POST("/products", pc.CreateProductControllers())
	auth.PUT("/products/:id", pc.UpdateProductCtrl())
	auth.DELETE("/products/:id", pc.DeleteProductCtrl())

	e.Logger.Fatal(e.Start(":8000"))

}
