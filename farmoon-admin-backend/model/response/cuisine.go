package response

import "github.com/kanyuanzhi/farmoon-admin/farmoon-admin-backend/model"

type ListCuisines struct {
	Cuisines []model.CuisineInfo `json:"cuisines"`
}

type AddCuisine struct {
	Cuisine model.CuisineInfo `json:"cuisine"`
}
