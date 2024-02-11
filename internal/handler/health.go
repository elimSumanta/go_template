package handler

import (
	"github.com/bitwyre/bitwyre/gateway/rest/p2p_api/internal/pkg/helper"
	"github.com/gin-gonic/gin"
)

// HealthCheck handler
// @Summary HealthCheck
// @Tags HealthCheck
// @version 1.0
// @produce application/json
// @Success 200
// @Router /public/healthcheck [get]
func Health(ctx *gin.Context) {
	helper.HttpSuccess("Health", nil, ctx)
}
