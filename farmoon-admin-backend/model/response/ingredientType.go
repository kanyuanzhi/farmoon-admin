package response

import "github.com/kanyuanzhi/farmoon-admin/farmoon-admin-backend/model"

type CountIngredientTypes struct {
	Count int64 `json:"count"`
}

type ListIngredientTypes struct {
	IngredientTypes []model.IngredientTypeInfo `json:"ingredientTypes"`
}

type AddIngredientType struct {
	IngredientType model.IngredientTypeInfo `json:"ingredientType"`
}
