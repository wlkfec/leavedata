package main

import (
	_ "github.com/go-sql-driver/mysql"
	"makeuse/entity"
)

type User struct {
	Name    string `json:"name"`
	Address string `json:"address"`
}

func main() {
	talkDetails := entity.GetTalkDetailsFromDB()

	getProjectInfoFromDB := talkDetails.GetProjectInfoFromDB()

	talkDetail := talkDetails[48]

	_ = talkDetail.BuildMqPushData(getProjectInfoFromDB[talkDetail.ProjectId])
}
