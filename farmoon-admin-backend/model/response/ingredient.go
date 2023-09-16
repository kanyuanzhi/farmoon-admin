package response

import "github.com/kanyuanzhi/farmoon-admin/farmoon-admin-backend/model"

type CountIngredients struct {
	Count int64 `json:"count"`
}

type ListIngredients struct {
	Ingredients []model.IngredientInfo `json:"ingredients"`
}

type AddIngredient struct {
	Ingredient model.IngredientInfo `json:"ingredient"`
}
