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

type AddDish struct {
	Dish model.DishInfo `json:"dish"`
}

type UpdateDishWithSteps struct {
	Dish model.DishInfo `json:"dish"`
}

type CopyDish struct {
	Dish model.DishInfo `json:"dish"`
}

type ExportSteps struct {
	Dishes []model.DishInfo `json:"dishes"`
}

type ListOwners struct {
	Owners []string `json:"owners"`
}

type GetQrCode struct {
	QrCode []byte `json:"qrCode"`
}
