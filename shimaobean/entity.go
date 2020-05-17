package shimaobean

import (
	"github.com/google/uuid"
	"time"
)

type PostCustomerDataDTO struct {
	Head       string                   `json:"head"`
	CreateTime int64                    `json:"createTime"`
	Data       PostCustomerDataInnerDTO `json:"data"`
	MerchantId string                   `json:"merchantId"`
}

func CreateShiMaoPushData() *PostCustomerDataDTO {
	return &PostCustomerDataDTO{
		Head:       "getFiledRecord",
		CreateTime: time.Now().UnixNano() / 1e6,
		MerchantId: "808FD400684443C88E02B9A8529B5F39",
		Data: PostCustomerDataInnerDTO{
			RingDuration:  10,
			CustomerState: 1,
			Uuid:          uuid.New().String(),
			CallTime:      "1990-01-01 12:12:12",
			Duration:      100,
			Phone:         "18666668888",
			CallStatus:    0,
			FileName:      "http://ig-ivr.oss-cn-shanghai.aliyuncs.com/smartphone3cf1e7ddf079813191f51066e3b51c9f.wav",
			Name:          "王力宏",
			Sex:           "MALE",
			Connect:       false,
			SalerName:     "邱爱芳",
			SalerId:       "AIFANG_QIU",
			IsExist:       "true",
			Note:          "asdasasdasdasdasasdasdasd",
			Location:      "北京",
			ChannelType:   "已分配来电",
			Account:       "xxx",
			AccountName:   "asdasdasdasd",
			Tags:          "dd",
			Intent:        IntentValue{},
		},
	}
}

type PostCustomerDataInnerDTO struct {
	RingDuration  int         `json:"ringDuration"`
	CustomerState int         `json:"customerState"`
	Uuid          string      `json:"uuid"`
	CallTime      string      `json:"callTime"`
	Duration      int         `json:"duration"`
	Phone         string      `json:"phone"`
	CallStatus    int         `json:"callStatus"`
	FileName      string      `json:"fileName"`
	Name          string      `json:"name"`
	Sex           string      `json:"sex"`
	Connect       bool        `json:"connect"`
	SalerName     string      `json:"salerName"`
	SalerId       string      `json:"salerId"`
	IsExist       string      `json:"isExist"`
	Note          string      `json:"note"`
	Location      string      `json:"location"`
	ChannelType   string      `json:"channelType"`
	Account       string      `json:"account"`
	AccountName   string      `json:"accountName"`
	Tags          string      `json:"tags"`
	Intent        IntentValue `json:"intent"`
}

type IntentValue struct {
	Level string `json:"level"`
	Name  string `json:"name"`
	Id    string `json:"id"`
}
