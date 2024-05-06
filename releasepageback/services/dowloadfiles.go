package services

import (
	"releasepageback/model"
)

func GetallDownloadfiles() []model.Downloadfiles {
	df := model.Downloadfiles{}
	dfs := df.GetAll()

	return dfs

}
