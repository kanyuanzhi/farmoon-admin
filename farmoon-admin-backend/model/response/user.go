package response

import "github.com/kanyuanzhi/farmoon-admin/farmoon-admin-backend/model"

type CountUsers struct {
	Count int64 `json:"count"`
}

type ListUsers struct {
	Users []model.UserInfo `json:"users"`
	Count int64            `json:"count"`
}

type UpdateUser struct {
}

type UpdatePassword struct {
}

type AddUser struct {
	model.UserInfo
}
