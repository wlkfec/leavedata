package entity

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/orm"
	"io/ioutil"
	"makeuse/config"
	"makeuse/constant"
	"makeuse/shimaobean"
	"strings"
	"time"
)

// ===== ===== ===== ===== ===== ===== < sm original data > ===== ===== ===== ===== ===== =====
type Root struct {
	UpdateFieldParameter FieldParameter
}
type FieldParameter struct {
	Head       string             `json:"head"`
	MerchantId string             `json:"merchantId"`
	CreateTime string             `json:"createTime"`
	Data       FieldParameterData `json:"data"`
}
type FieldParameterData struct {
	Intents Intents       `json:"intents"`
	Tags    []*TagsValues `json:"tags"`
}
type Intents struct {
	IsRequire bool            `json:"isRequire"`
	Values    []IntentsValues `json:"values"`
}
type IntentsValues struct {
	Level string `json:"level"`
	Name  string `json:"name"`
	Id    string `json:"id"`
}
type TagsValues struct {
	IsRequired bool        `json:"isRequired"`
	Values     interface{} `json:"values"`
	Name       string      `json:"name"`
	Index      int         `json:"index"`
	Modifiable bool        `json:"modifiable"`
	IsRadio    bool        `json:"isRadio"`
	Key        string      `json:"key"`
}
type TagsValuesInner struct {
	Name string `json:"name"`
}

// ===== ===== ===== ===== ===== ===== < end > ===== ===== ===== ===== ===== =====

// ===== ===== ===== ===== ===== ===== < DB original data > ===== ===== ===== ===== ===== =====
type CustomerDetail struct {
	Id               int
	UserName         string
	PhoneNumber      string
	CallPhoneNumber  string
	Gender           string
	AssignSaler      string
	SalerName        string
	Type             string
	TaskId           int
	CreateDate       time.Time
	CreateBy         string
	BuildingId       string
	ProjectId        string
	IsNotInput       int
	IsNotInputNumber int
	CustomerType     string
}
type CustomerTag struct {
	Id           int
	KnowChannel  string
	Attention    string
	Intent       string
	Consumer     string
	HouseType    string
	HouseTarget  string
	ConsumerType string
	IntentHouse  string
	Resistance   string
	BuyIntent    string
	Note         string
	CustomerId   int
}
type BaseOrganization struct {
	Id               string `orm:"colunm(id);pk"`
	OrganizationName string
	OrganizationType string
	ParentId         string
	Active           int
	Sort             int
	CreateDate       time.Time
	UpdateDate       time.Time
	ProjectGuid      string
}
type TalkDetail struct {
	Id                  int
	PhoneNumber         string
	PhoneName           string
	PhoneType           string
	TalkType            int
	TalkDate            time.Time
	TalkCostTime        int
	RingCostTime        int
	SalerId             string
	TalkContent         string
	TalkAudio           string
	TaskId              int
	BuildingId          string
	ProjectId           string
	PhoneOperator       string
	PhoneAttribution    string
	PhoneSystem         int
	NumberWashingResult int
	SensitiveWordCount  int
	Imei                string
}

// ===== ===== ===== ===== ===== ===== < end > ===== ===== ===== ===== ===== =====

// ===== ===== ===== ===== ===== ===== < type definition > ===== ===== ===== ===== ===== =====
type DBTalkData []TalkDetail

// ===== ===== ===== ===== ===== ===== < end > ===== ===== ===== ===== ===== =====

// ===== ===== ===== ===== ===== ===== < function > ===== ===== ===== ===== ===== =====

/**
解析原始数据格式模板
*/
func CreateRoot() *Root {
	var (
		fileBytes []byte
		err       error
		root      Root
	)
	if fileBytes, err = ioutil.ReadFile("all_data.json"); err != nil {
		fmt.Println(err)
		return nil
	}
	if err = json.Unmarshal(fileBytes, &root); err != nil {
		fmt.Println(err)
		return nil
	}
	return &root
}

type TagsValuesTemp struct {
	IsRequired bool        `json:"isRequired"`
	Values     interface{} `json:"values"`
	Name       string      `json:"name"`
	Index      int         `json:"index"`
	Modifiable bool        `json:"modifiable"`
	IsRadio    bool        `json:"isRadio"`
	Key        string      `json:"key"`
}

