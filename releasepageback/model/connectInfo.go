package model

import (
	. "releasepageback/database"
)

type ConnectInfo struct {
	Id      int    `json:"id"`
	Envname string `json:"envname"`
	Type    string `json:"type"`
	Account string `json:"account"`
	Passwd  string `json:"passwd"`
}

func (conn ConnectInfo) GetAll() (conns []ConnectInfo) {

	Orm.Find(&conns)
	return conns
}

// 设置模型对应的表名
func (ConnectInfo) TableName() string {
	return "buildimageconninfo"
}
