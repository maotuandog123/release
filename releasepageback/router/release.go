package router

import (
	"github.com/gin-gonic/gin"
	"releasepageback/api"
	"releasepageback/utils"
)

func InitReleaseRouter(Router *gin.RouterGroup) {
	ReleaseRouter := Router.Group("release")
	{
		utils.Register(ReleaseRouter, []string{"POST", "GET"}, "info", api.Releases)
		utils.Register(ReleaseRouter, []string{"POST", "GET"}, "dfile", api.GetallDowloadfiles)
		utils.Register(ReleaseRouter, []string{"POST", "GET"}, "conninfo", api.GetallConn)
	}
}
