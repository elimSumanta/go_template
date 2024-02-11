package middleware

import (
	"github.com/bitwyre/bitwyre/gateway/rest/p2p_api/internal/pkg/library"
	"github.com/gin-gonic/gin"
)

func (js *JwtService) SetUpCors() gin.HandlerFunc {
	return func(c *gin.Context) {
		if library.GetConfig.EnableCors {
			c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
			c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
			c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
			c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")
		}

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
