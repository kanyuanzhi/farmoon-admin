package model

import "github.com/kanyuanzhi/farmoon-admin/farmoon-admin-backend/global"

type SysUser struct {
	global.FXModel
	Username string `json:"username" gorm:"comment:用户名;not null;uniqueIndex;"`
	Nickname string `json:"nickname" gorm:"comment:用户昵称;"`
	RealName string `json:"realName" gorm:"comment:真实姓名;"`
	Password string `json:"-"  gorm:"comment:用户密码;"`
	Avatar   string `json:"avatar" gorm:"comment:头像"`
	Gender   string `json:"gender" gorm:"comment:性别;default:gender_unknown;"`
	Mobile   string `json:"mobile" gorm:"comment:手机号;"`
	Email    string `json:"email" gorm:"comment:邮箱;"`
}
