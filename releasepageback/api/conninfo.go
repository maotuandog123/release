package api

import (
	"github.com/gin-gonic/gin"
	"releasepageback/services"
)

func GetallConn(ctx *gin.Context) {
	conn := services.GetallConninfo()
	ctx.JSON(200, conn)
}