/**
获取用户标签
*/
func (customerDetail *CustomerDetail) GetCustomerTagByCustomerId() *CustomerTag {
	o := orm.NewOrm()
	org := CustomerTag{
		CustomerId: customerDetail.Id,
	}
	if err := o.Read(&org, "customer_id"); err != nil {
		fmt.Println(err)
		return nil
	}
	return &org
}

/**
根据通话记录获得对应的客户信息
*/
func (talkDetail *TalkDetail) GetCustomerDetailByTalkDetail() *CustomerDetail {
	o := orm.NewOrm()
	org := CustomerDetail{
		PhoneNumber: talkDetail.PhoneNumber,
		BuildingId:  talkDetail.BuildingId,
	}
	if err := o.Read(&org, "phone_number", "building_id"); err != nil {
		fmt.Println(err)
		return nil
	}
	return &org
}

/**
获取数据库中所有的通话记录
*/
func GetTalkDetailsFromDB() DBTalkData {
	var result []TalkDetail
	o := orm.NewOrm()
	var maps []orm.Params
	querySeter := o.QueryTable("talk_detail")
	querySeter.Values(&maps)
	for _, valu := range maps {
		if bytes, err := json.Marshal(valu); err == nil {
			talk := TalkDetail{}
			_ = json.Unmarshal(bytes, &talk)
			result = append(result, talk)
		}
	}
	return result
}

/**
当前涉及的所有项目
*/
func (dbTalkData DBTalkData) getprojectXref() map[string]bool {
	projectXref := make(map[string]bool)
	for _, talk := range dbTalkData {
		if _, ok := projectXref[talk.ProjectId]; !ok {
			projectXref[talk.ProjectId] = true
		}
	}
	return projectXref
}

/**
获取项目 id 和 guid 的对应关系
*/
func (dbTalkData DBTalkData) GetProjectInfoFromDB() map[string]string {
	xref := make(map[string]string)
	for k, _ := range dbTalkData.getprojectXref() {
		o := orm.NewOrm()
		org := BaseOrganization{
			Id: k,
		}
		if err := o.Read(&org, "id"); err != nil {
			fmt.Println(err)
			continue
		}
		xref[org.Id] = org.ProjectGuid
	}
	return xref
}

