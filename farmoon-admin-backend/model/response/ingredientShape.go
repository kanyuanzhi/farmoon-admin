package response

import "github.com/kanyuanzhi/farmoon-admin/farmoon-admin-backend/model"

type CountIngredientShapes struct {
	Count int64 `json:"count"`
}

type ListIngredientShapes struct {
	IngredientShapes []model.IngredientShapeInfo `json:"ingredientShapes"`
}

type AddIngredientShape struct {
	IngredientShape model.IngredientShapeInfo `json:"ingredientShape"`
}
