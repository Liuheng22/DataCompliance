package v1_upload

import (
	"DataCompliance/data/text"
	"DataCompliance/pkg/e"
	"strings"

	"github.com/gin-gonic/gin"
)

func UploadText(c *gin.Context) {
	testdata := &text.Test{}
	err := c.BindJSON(testdata)
	//错误就返回错误
	if err != nil {
		c.JSON(e.ERROR, gin.H{
			"message": e.GetMsg(e.ERROR),
		})
		return
	}
	//正确处理
	resdata := *HandleText(testdata)
	c.JSON(e.SUCCESS, gin.H{
		"message": e.GetMsg(e.SUCCESS),
		"data":    resdata,
	})
}

//处理Test
func HandleText(data *text.Test) *text.Testres {
	//初始化结构
	res := &text.Testres{Data: make([]text.Rowres, 0)}

	//处理每一行然后添加问题
	for _, row := range data.Data {
		res.Data = append(res.Data, *HandleRowdata(&row))
	}
	return res
}

//处理Rowdata
func HandleRowdata(row *text.Rowdata) *text.Rowres {
	resrow := &text.Rowres{Key: row.Key, Name: row.Name, Age: row.Age, Phone: row.Phone, Address: row.Address, Id: row.Id, Problems: make([]text.Problem, 0)}
	// 发现并且添加问题
	// name问题
	problems := make([]text.Problem, 0)
	HandleName(row.Name, problems)

	// phone问题
	HandlePhone(row.Phone, problems)

	// address问题
	HandleAddress(row.Address, problems)

	// id问题
	HandleId(row.Id, problems)

	// 将问题汇总
	resrow.Problems = append(resrow.Problems, problems...)
	return resrow
}

// 处理name问题
func HandleName(name string, problems []text.Problem) {
	// 名字的问题
	// 首先名字少于等于一个字，认为名字不完整
	if len(name) <= 1 {
		problems = append(problems, text.Problem{
			ID:          len(problems),
			Col:         "Name",
			Seriousness: "caution",
			Type:        "完整性",
			Description: "缺少姓名信息",
		})
		return
	}

	// 名字的隐私性
	count := strings.Count(name, "*")
	// 有隐蔽，没问题
	if count != 0 {
		return
	}
	namerune := []rune(name)
	// 有错构造新的name
	newname := make([]rune, 2)
	newname[0] = namerune[0]
	newname[1] = rune('*')
	if len(namerune) > 2 {
		newname = append(newname, namerune[len(newname)-1])
	}
	problems = append(problems, text.Problem{
		ID:          len(problems),
		Col:         "Name",
		Seriousness: "risky",
		Type:        "隐私性",
		Description: "具体姓名需要隐去",
		Fix:         string(newname),
	})

	// // 判断是否都是***
	// flag := true
	//
	// for i := range check {
	// 	if check[i] != rune('*') {
	// 		check[i] = rune('*')
	// 		flag = false
	// 	}
	// }
	// if flag {
	// 	return
	// }
	// // 新的名字
	// var newname string
	// if len(namrune) > 2 {
	// 	strings.Split(newname, "")
	// }
}

// 处理phone问题
func HandlePhone(phone string, problems []text.Problem) {
	// phone问题处理
	// phone缺失
	if len(phone) == 0 {
		problems = append(problems, text.Problem{
			ID:          len(problems),
			Col:         "Phone",
			Seriousness: "caution",
			Type:        "完整性",
			Description: "缺少电话信息",
		})
		//	电话位数不对，那么不需要考虑后续的东西了
		return
	}
	// phone的位数
	if len(phone) != 11 {
		problems = append(problems, text.Problem{
			ID:          len(problems),
			Col:         "Phone",
			Seriousness: "critical",
			Type:        "规范性",
			Description: "phone格式与常规不符,且具体电话号码需隐去",
		})
		//	电话位数不对，那么不需要考虑后续的东西了
		return
	}

	// 电话隐私性问题
	// 找****的位置
	// 找是否有****
	cnt := strings.Count(phone, "****")
	pos := strings.Index(phone, "****")
	if cnt > 0 && pos == 3 {
		// 隐私上没有错误
		return
	}
	// 隐私上有问题
	// 生成正确的phone
	newphone := string(phone[0:3] + "****" + phone[7:])
	problems = append(problems, text.Problem{
		ID:          len(problems),
		Col:         "Phone",
		Seriousness: "critical",
		Type:        "隐私性",
		Description: "根据国家法律法规,具体电话号码需隐去",
		Fix:         newphone,
	})
	//	电话位数不对，那么不需要考虑后续的东西了
	return
}

// 处理address问题
func HandleAddress(address string, problems []text.Problem) {

}

// 处理id问题
func HandleId(id string, problems []text.Problem) {

}
