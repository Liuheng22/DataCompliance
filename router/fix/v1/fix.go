package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//修复电话
func FixPhone(c *gin.Context) {
	//获取电话
	phone := c.Query("phone")
	//检查是否为别的错误
	if len(phone) != 11 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid phone number",
		})
		return
	}
	//修改phone的数据，变成中间四位为星号
	phone = phone[:3] + "****" + phone[7:]
	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
		"phone":   phone,
	})
}

//修复地址
func FixAddress(c *gin.Context) {
	//获取地址
	address := c.Query("address")
	//地址无效则返回错误提示
	if len(address) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "No address!",
		})
		return
	}
	//如果地址有效
	address = address[:len(address)-6] + "******"
	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
		"address": address,
	})
}

//修复名字
func FixName(c *gin.Context) {
	//获取名字
	name := c.Query("name")
	//名字无效则错误提示
	if len(name) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "No name!",
		})
		return
	}
	//如果名字有效
	name = name[:1] + "**"
	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
		"name":    name,
	})
}

//添加电话
func AddPhone(c *gin.Context) {

}

//添加地址
func AddAddress(c *gin.Context) {

}

//添加名字
func AddName(c *gin.Context) {

}

//检查电话
func CheckPhone(c *gin.Context) {

}

//检查地址
func CheckAddress(c *gin.Context) {

}

//检查名字
func CheckName(c *gin.Context) {

}
