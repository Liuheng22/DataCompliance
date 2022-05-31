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
	resdata := HandleText(testdata)
	return
}

func HandleText(data *text.Test) *text.Testres {
	res := &text.Testres{}
}
