package middleware

import (
	"net/http"

	"github.com/bitwyre/bitwyre/gateway/rest/p2p_api/internal/pkg/helper"
	"github.com/bitwyre/bitwyre/gateway/rest/p2p_api/internal/pkg/library"
	"github.com/gin-gonic/gin"
)

func BasicAPIKeyMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		apiKey := c.GetHeader("X-Api-Key")

		if apiKey == "" || apiKey != library.GetConfig.BasicApiKey {
			helper.HttpErrorResponse(http.StatusUnauthorized, helper.BaseErrorResponseSchema{
				Code:    "UNAUTHORIZED",
				Message: "Invalid API-Key. Please try again or generate new pair.",
			}, nil, c)
			c.Abort()
		}

		c.Writer.Header().Set("Content-Type", "application/json; charset=utf-8")
		c.Next()
	}
}
