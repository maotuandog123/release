package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"releasepageback/constants"
	. "releasepageback/database"
	. "releasepageback/services"
)

/**
用户认证登陆
*/

func UserAuthenticate(ctx *gin.Context) {
	/**
	用户认证登陆
	*/
	RedisConn := RedisPool1.Get()
	_, _ = RedisConn.Do("SET", "age", 19)
	ctx.String(http.StatusOK, "hello Urils!")
}

/**
创建用户
*/

func UserCreate(ctx *gin.Context) {

	id, err := CreateUser(ctx)
	if err != nil || id < 1 {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    constants.CodeCreateUserFail,
			"message": constants.CreateUserFail,
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"code":    constants.CodeSuccess,
		"message": constants.CreateUserSuccess,
		"data":    id,
	})
}
