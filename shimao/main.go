package main

import (
	_ "github.com/go-sql-driver/mysql"
	"makeuse/basebean"
)

var (
	fileBytes  []byte
	err        error
	rootResult basebean.Root
	myConfig   *basebean.Config
)

func main() {

	//// 世贸数据模板
	//_ = basebean.BuildRootByFile()
	//
	//// 所有的通话记录
	//talkDetails := basebean.GetTalkDetailFromDB()
	//
	//// 项目和项目guid的对应关系
	//getProjectInfoFromDB := talkDetails.GetProjectInfoFromDB()
	//fmt.Println(getProjectInfoFromDB)
	//
	//// 依次构建推送字符串
	//for _, ele := range talkDetails{
	//	_ = ele.BuildMqPushData(getProjectInfoFromDB[ele.ProjectId])
	//}

}

//myConfig = new(basebean.Config)
//if _, err = toml.DecodeFile("config.toml", myConfig); err != nil {
//fmt.Println(err)
//return
//}
//fmt.Println(myConfig.Mysql.Username)
