package router

import (
	"DataCompliance/pkg/cors"
	v1 "DataCompliance/router/fix/v1"
	"DataCompliance/router/upload/v1_upload"
	"net/http"

	"github.com/gin-gonic/gin"
)

func StartRouter() *gin.Engine {
	router := gin.Default()
	router.Use(cors.Cors())
	api := router.Group("/api")
	{
		api.GET("/me", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "hello world!"})
		})
	}

	upload := router.Group("/upload")
	{
		upload.POST("/text", v1_upload.UploadText)
		upload.Use(cors.Cors())
	}

	// 修复功能
	fix := router.Group("/fix")
	{
		//修复电话
		fix.GET("/phone", v1.FixPhone)
		//修复地址
		fix.GET("/address", v1.FixAddress)
		//修复名字
		fix.GET("/name", v1.FixName)
		//添加电话
		fix.POST("/phone", v1.AddPhone)
		//添加名字
		fix.POST("/name", v1.AddName)
		//添加地址
		fix.POST("/address", v1.AddAddress)

	}
	return router
}
