package model

import (
	"fmt"
	. "releasepageback/database"
)

type Release struct {
	Id           int    `gorm:"primaryKey;column:id" json:"id"`
	X86Name      string `gorm:"column:x86name" json:"x86name"`
	X86Path      string `gorm:"column:x86path" json:"x86path"`
	Type         string `gorm:"column:type" json:"type"`
	Version      string `gorm:"column:version" json:"version"`
	Md5x86Path   string `gorm:"column:md5x86path" json:"md5x86path"`
	Md5x86name   string `gorm:"column:md5x86name" json:"md5x86name"`
	Comment      string `gorm:"column:comment" json:"comment"`
	ArrchName    string `gorm:"column:arrchname" json:"arrchname"`
	ArrchPath    string `gorm:"column:arrchpath" json:"arrchpath"`
	Md5arrchPath string `gorm:"column:md5arrchpath" json:"md5arrchpath"`
	Md5arrchname string `gorm:"column:md5arrchname" json:"md5arrchname"`
}

func (release Release) GetreleaseByType(tp string) []Release {
	re := []Release{}
	Orm.Find(&re, "type = ?", tp)
	fmt.Println(re)
	return re
}

func (Release) TableName() string {
	return "releaseinfo"
}
