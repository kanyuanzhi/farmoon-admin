package public

import (
	"github.com/gin-gonic/gin"
	"github.com/kanyuanzhi/farmoon-admin/farmoon-admin-backend/global"
	"github.com/kanyuanzhi/farmoon-admin/farmoon-admin-backend/model"
	"github.com/kanyuanzhi/farmoon-admin/farmoon-admin-backend/model/request"
	"github.com/kanyuanzhi/farmoon-admin/farmoon-admin-backend/model/response"
	"github.com/kanyuanzhi/farmoon-admin/farmoon-admin-backend/utils"
	"github.com/lib/pq"
)

type RegisterApi struct{}

func (api *RegisterApi) Register(c *gin.Context) {
	var registerRequest request.Register
	if err := request.ShouldBindJSON(c, &registerRequest); err != nil {
		response.ErrorMessage(c, err.Error())
		return
	}
	if registerRequest.RepeatPassword != registerRequest.Password {
		response.ErrorMessage(c, "两次输入的密码不一致")
		return
	}
	if result := global.FXDb.Where("username = ?", registerRequest.Username).First(&model.SysUser{}); result.RowsAffected == 1 {
		response.ErrorMessage(c, "用户名已存在")
		return
	}
	passwordEncoded, _ := utils.EncodeBcrypt(registerRequest.Password)

	imageData, err := utils.LoadLocalImage()
	if err != nil {
		response.ErrorMessage(c, err.Error())
		return
	}

	user := model.SysUser{
		Username: registerRequest.Username,
		Password: passwordEncoded,
		Avatar:   imageData,
		Gender:   "N", // 未知性别
		Roles:    make(pq.StringArray, 0),
	}

	if result := global.FXDb.Create(&user); result.Error != nil {
		response.ErrorMessage(c, result.Error.Error())
		return
	}

	response.SuccessMessage(c, "注册成功")
}
