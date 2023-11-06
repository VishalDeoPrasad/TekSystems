package middleware

import "github.com/gin-gonic/gin"

func Auth() gin.HandlerFunc {
	return gin.BasicAuth(gin.Accounts{
		"Vishal": "12345",
	})
}
