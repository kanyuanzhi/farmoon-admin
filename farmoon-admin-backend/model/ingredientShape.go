package model

import "github.com/kanyuanzhi/farmoon-admin/farmoon-admin-backend/global"

type SysIngredientShape struct {
	global.FXModel
	Name string `json:"name" gorm:"comment:食材形状;"`
}

type IngredientShapeInfo struct {
	Id   uint   `json:"id"`
	Sort uint   `json:"sort"`
	Name string `json:"name"`
}
