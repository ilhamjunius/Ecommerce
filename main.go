package main

import (
	"ecommerce/configs"
	authCon "ecommerce/delivery/controllers/auth"
	categoryCon "ecommerce/delivery/controllers/category"
	productCon "ecommerce/delivery/controllers/product"
	shoppingCartCon "ecommerce/delivery/controllers/shoppingcart"
	userCon "ecommerce/delivery/controllers/user"

	"ecommerce/delivery/routes"
	"ecommerce/repository/auth"
	"ecommerce/repository/category"
	product "ecommerce/repository/products"
	"ecommerce/repository/shoppingcart"
	user "ecommerce/repository/users"
	"ecommerce/utils"
	"fmt"
	"log"

	"github.com/labstack/echo/v4"
)

func main() {

	config := configs.GetConfig()

	db := utils.InitDB(config)
	productRepo := product.NewProductRepo(db)
	productController := productCon.NewProductControllers(productRepo)
	userRepo := user.NewUserRepo(db)
	userController := userCon.NewUsersControllers(userRepo)
	shoppingCartRepo := shoppingcart.NewShoppingCartRepo(db)
	shoppingCartController := shoppingCartCon.NewShoppingCartControllers(shoppingCartRepo)
	categoryRepo := category.NewCategoryRepo(db)
	categoryController := categoryCon.NewCategoryControllers(categoryRepo)
	authRepo := auth.NewAuthRepo(db)
	authController := authCon.NewAuthControllers(authRepo)
	e := echo.New()
	routes.RegisterPath(e, userController, productController, shoppingCartController, categoryController, authController)

	log.Fatal(e.Start(fmt.Sprintf(":%d", config.Port)))

}
