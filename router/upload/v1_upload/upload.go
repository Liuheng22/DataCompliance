package v1_upload

import (
	"DataCompliance/data/text"
	"DataCompliance/pkg/e"

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
	// 名字的问题,
}

// 处理phone问题
func HandlePhone(phone string, problems []text.Problem) {

}

// 处理address问题
func HandleAddress(address string, problems []text.Problem) {

}

// 处理id问题
func HandleId(id string, problems []text.Problem) {

}
