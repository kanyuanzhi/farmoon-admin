package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/kanyuanzhi/farmoon-admin/farmoon-admin-backend/api/v1/public"
)

func InitPublicRouter(routerGroup *gin.RouterGroup) {
	loginApi := public.LoginApi{}
	registerApi := public.RegisterApi{}
	routerGroup.POST("login", loginApi.Login)
	routerGroup.POST("register", registerApi.Register)
}
