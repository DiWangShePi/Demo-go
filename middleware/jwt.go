package middleware

import (
	"Demo/utils"
	"errors"
	"github.com/gin-gonic/gin"
)

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.PostForm("token")
		if token == "" {
			c.JSON(200, gin.H{
				"status":  "Error",
				"message": "Can not find token",
			})
			c.Abort()
			return
		}

		claims, err := utils.ParseToken(token)
		if err != nil {
			if errors.Is(err, utils.TokenExpired) {
				c.JSON(200, gin.H{
					"status":  "Error",
					"message": "Token has expired",
				})
				c.Abort()
				return
			}
			c.JSON(200, gin.H{
				"status":  "Error",
				"message": err,
			})
			c.Abort()
			return
		}

		c.Set("claims", claims)
		c.Next()
	}
}
