package initialize

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"releasepageback/middleware"
	"releasepageback/router"
)

func InitRouter() *gin.Engine {
	// 1. 创建路由
	Router := gin.Default()
	Router.Use(middleware.ExceptionMiddleware)
	Router.Use(middleware.GinLogger(), middleware.GinRecovery(true))

	// 注册中间件
	Router.Use(middleware.Cors(), middleware.GinLogger(), middleware.GinRecovery(true))
	// 2. Api路由分组
	ApiGroup := Router.Group("/api")
	// 3. 初始化用户相关路由
	router.InitReleaseRouter(ApiGroup)
	zap.S().Info("路由初始化完成....")
	return Router
}
