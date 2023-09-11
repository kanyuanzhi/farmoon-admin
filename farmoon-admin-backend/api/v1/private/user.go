package private

import (
	"encoding/base64"
	"github.com/gin-gonic/gin"
	"github.com/kanyuanzhi/farmoon-admin/farmoon-admin-backend/global"
	"github.com/kanyuanzhi/farmoon-admin/farmoon-admin-backend/model"
	"github.com/kanyuanzhi/farmoon-admin/farmoon-admin-backend/model/request"
	"github.com/kanyuanzhi/farmoon-admin/farmoon-admin-backend/model/response"
	"github.com/kanyuanzhi/farmoon-admin/farmoon-admin-backend/utils"
	"strings"
)

type UserApi struct{}

func (api *UserApi) Count(c *gin.Context) {
	var countUsersRequest request.CountUsers
	if err := request.ShouldBindQuery(c, &countUsersRequest); err != nil {
		response.ErrorMessage(c, err.Error())
		return
	}

	filterDb, err := request.GenerateUserQueryCondition(countUsersRequest.Filter, countUsersRequest.EnableGenderFilter,
		countUsersRequest.GenderFilter, countUsersRequest.EnableRolesFilter, strings.Split(countUsersRequest.RolesFilter, ","))
	if err != nil {
		response.ErrorMessage(c, err.Error())
		return
	}

	var count int64
	if err := filterDb.Count(&count).Error; err != nil {
		response.ErrorMessage(c, err.Error())
		return
	}

	countUserResponse := response.CountUsers{
		Count: count,
	}

	response.SuccessData(c, countUserResponse)
}

func (api *UserApi) List(c *gin.Context) {
	var listUsersRequest request.ListUsers
	if err := request.ShouldBindQuery(c, &listUsersRequest); err != nil {
		response.ErrorMessage(c, err.Error())
		return
	}

	filterDb, err := request.GenerateUserQueryCondition(listUsersRequest.Filter, listUsersRequest.EnableGenderFilter,
		listUsersRequest.GenderFilter, listUsersRequest.EnableRolesFilter, strings.Split(listUsersRequest.RolesFilter, ","))
	if err != nil {
		response.ErrorMessage(c, err.Error())
		return
	}

	var users []model.SysUser
	if err := filterDb.Limit(listUsersRequest.PageSize).Offset((listUsersRequest.PageIndex - 1) * listUsersRequest.PageSize).
		Order("id").Find(&users).Error; err != nil {
		response.ErrorMessage(c, err.Error())
		return
	}

	usersInfo := []model.UserInfo{}
	for _, user := range users {
		usersInfo = append(usersInfo, model.UserInfo{
			Id:       user.Id,
			Avatar:   base64.StdEncoding.EncodeToString(user.Avatar),
			Username: user.Username,
			Nickname: user.Nickname,
			RealName: user.RealName,
			Gender:   user.Gender,
			Mobile:   user.Mobile,
			Email:    user.Email,
			Roles:    user.Roles,
		})
	}

	listUsersResponse := response.ListUsers{
		Users: usersInfo,
	}

	response.SuccessData(c, listUsersResponse)
}

func (api *UserApi) Update(c *gin.Context) {
	var updateUserRequest request.UpdateUser
	if err := request.ShouldBindJSON(c, &updateUserRequest); err != nil {
		response.ErrorMessage(c, err.Error())
		return
	}

	user := model.SysUser{
		FXModel: global.FXModel{
			Id: updateUserRequest.Id,
		},
		Nickname: updateUserRequest.Nickname,
		RealName: updateUserRequest.RealName,
		Gender:   updateUserRequest.Gender,
		Mobile:   updateUserRequest.Mobile,
		Email:    updateUserRequest.Email,
		Roles:    updateUserRequest.Roles,
	}

	if err := global.FXDb.Model(&user).Omit("id").Updates(user).Error; err != nil {
		response.ErrorMessage(c, err.Error())
		return
	}

	response.SuccessMessage(c, "更新成功")
}

func (api *UserApi) Delete(c *gin.Context) {
	var deleteUserRequest request.DeleteUsers
	if err := request.ShouldBindJSON(c, &deleteUserRequest); err != nil {
		response.ErrorMessage(c, err.Error())
		return
	}

	var users []model.SysUser
	for _, id := range deleteUserRequest.Ids {
		users = append(users, model.SysUser{
			FXModel: global.FXModel{
				Id: id,
			},
		})
	}

	if err := global.FXDb.Delete(&users).Error; err != nil {
		response.ErrorMessage(c, err.Error())
		return
	}

	response.SuccessMessage(c, "删除成功")
}

func (api *UserApi) UpdatePassword(c *gin.Context) {
	var updatePasswordRequest request.UpdatePassword
	if err := request.ShouldBindJSON(c, &updatePasswordRequest); err != nil {
		response.ErrorMessage(c, err.Error())
		return
	}

	passwordEncoded, _ := utils.EncodeBcrypt(updatePasswordRequest.Password)
	user := model.SysUser{
		FXModel:  global.FXModel{Id: updatePasswordRequest.Id},
		Password: passwordEncoded,
	}

	if err := global.FXDb.Model(&user).Update("password", user.Password).Error; err != nil {
		response.ErrorMessage(c, err.Error())
		return
	}

	response.SuccessMessage(c, "更新密码成功")
}

func (api *UserApi) Add(c *gin.Context) {
	var addUserRequest request.AddUser
	if err := request.ShouldBindJSON(c, &addUserRequest); err != nil {
		response.ErrorMessage(c, err.Error())
		return
	}

	if result := global.FXDb.Where("username = ?", addUserRequest.Username).First(&model.SysUser{}); result.RowsAffected == 1 {
		response.ErrorMessage(c, "用户名已存在")
		return
	}

	passwordEncoded, _ := utils.EncodeBcrypt(addUserRequest.Password)

	imageData, err := utils.LoadDefaultProfilePhoto()
	if err != nil {
		response.ErrorMessage(c, err.Error())
		return
	}

	user := model.SysUser{
		Username: addUserRequest.Username,
		Nickname: addUserRequest.Nickname,
		RealName: addUserRequest.RealName,
		Password: passwordEncoded,
		Avatar:   imageData,
		Gender:   addUserRequest.Gender,
		Mobile:   addUserRequest.Mobile,
		Email:    addUserRequest.Email,
		Roles:    addUserRequest.Roles,
	}

	if err := global.FXDb.Create(&user).Error; err != nil {
		response.ErrorMessage(c, err.Error())
		return
	}

	addUserResponse := response.AddUser{
		UserInfo: model.UserInfo{
			Id:       user.Id,
			Username: user.Username,
			Nickname: user.Nickname,
			RealName: user.RealName,
			Avatar:   base64.StdEncoding.EncodeToString(user.Avatar),
			Gender:   user.Gender,
			Mobile:   user.Mobile,
			Email:    user.Email,
			Roles:    user.Roles,
		},
	}

	response.SuccessMessageData(c, addUserResponse, "添加用户成功")
}
