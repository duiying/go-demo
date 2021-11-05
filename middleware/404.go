package middleware

import "github.com/gin-gonic/gin"

func NotFound() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(404, gin.H{
			"code":    404,
			"message": "api not found",
			"data":    "",
		})
		return
	}
}
