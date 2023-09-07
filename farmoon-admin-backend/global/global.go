package global

import (
	"github.com/kanyuanzhi/farmoon-admin/farmoon-admin-backend/config"
	"gorm.io/gorm"
	"time"
)

var (
	FXConfig config.Config
	FXDb     *gorm.DB
)

type FXModel struct {
	Id        uint      `json:"id" gorm:"comment:id;primarykey;autoIncrement;"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Sort      uint      `json:"sort" gorm:"comment:排序;default:1;"`
	Memo      string    `json:"memo" gorm:"comment:备注描述;type:text;"`
}
