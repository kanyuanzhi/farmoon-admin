package request

type CountDishes struct {
	Filter              string `json:"filter" form:"filter"`
	EnableCuisineFilter bool   `json:"enableCuisineFilter" form:"enableCuisineFilter"`
	CuisineFilter       string `json:"cuisineFilter" form:"cuisineFilter"`
	EnableOwnerFilter   bool   `json:"enableOwnerFilter" form:"enableOwnerFilter"`
	OwnerFilter         string `json:"ownerFilter" form:"ownerFilter"`
	IsOfficial          bool   `json:"isOfficial" form:"isOfficial"`
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

type UpdateDishWithSteps struct {
	UpdateDish
	Steps []map[string]interface{} `json:"steps"`
}

type DeleteDishes struct {
	Ids []uint `json:"ids"`
}

type AddDish struct {
	Name    string                   `json:"name"`
	Cuisine uint                     `json:"cuisine"`
	Steps   []map[string]interface{} `json:"steps"`
}

type CopyDish struct {
	Id uint `json:"id"`
}

type ToppingDish struct {
	Id uint `json:"id"`
}

type ExportSteps struct {
	Ids []uint `json:"ids"`
}

type ListOwners struct {
	PageIndex int `json:"pageIndex" form:"pageIndex"`
	PageSize  int `json:"pageSize" form:"pageSize"`
}

type AddToOfficials struct {
	Id uint `json:"id"`
}

type GetDishQrCode struct {
	Id uint `json:"id" form:"id"`
}
