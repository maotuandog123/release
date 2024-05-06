package api

import (
	"github.com/gin-gonic/gin"
	"releasepageback/services"
)

func GetallDowloadfiles(ctx *gin.Context) {
	dfs := services.GetallDownloadfiles()
	ctx.JSON(200, dfs)
}
