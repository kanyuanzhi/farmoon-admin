package boot

import (
	"github.com/gin-gonic/gin"
	"github.com/kanyuanzhi/farmoon-admin/farmoon-admin-backend/middleware"
	"github.com/kanyuanzhi/farmoon-admin/farmoon-admin-backend/router/v1"
)

func Router() *gin.Engine {
	router := gin.Default()
	gin.SetMode(gin.ReleaseMode)
	router.Use(middleware.Cors())

	apiV1 := router.Group("/farmoon-api/v1")

	// 不用验证的路由
	publicGroup := apiV1.Group("/public")
	v1.InitPublicRouter(publicGroup)

	// 需要验证的路由
	privateGroup := apiV1.Group("/private")
	v1.InitPrivateRouter(privateGroup)

	return router
}
