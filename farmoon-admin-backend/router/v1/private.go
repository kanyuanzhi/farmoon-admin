package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/kanyuanzhi/farmoon-admin/farmoon-admin-backend/api/v1/private"
)

func InitPrivateRouter(routerGroup *gin.RouterGroup) {
	userApi := &private.UserApi{}
	routerGroup.GET("user/list", userApi.List)
	routerGroup.GET("user/count", userApi.Count)
	routerGroup.PUT("user/update", userApi.Update)
	routerGroup.DELETE("user/delete", userApi.Delete)
	routerGroup.PUT("user/update-roles", userApi.UpdateRoles)
	routerGroup.PUT("user/update-password", userApi.UpdatePassword)
	routerGroup.POST("user/update-avatar", userApi.UpdateAvatar)
	routerGroup.POST("user/add", userApi.Add)

	cuisineApi := &private.CuisineApi{}
	routerGroup.GET("cuisine/list", cuisineApi.List)
	routerGroup.PUT("cuisine/update-name", cuisineApi.UpdateName)
	routerGroup.PUT("cuisine/update-unDeletable", cuisineApi.UpdateUnDeletable)
	routerGroup.PUT("cuisine/update-sort", cuisineApi.UpdateSort)
	routerGroup.DELETE("cuisine/delete", cuisineApi.Delete)
	routerGroup.POST("cuisine/add", cuisineApi.Add)

	dishApi := &private.DishApi{}
	routerGroup.GET("dish/list", dishApi.List)
	routerGroup.GET("dish/count", dishApi.Count)
	routerGroup.PUT("dish/update", dishApi.Update)
	routerGroup.DELETE("dish/delete", dishApi.Delete)
	routerGroup.POST("dish/update-image", dishApi.UpdateImage)
}
