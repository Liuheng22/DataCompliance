package cors

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Cors() gin.HandlerFunc {
	c := cors.Config{
		AllowOrigins: []string{"http://localhost:3000"},
		//AllowAllOrigins: true,
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "PATCH"},
		AllowHeaders: []string{"Content-Type", "Access-Token", "Authorization"},
		MaxAge:       6 * time.Hour,
	}

	return cors.New(c)
}
