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
	index := 0
	rowres := &text.Rowres{}
	for _, row := range data.Data {
		rowres, index = HandleRowdata(&row, index)
		res.Data = append(res.Data, *rowres)
	}
	return res
}

//处理Rowdata
func HandleRowdata(row *text.Rowdata, index int) (*text.Rowres, int) {
	resrow := &text.Rowres{Key: row.Key, Name: row.Name, Age: row.Age, Phone: row.Phone, Address: row.Address, Id: row.Id, Problems: make([]text.Problem, 0)}
	// 发现并且添加问题
	// name问题
	problems := make([]text.Problem, 0)
	problems = append(problems, HandleName(row.Name, index+len(problems))...)

	// phone问题
	problems = append(problems, HandlePhone(row.Phone, index+len(problems))...)

	// address问题
	problems = append(problems, HandleAddress(row.Address, index+len(problems))...)

	// id问题
	problems = append(problems, HandleId(row.Id, index+len(problems))...)

	// 将问题汇总
	resrow.Problems = append(resrow.Problems, problems...)
	return resrow, index + len(resrow.Problems)
}

// 处理name问题
func HandleName(name string, index int) []text.Problem {
	// 名字的问题
	problems := make([]text.Problem, 0)
	namerune := []rune(name)
	// 首先名字少于等于一个字，认为名字不完整
	if len(namerune) <= 1 {
		problems = append(problems, text.Problem{
			ID:          index + len(problems),
			Col:         "Name",
			Seriousness: "normal",
			Type:        "完整性",
			Description: "缺少姓名信息",
		})
		return problems
	}

	// 名字的隐私性
	count := strings.Count(name, "*")
	// 有隐蔽，没问题
	if count != 0 {
		return problems
	}
	// 有错构造新的name
	newname := make([]rune, 2)
	newname[0] = namerune[0]
	newname[1] = rune('*')
	if len(namerune) > 2 {
		newname = append(newname, namerune[len(namerune)-1])
	}
	problems = append(problems, text.Problem{
		ID:          index + len(problems),
		Col:         "Name",
		Seriousness: "risky",
		Type:        "隐私性",
		Description: "具体姓名需要隐去",
		Fix:         string(newname),
	})
	return problems
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
func HandlePhone(phone string, index int) []text.Problem {
	// phone问题处理
	problems := make([]text.Problem, 0)
	// phone缺失
	if len(phone) == 0 {
		problems = append(problems, text.Problem{
			ID:          index + len(problems),
			Col:         "Phone",
			Seriousness: "caution",
			Type:        "完整性",
			Description: "缺少电话信息",
		})
		//	电话位数不对，那么不需要考虑后续的东西了
		return problems
	}
	// phone的位数
	if len(phone) != 11 {
		problems = append(problems, text.Problem{
			ID:          index + len(problems),
			Col:         "Phone",
			Seriousness: "critical",
			Type:        "规范性",
			Description: "phone格式与常规不符,且具体电话号码需隐去",
		})
		//	电话位数不对，那么不需要考虑后续的东西了
		return problems
	}

	// 电话隐私性问题
	// 找****的位置
	// 找是否有****
	cnt := strings.Count(phone, "****")
	pos := strings.Index(phone, "****")
	if cnt > 0 && pos == 3 {
		// 隐私上没有错误
		return problems
	}
	// 隐私上有问题
	// 生成正确的phone
	newphone := string(phone[0:3] + "****" + phone[7:])
	problems = append(problems, text.Problem{
		ID:          index + len(problems),
		Col:         "Phone",
		Seriousness: "critical",
		Type:        "隐私性",
		Description: "根据国家法律法规,具体电话号码需隐去",
		Fix:         newphone,
	})
	//	电话位数不对，那么不需要考虑后续的东西了
	return problems
}

// 处理address问题
func HandleAddress(address string, index int) []text.Problem {
	// address问题处理
	problems := make([]text.Problem, 0)
	addressrune := []rune(address)
	if len(addressrune) == 0 {
		problems = append(problems, text.Problem{
			ID:          index + len(problems),
			Col:         "Address",
			Seriousness: "caution",
			Type:        "完整性",
			Description: "缺少地址信息",
		})
		//	没有地址，就不需要后面的了
		return problems
	}
	if len(addressrune) <= 6 {
		problems = append(problems, text.Problem{
			ID:          index + len(problems),
			Col:         "Address",
			Seriousness: "caution",
			Type:        "完整性",
			Description: "缺少地址信息",
		})
		//	没有地址，就不需要后面的了
		return problems
	}

	//地址隐私问题
	cnt := strings.Count(address, "****")
	pos := strings.Index(address, "****")

	if cnt > 0 && pos == len(addressrune)-4 {
		return problems
	}
	// 隐私上有问题
	// 生成正确的phone
	newaddress := string(addressrune[:len(addressrune)-4]) + "****"
	problems = append(problems, text.Problem{
		ID:          index + len(problems),
		Col:         "Address",
		Seriousness: "critical",
		Type:        "隐私性",
		Description: "具体门牌号可能带来风险，建议隐去",
		Fix:         newaddress,
	})
	// 返回
	return problems
}

// 处理id问题
func HandleId(id string, index int) []text.Problem {
	// 问题
	problems := make([]text.Problem, 0)
	// 缺少身份证信息
	if len(id) == 0 {
		problems = append(problems, text.Problem{
			ID:          index + len(problems),
			Col:         "ID",
			Seriousness: "risky",
			Type:        "完整性",
			Description: "缺少身份证信息",
		})
	}
	// 身份证id位数不对
	if len(id) != 18 {
		problems = append(problems, text.Problem{
			ID:          index + len(problems),
			Col:         "ID",
			Seriousness: "caution",
			Type:        "规范性",
			Description: "身份证格式与常规格式不符",
		})
	}

	// 身份证隐私问题
	count := strings.Count(id, "********")
	pos := strings.Index(id, "********")
	if count > 0 && pos == 6 {
		return problems
	}
	newid := id[0:6] + "********" + id[14:]
	problems = append(problems, text.Problem{
		ID:          index + len(problems),
		Col:         "ID",
		Seriousness: "critical",
		Type:        "隐私性",
		Description: "根据国家法律法规，身份证信息需要隐去",
		Fix:         newid,
	})
	return problems
}
