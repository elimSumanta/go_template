package cmd

import (
	handler "github.com/bitwyre/bitwyre/gateway/rest/p2p_api/internal/handler"
	"github.com/gin-gonic/gin"
	rkgin "github.com/rookie-ninja/rk-gin/v2/boot"
)

type ClientRoutes struct {
	Private *gin.RouterGroup
	Public  *gin.RouterGroup
	Service *Service
}

func InitRouter(service *Service) *ClientRoutes {
	gitRoute := rkgin.GetGinEntry("p2p-service-go")
	gitRoute.Router.NoRoute(noRouteError)
	gitRoute.Router.GET("/", handler.Health)

	// Register handler
	publicRouter := rkgin.GetGinEntry("p2p-service-go").Router.Group("/public")
	privateRouter := rkgin.GetGinEntry("p2p-service-go").Router.Group("/private")
	privateRouter.Use(service.IJWTMiddlewareP2P.JWTMiddlewareP2P())
	privateRouter.Use(service.IJWTMiddlewareP2P.SetUpCors())

	return &ClientRoutes{
		Private: privateRouter,
		Public:  publicRouter,
		Service: service,
	}
}

func noRouteError(c *gin.Context) {
	c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "Path": c.Request.URL.Path})
}

func (cr *ClientRoutes) InitPrivate() {
	cr.Private.GET("user/profile", cr.Service.GetProfile)
	cr.Private.GET("main/ads/list", cr.Service.GetAdsList)
	cr.Private.GET("main/crypto/list", cr.Service.GetCryptoList)
}

func (cr *ClientRoutes) InitPublic() {
	cr.Public.GET("/healthcheck", handler.Health)
}
