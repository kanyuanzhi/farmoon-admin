package global

import (
	"github.com/kanyuanzhi/farmoon-admin/farmoon-admin-backend/config"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	FXConfig config.Config
	FXDb     *gorm.DB
	FXLogger *zap.Logger
)

type FXModel struct {
	Id        uint   `json:"id" gorm:"comment:id;primaryKey;autoIncrement;"`
	CreatedAt int64  `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt int64  `json:"updated_at" gorm:"autoUpdateTime"`
	Sort      uint   `json:"sort" gorm:"comment:排序;default:1;"`
	Memo      string `json:"memo" gorm:"comment:备注描述;type:text;"`
}
