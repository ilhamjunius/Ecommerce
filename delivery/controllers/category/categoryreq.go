package category

type CreateCategoryRequestFormat struct {
	ID           uint   `json:"id" form:"id"`
	CategoryType string `json:"category_type" form:"category_type"`
}

type PutCategoryRequestFormat struct {
	ID           uint   `json:"id" form:"id"`
	CategoryType string `json:"category_type" form:"category_type"`
}
