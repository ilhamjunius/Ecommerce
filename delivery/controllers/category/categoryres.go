package category

import "ecommerce/entities"

type CreateCategoryResponseFormat struct {
	Message string              `json:"message"`
	Data    []entities.Category `json:"data"`
}

type GetCategoryResponseFormat struct {
	Message interface{}         `json:"message"`
	Data    []entities.Category `json:"data"`
}
type GetCategoriesResponseFormat struct {
	Message string              `json:"message"`
	Data    []entities.Category `json:"data"`
}
type DeleteCategoryResponseFormat struct {
	Message string `json:"message"`
}

type PutCategoryResponseFormat struct {
	Message string              `json:"message"`
	Data    []entities.Category `json:"data"`
}
