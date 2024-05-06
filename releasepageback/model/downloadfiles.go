package model

import . "releasepageback/database"

type Downloadfiles struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Filename string `json:"filename"`
	Path     string `json:"path"`
}

func (Downloadfiles) TableName() string {
	return "commentfiles"
}
func (f Downloadfiles) GetAll() (dfs []Downloadfiles) {

	Orm.Find(&dfs)
	return dfs
}
