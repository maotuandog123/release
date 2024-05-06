package services

import (
	"github.com/gin-gonic/gin"
	. "releasepageback/model"
	. "releasepageback/utils"
)

/**
创建用户
*/

func CreateUser(ctx *gin.Context) (uint, error) {
	user := User{}
	var err error
	if err = ctx.ShouldBindJSON(&user); err != nil {
		return 0, err
	}

	user.HashPassword, err = MakeHashPassword(user.Password)
	if err != nil {
		return 0, err
	}

	return user.Insert()
}

/**
获取指定ID的用户
*/

func GetUserById(id uint) (user User) {
	user = User{}
	user.GetOneById(id)
	return user
}

/**
根据账户信息(用户名、手机、邮箱)获取用户
*/

func GetUserByAccount(account string) (user User) {
	user = User{}
	user.GetOneByAccount(account)
	return user
}

/**
获取所有用户
*/

func GetAllUser() []User {
	user := User{}
	return user.GetAll()
}

/**
更新密码
*/

func ChangeUserPassword(user User, RawPassword string) {
	password, _ := MakeHashPassword(RawPassword)
	user.ChangePassword(password)
}
