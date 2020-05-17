package main

import (
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Name    string `json:"name"`
	Address string `json:"address"`
}

func main() {
	//var nowTime = time.Now()
	//fmt.Println(util.GetTimeStr(nowTime))
	//excel.ReadExcel()
	//	2020-04-27 00:26:21.000
	//  2020-04-15 15:41:33 +0800 CST
	//talkDetails := entity.GetTalkDetailsFromDB()
	//for _, talk := range talkDetails{
	//	fmt.Println(talk.TalkDate)
	//}
	//getProjectInfoFromDB := talkDetails.GetProjectInfoFromDB()
	//talkDetail := talkDetails[48]
	//_ = talkDetail.BuildMqPushData(getProjectInfoFromDB[talkDetail.ProjectId])
}
