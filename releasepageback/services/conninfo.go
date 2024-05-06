package services

import (
	"fmt"
	"releasepageback/model"
)

func GetallConninfo() []model.ConnectInfo {
	conn := model.ConnectInfo{}
	conns := conn.GetAll()
	fmt.Println(conns)
	return conns

}
