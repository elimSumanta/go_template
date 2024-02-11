package user

import (
	"github.com/bitwyre/bitwyre/gateway/rest/p2p_api/internal/model"
	"github.com/bitwyre/bitwyre/gateway/rest/p2p_api/internal/pkg/helper"
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

// User handler
// @Summary Get User Profile
// @Id user
// @Tags User
// @version 1.0
// @Param Authorization header string true "With the bearer started"
// @produce application/json
// @Success 200 {object} helper.BaseResponseSchema
// @Router /private/user/profile [get]
func (h *UserHandler) GetProfile(ctx *gin.Context) {
	userId := ctx.Request.Header.Get("x-user-id")
	span := trace.SpanFromContext(ctx.Request.Context())
	span.SetAttributes(attribute.String("userId", userId))

	helper.HttpSuccess(model.UserProfile{
		UserID: userId,
	}, nil, ctx)

}
