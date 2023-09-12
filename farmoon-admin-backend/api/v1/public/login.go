package public

import (
	"encoding/base64"
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
		response.ErrorMessage(c, err.Error())
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

	if user.Roles == nil {
		user.Roles = []string{}
	}

	loginResponse := response.Login{
		UserInfo: model.UserInfo{
			Id:       user.Id,
			Username: user.Username,
			Nickname: user.Nickname,
			RealName: user.RealName,
			Avatar:   "data:image/png;base64," + base64.StdEncoding.EncodeToString(user.Avatar),
			Gender:   user.Gender,
			Mobile:   user.Mobile,
			Email:    user.Email,
			Roles:    user.Roles,
		},
		Token: utils.CreateToken(user.Username),
	}

	response.SuccessMessageData(c, loginResponse, "登录成功")
}
