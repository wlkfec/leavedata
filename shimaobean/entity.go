package shimaobean

import (
	"github.com/google/uuid"
	"time"
)

type PostCustomerDataDTO struct {
	Head       string
	CreateTime time.Time
	Data       PostCustomerDataInnerDTO
	MerchantId string
}

func CreateShiMaoPushData() *PostCustomerDataDTO {
	return &PostCustomerDataDTO{
		Head:       "getFiledRecord",
		CreateTime: time.Now(),
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
	RingDuration  int
	CustomerState int
	Uuid          string
	CallTime      string
	Duration      int
	Phone         string
	CallStatus    int
	FileName      string
	Name          string
	Sex           string
	Connect       bool
	SalerName     string
	SalerId       string
	IsExist       string
	Note          string
	Location      string
	ChannelType   string
	Account       string
	AccountName   string
	Tags          string
	Intent        IntentValue
}

type IntentValue struct {
	Level string
	Name  string
	Id    string
}
