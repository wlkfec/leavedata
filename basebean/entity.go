package basebean

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/orm"
	"io/ioutil"
	"makeuse/shimaobean"
	"time"
)

type Root struct {
	UpdateFieldParameter FieldParameter
}

func BuildRootByFile() *Root {
	var (
		fileBytes []byte
		err       error
		root      Root
	)
	if fileBytes, err = ioutil.ReadFile("content.json"); err != nil {
		fmt.Println("文件读取失败")
		fmt.Println(err)
		return nil
	}
	if err = json.Unmarshal(fileBytes, &root); err != nil {
		fmt.Println(err)
		return nil
	}
	return &root
}

type FieldParameter struct {
	Head       string             `json:"head"`
	MerchantId string             `json:"merchantId"`
	CreateTime string             `json:"createTime"`
	Data       FieldParameterData `json:"data"`
}

type FieldParameterData struct {
	Intents Intents      `json:"intents"`
	Tags    []TagsValues `json:"tags"`
}

// ----- ----- ----- ----- ----- ----- ----- -----

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
	IsRequired bool              `json:"isRequired"`
	Values     []TagsValuesInner `json:"values"`
	Name       string            `json:"name"`
	Index      int               `json:"index"`
	Modifiable bool              `json:"modifiable"`
	IsRadio    bool              `json:"isRadio"`
	Key        string            `json:"key"`
}

type TagsValuesInner struct {
	Name string `json:"name"`
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

func (customerDetail *CustomerDetail) GetCustomerTagByCustomerId() *CustomerTag {

	o := orm.NewOrm()
	org := CustomerTag{
		CustomerId: customerDetail.Id,
	}
	if err := o.Read(&org, "customer_id"); err != nil {
		fmt.Println(err)
	}

	return &org
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

// 定义通话记录的结构体对象
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

func (talkDetail *TalkDetail) GetCustomerDetailByTalkDetail() *CustomerDetail {

	o := orm.NewOrm()
	org := CustomerDetail{
		PhoneNumber: talkDetail.PhoneNumber,
		BuildingId:  talkDetail.BuildingId,
	}
	if err := o.Read(&org, "customer_id"); err != nil {
		fmt.Println(err)
	}

	return &org
}

/**
构建推送内容
*/
func (talkDetail *TalkDetail) BuildMqPushData(projectGuid string) string {
	// 构建实际的推送字符串
	baseData := shimaobean.CreateShiMaoPushData()
	baseData.Data.RingDuration = talkDetail.RingCostTime
	baseData.Data.Duration = talkDetail.TalkCostTime
	baseData.Data.Phone = talkDetail.PhoneNumber
	baseData.Data.CallStatus = talkDetail.TalkType
	baseData.Data.FileName = talkDetail.TalkAudio
	baseData.Data.CallTime = talkDetail.TalkDate.Format("2006-01-02 15:04:05")
	baseData.MerchantId = projectGuid

	// 获取客户和客户标签
	customerDetail := talkDetail.GetCustomerDetailByTalkDetail()
	customerTag := customerDetail.GetCustomerTagByCustomerId()
	baseData.Data.Name = customerDetail.UserName
	if "男" == customerDetail.Gender {
		baseData.Data.Sex = "MALE"
	}
	if "女" == customerDetail.Gender {
		baseData.Data.Sex = "FEMALE"
	}
	baseData.Data.SalerId = customerDetail.AssignSaler
	baseData.Data.SalerName = customerDetail.SalerName
	baseData.Data.Note = customerTag.Note

	// 处理意向类型 intent
	root := BuildRootByFile()
	buyIntent := customerTag.BuyIntent
	if buyIntent == "" {
		root.UpdateFieldParameter.Data.Intents.Values = make([]IntentsValues, 0)
	} else {
		intentValues := root.UpdateFieldParameter.Data.Intents.Values
		for _, iv := range intentValues {
			if buyIntent == iv.Name {
				baseData.Data.Intent = shimaobean.IntentValue{
					Level: iv.Level,
					Name:  iv.Name,
					Id:    iv.Id,
				}
				break
			}
		}
	}

	// 处理各种标签
	tagsValues := root.UpdateFieldParameter.Data.Tags
	for _, tagValue := range tagsValues {

	}

	return ""
}

func init() {
	if err := orm.RegisterDataBase(
		"default",
		"mysql",
		"root:Unisound@123@tcp(192.168.5.184:3306)/test"); err != nil {
		fmt.Println(err)
		return
	}
	orm.RegisterModel(new(TalkDetail), new(BaseOrganization), new(CustomerDetail), new(CustomerTag))
}

func GetTalkDetailFromDB() DBTalkData {
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

// 解析配置信息
type Config struct {
	Mysql struct {
		Dbname   string
		Username string
		Password string
	}
}

type DBTalkData []TalkDetail

func (dbTalkData DBTalkData) getprojectXref() map[string]bool {
	projectXref := make(map[string]bool)
	for _, talk := range dbTalkData {
		if _, ok := projectXref[talk.ProjectId]; !ok {
			projectXref[talk.ProjectId] = true
		}
	}
	return projectXref
}

func (dbTalkData DBTalkData) GetProjectInfoFromDB() map[string]string {
	xref := make(map[string]string)
	for k, _ := range dbTalkData.getprojectXref() {
		// 查询出对应的组织数据
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
