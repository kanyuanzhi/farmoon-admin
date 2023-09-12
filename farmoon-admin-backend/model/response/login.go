package response

import "github.com/kanyuanzhi/farmoon-admin/farmoon-admin-backend/model"

type Login struct {
	model.UserInfo
	Token string `json:"token"`
}
