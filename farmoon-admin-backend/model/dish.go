package model

import (
	"github.com/google/uuid"
	"github.com/kanyuanzhi/farmoon-admin/farmoon-admin-backend/global"
)

type SysDish struct {
	global.FXModel
	Name        string                              `json:"name" gorm:"comment:菜品名称;"`
	UUID        uuid.UUID                           `json:"uuid" gorm:"comment:唯一标识符;not null;uniqueIndex;"`
	Steps       []map[string]interface{}            `json:"steps" gorm:"json;comment:步骤;type:json"`
	CustomSteps map[string][]map[string]interface{} `json:"customSteps" gorm:"json;comment:口味步骤;type:json"`
	Image       []byte                              `json:"image" gorm:"comment:菜品图像;"`
	Cuisine     uint                                `json:"cuisine" gorm:"comment:菜系;default:1"`
}

type DishInfo struct {
	Id          uint                                `json:"id"`
	Name        string                              `json:"name"`
	UUID        uuid.UUID                           `json:"uuid"`
	Steps       []map[string]interface{}            `json:"steps"`
	CustomSteps map[string][]map[string]interface{} `json:"customSteps"`
	Image       string                              `json:"image"`
	Cuisine     uint                                `json:"cuisine"`
}
