package services

import (
	"releasepageback/model"
)

func GetRelease(tp string) []model.Release {
	release := model.Release{}
	res := release.GetreleaseByType(tp)

	return res
}
