package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"releasepageback/services"
)

func Releases(ctx *gin.Context) {
	tp := ctx.Query("type") // 使用类型断言将接口类型转换为string类型
	//if !ok {
	//	fmt.Println("key type doesn't exist")
	//	// 如果转换失败，处理错误或者返回适当的响应给客户端
	//	ctx.JSON(http.StatusBadRequest, gin.H{"error": "Type assertion failed"})
	//	return
	//}
	// 使用转换后的字符串进行业务逻辑处理
	res := services.GetRelease(tp)
	ctx.JSON(http.StatusOK, res)
}
