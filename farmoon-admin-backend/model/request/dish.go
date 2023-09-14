package request

type CountDishes struct {
	Filter              string `json:"filter" form:"filter"`
	EnableCuisineFilter bool   `json:"enableCuisineFilter" form:"enableCuisineFilter"`
	CuisineFilter       string `json:"cuisineFilter" form:"cuisineFilter"`
}

type ListDishes struct {
	PageIndex int `json:"pageIndex" form:"pageIndex"`
	PageSize  int `json:"pageSize" form:"pageSize"`
	CountDishes
}

type UpdateDish struct {
	Id      uint   `json:"id"`
	Name    string `json:"name"`
	Cuisine uint   `json:"cuisine"`
}

type DeleteDishes struct {
	Ids []uint `json:"ids"`
}
