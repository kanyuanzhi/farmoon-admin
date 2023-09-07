package request

import (
	"github.com/gin-gonic/gin"
	"github.com/kanyuanzhi/farmoon-admin/farmoon-admin-backend/model/response"
)

func ShouldBindJSON(c *gin.Context, obj interface{}) error {
	if err := c.ShouldBindJSON(obj); err != nil {
		//global.GqaLogger.Error(global.GqaConfig.System.BindError, zap.Any("err", err))
		response.ErrorMessage(c, err.Error())
		return err
	}
	return nil
}
