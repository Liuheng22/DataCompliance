package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func StartRouter() *gin.Engine {
	router := gin.Default()
	api := router.Group("/api")
	{
		api.GET("/me", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "hello world!"})
		})
	}
	return router
}
