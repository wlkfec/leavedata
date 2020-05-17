package main

import (
	_ "github.com/go-sql-driver/mysql"
	"makeuse/entity"
)

func main() {

	entity.CreateKnowChannelTemp()

	//talkDetails := entity.GetTalkDetailsFromDB()
	//
	//getProjectInfoFromDB := talkDetails.GetProjectInfoFromDB()
	//
	//talkDetail := talkDetails[48]
	//
	//customerDetail := talkDetail.GetCustomerDetailByTalkDetail()
	//
	//customerDetail.GetCustomerTagByCustomerId()
	//
	//mqPushData := talkDetail.BuildMqPushData(getProjectInfoFromDB[talkDetail.ProjectId])
	//fmt.Println(" ..... ..... ..... ..... ..... ..... ..... ..... ")
	////marshal, _ := json.Marshal(knowChannelValues)
	//f, _ := os.OpenFile("demo.txt", os.O_CREATE|os.O_RDWR, 0660)
	//defer f.Close()
	//f.WriteString(mqPushData)
	//fmt.Println(" ..... ..... ..... ..... ..... ..... ..... ..... ")
}
