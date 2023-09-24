package rpcResponse

import (
	"github.com/google/uuid"
	"github.com/kanyuanzhi/farmoon-admin/farmoon-admin-backend/model"
)

type FetchOfficialDishes struct {
	NewAddedDishes   []model.SysDish `json:"newAddedDishes"`
	UpdatedDish      []model.SysDish `json:"updatedDish"`
	DeletedDishUUIDs []uuid.UUID     `json:"deletedDishUUIDs"`
}
