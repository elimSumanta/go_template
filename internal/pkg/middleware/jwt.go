package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/bitwyre/bitwyre/gateway/rest/p2p_api/internal/pkg/helper"
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel"
)

type CheckUserAuthResp struct {
	Token            string `json:"token,omitempty"`
	UserId           string `json:"userId,omitempty"`
	Email            string `json:"email,omitempty"`
	HasMasterAccess  bool   `json:"has_master_access,omitempty"`
	IsAllowedTrading bool   `json:"is_allowed_trading,omitempty"`
	IsReadOnly       bool   `json:"is_read_only,omitempty"`
}

func (js *JwtService) JWTMiddlewareP2P() gin.HandlerFunc {
	return func(c *gin.Context) {
		helper.Block{
			Try: func() {
				bearerToken := c.GetHeader("authorization")
				token := parseToken(bearerToken)

				// Create a tracer span
				tr := otel.Tracer("http")
				ctx, span := tr.Start(c.Request.Context(), fmt.Sprintf("%s %s", c.Request.Method, c.Request.RequestURI))
				defer span.End()

				parsedToken, err := js.IJWTService.ParseNoValidate(token)
				if err != nil {
					helper.Throw(err)
				}

				userPubKey, err := js.IJWTService.GetUserJWTKey(parsedToken.Subject, c)
				if err != nil {
					helper.Throw(err)
				}

				validToken, err := js.IJWTService.ParseValidate(token, userPubKey.RSAPublicKey)
				if err != nil {
					helper.Throw(err)
				}

				_, err = js.IUserAuthentication.FindByUUID(validToken.Subject, ctx)
				if err != nil {
					helper.Throw(err)
				}

				c.Writer.Header().Set("Content-Type", "application/json; charset=utf-8")
				c.Request.Header.Add("x-user-id", validToken.Subject)
				c.Next()
			},
			Catch: func(e helper.Exception) {
				// logrus.Errorf("Caught %v", e)
				helper.HttpErrorResponse(http.StatusUnauthorized, helper.BaseErrorResponseSchema{
					Code:    "UNAUTHORIZED",
					Message: "Invalid authentication header",
				}, nil, c)
				c.Abort()
			},
			Finally: nil,
		}.Do()
	}
}

// Parse JWT Token from Request Header
// Return string to be used by validateToken() later
func parseToken(bearerToken string) string {
	splitToken := strings.Split(bearerToken, " ")
	if bearerToken == "" || len(splitToken) < 2 {
		helper.Throw("Invalid JWT Format")
	}

	return splitToken[1]
}
