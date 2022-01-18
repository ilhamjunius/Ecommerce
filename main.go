package main

import (
	"ecommerce/configs"
	"ecommerce/utils"
	"fmt"
)

func main() {
	config := configs.GetConfig()

	db := utils.InitDB(config)
	fmt.Println(db)
}
