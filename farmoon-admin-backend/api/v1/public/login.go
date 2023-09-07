package public

import (
	"github.com/gin-gonic/gin"
	"github.com/kanyuanzhi/farmoon-admin/farmoon-admin-backend/global"
	"github.com/kanyuanzhi/farmoon-admin/farmoon-admin-backend/model"
	"github.com/kanyuanzhi/farmoon-admin/farmoon-admin-backend/model/request"
	"github.com/kanyuanzhi/farmoon-admin/farmoon-admin-backend/model/response"
	"github.com/kanyuanzhi/farmoon-admin/farmoon-admin-backend/utils"
)

type LoginApi struct{}

func (api *LoginApi) Login(c *gin.Context) {
	var loginRequest request.Login
	if err := request.ShouldBindJSON(c, &loginRequest); err != nil {
		return
	}

	var user model.SysUser
	if result := global.FXDb.Where("username = ?", loginRequest.Username).First(&user); result.RowsAffected == 0 {
		response.ErrorMessage(c, "用户名或密码错误")
		return
	}

	if !utils.CompareBcrypt(user.Password, loginRequest.Password) {
		response.ErrorMessage(c, "用户名或密码错误")
		return
	}

	loginResponse := response.Login{
		Avatar:   user.Avatar,
		Username: user.Username,
		Nickname: user.Nickname,
		RealName: user.RealName,
		Token:    utils.CreateToken(user.Username),
	}

	response.SuccessMessageData(c, loginResponse, "登录成功")
}