/**
构建推送内容
*/
func (talkDetail *TalkDetail) BuildMqPushData(projectGuid string) string {

	var knowChannelValues []*TagsValues

	baseData := shimaobean.CreateShiMaoPushData()
	baseData.Data.RingDuration = talkDetail.RingCostTime
	baseData.Data.Duration = talkDetail.TalkCostTime
	baseData.Data.Phone = talkDetail.PhoneNumber
	baseData.Data.CallStatus = talkDetail.TalkType
	baseData.Data.FileName = talkDetail.TalkAudio
	baseData.Data.CallTime = talkDetail.TalkDate.Format("2006-01-02 15:04:05")
	baseData.MerchantId = projectGuid
	customerDetail := talkDetail.GetCustomerDetailByTalkDetail()
	customerTag := customerDetail.GetCustomerTagByCustomerId()
	baseData.Data.Name = customerDetail.UserName
	if constant.Male == customerDetail.Gender {
		baseData.Data.Sex = "MALE"
	}
	if constant.FEMALE == customerDetail.Gender {
		baseData.Data.Sex = "FEMALE"
	}
	baseData.Data.SalerId = customerDetail.AssignSaler
	baseData.Data.SalerName = customerDetail.SalerName
	baseData.Data.Note = customerTag.Note

	root := CreateRoot()
	if customerTag.BuyIntent == "" {
		root.UpdateFieldParameter.Data.Intents.Values = make([]IntentsValues, 0)
	} else {
		intentValues := root.UpdateFieldParameter.Data.Intents.Values
		for _, iv := range intentValues {
			if customerTag.BuyIntent == iv.Name {
				baseData.Data.Intent = shimaobean.IntentValue{
					Level: iv.Level,
					Name:  iv.Name,
					Id:    iv.Id,
				}
				break
			}
		}
	}

	tagsValues := root.UpdateFieldParameter.Data.Tags
	for _, tagValue := range tagsValues {
		switch tagValue.Name {
		case constant.Gender:
			if customerDetail.Gender != "" {
				tagValue.Values = []TagsValuesInner{{Name: customerDetail.Gender}}
			} else {
				tagValue.Values = []TagsValuesInner{}
			}
		case constant.IntentionalFormat:
			if customerTag.Intent != "" {
				tagValue.Values = []TagsValuesInner{{Name: customerTag.Intent}}
			} else {
				tagValue.Values = []TagsValuesInner{}
			}
		case constant.CustomerGroup:
			if customerTag.Consumer != "" {
				tagValue.Values = []TagsValuesInner{{Name: customerTag.Consumer}}
			} else {
				tagValue.Values = []TagsValuesInner{}
			}
		case constant.IntentionType:
			if customerTag.HouseType != "" {
				tagValue.Values = []TagsValuesInner{{Name: customerTag.HouseType}}
			} else {
				tagValue.Values = []TagsValuesInner{}
			}
		case constant.HomePurpose:
			if customerTag.HouseTarget != "" {
				tagValue.Values = []TagsValuesInner{{Name: customerTag.HouseTarget}}
			} else {
				tagValue.Values = []TagsValuesInner{}
			}
		case constant.CustomerType:
			if customerTag.ConsumerType != "" {
				tagValue.Values = []TagsValuesInner{{Name: customerTag.ConsumerType}}
			} else {
				tagValue.Values = []TagsValuesInner{}
			}
		case constant.IntentToList:
			if customerTag.IntentHouse != "" {
				tagValue.Values = []TagsValuesInner{{Name: customerTag.IntentHouse}}
			} else {
				tagValue.Values = []TagsValuesInner{}
			}
		case constant.Resistance:
			if customerTag.Resistance != "" {
				tagValue.Values = []TagsValuesInner{{Name: customerTag.Resistance}}
			} else {
				tagValue.Values = []TagsValuesInner{}
			}
		case constant.FocusPoint:
			attention := customerTag.Attention
			if attention != "" {
				inners := make([]TagsValuesInner, 0)
				for _, str := range strings.Split(attention, ",") {
					inners = append(inners, TagsValuesInner{Name: str})
				}
				tagValue.Values = inners
			} else {
				tagValue.Values = []TagsValuesInner{}
			}
		case constant.CognitiveChannel:
			bytes, err := json.Marshal(tagValue.Values)
			if err != nil {
				fmt.Println(err)
			}
			err = json.Unmarshal(bytes, &knowChannelValues)
			if err != nil {
				fmt.Println(err)
			}
			if customerTag.KnowChannel != "" {
				c := make(map[string][]string)
				for _, s1 := range strings.Split(customerTag.KnowChannel, ",") {
					strs := strings.Split(s1, "-")
					if value, ok := c[strs[0]]; ok {
						c[strs[0]] = append(value, strs[1])
					} else {
						c[strs[0]] = []string{strs[1]}
					}
				}
				for _, kcV := range knowChannelValues {
					subContents := c[kcV.Name]
					var tagsValuesInner []TagsValuesInner
					for _, innerEle := range subContents {
						tagsValuesInner = append(tagsValuesInner, TagsValuesInner{Name: innerEle})
					}
					if len(tagsValuesInner) > 0 {
						kcV.Values = tagsValuesInner
					} else {
						kcV.Values = make([]TagsValuesInner, 0)
					}
				}
				tagValue.Values = knowChannelValues
			} else {
				for _, kcV := range knowChannelValues {
					kcV.Values = make([]TagsValuesInner, 0)
				}
				tagValue.Values = knowChannelValues
			}
		}

	}
	bytes1, _ := json.Marshal(baseData)
	bytes2, _ := json.Marshal(root.UpdateFieldParameter.Data.Tags)
	return strings.ReplaceAll(string(bytes1), "\"dd\"", string(bytes2))
}

// ===== ===== ===== ===== ===== ===== < end > ===== ===== ===== ===== ===== =====

func init() {
	if err := orm.RegisterDataBase(
		"default",
		"mysql",
		config.Config.MySQLDataSource); err != nil {
		fmt.Println(err)
		return
	}
	orm.RegisterModel(new(TalkDetail), new(BaseOrganization), new(CustomerDetail), new(CustomerTag))
}
