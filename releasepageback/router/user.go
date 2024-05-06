package router

import (
	"github.com/gin-gonic/gin"
	"releasepageback/utils"

	"releasepageback/api"
)

func InitUserRouter(Router *gin.RouterGroup) {
	/**
	用户相关的路由组
	*/
	UserRouter := Router.Group("user")

	{
		// 用户认证登陆
		utils.Register(UserRouter, []string{"GET", "POST"}, "authenticate", api.UserAuthenticate)
		// 用户创建
		utils.Register(UserRouter, []string{"POST"}, "", api.UserCreate)
	}

}
