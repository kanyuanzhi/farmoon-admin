package response

import "github.com/kanyuanzhi/farmoon-admin/farmoon-admin-backend/model"

type CountDishes struct {
	Count int64 `json:"count"`
}

type ListDishes struct {
	Dishes []model.DishInfo `json:"dishes"`
}

type UpdateDishImage struct {
	Image string `json:"image"`
}
