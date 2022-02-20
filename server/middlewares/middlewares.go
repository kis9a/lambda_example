package middlewares

import (
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// config := config.GetConfig()
		// reqKey := c.Request.Header.Get("X-Auth-Key")
		// reqSecret := c.Request.Header.Get("X-Auth-Secret")

		// zap.L().Info(reqKey)
		// zap.L().Info(reqSecret)

		// var key string
		// var secret string
		// if key = config.HTTP_AUTH_KEY; len(strings.TrimSpace(key)) == 0 {
		// 	c.AbortWithStatus(500)
		// }
		// if secret = config.HTTP_AUTH_SECRET; len(strings.TrimSpace(secret)) == 0 {
		// 	c.AbortWithStatus(401)
		// }
		// if key != reqKey || secret != reqSecret {
		// 	c.AbortWithStatus(401)
		// 	return
		// }
		c.Next()
	}
}
