package cmd

import (
	mainflow "github.com/bitwyre/bitwyre/gateway/rest/p2p_api/internal/handler/main_flow"
	userHandler "github.com/bitwyre/bitwyre/gateway/rest/p2p_api/internal/handler/user"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"

	pkgJWTMiddlerware "github.com/bitwyre/bitwyre/gateway/rest/p2p_api/internal/pkg/middleware"

	repoAds "github.com/bitwyre/bitwyre/gateway/rest/p2p_api/internal/repo/ads"
	repoJWT "github.com/bitwyre/bitwyre/gateway/rest/p2p_api/internal/repo/jwt-service"
	repoUser "github.com/bitwyre/bitwyre/gateway/rest/p2p_api/internal/repo/user-auth"

	ucAds "github.com/bitwyre/bitwyre/gateway/rest/p2p_api/internal/usecase/ads"
	"github.com/gin-gonic/gin"
)

type Service struct {
	IUserHandler
	IMainFlowHandler
	IJWTMiddlewareP2P
}

func NewService(db *gorm.DB, rds *redis.Client) (*Service, error) {

	// Initialize the Repository
	repoJWT := repoJWT.NewJWTService(db)
	repoUser := repoUser.NewUserAuthenticationRepo(db)
	repoAds := repoAds.NewAdsRepo(db, rds)
	// --------------------------------

	// Initialize the Usecase Service
	sJWTMiddleWare := pkgJWTMiddlerware.NewService(repoUser, repoJWT)
	ucAds := ucAds.InitAdsUsecase(ucAds.RepoInitiated{
		AdsRepo: repoAds,
	})
	// --------------------------------

	// Initialize the Handler
	iUserHandler := userHandler.Init()
	iMainFLowHandler := mainflow.Init(ucAds)
	// --------------------------------

	return &Service{
		IUserHandler:      iUserHandler,
		IMainFlowHandler:  iMainFLowHandler,
		IJWTMiddlewareP2P: sJWTMiddleWare,
	}, nil
}

// Define Interface
type IUserHandler interface {
	GetProfile(ctx *gin.Context)
}

type IMainFlowHandler interface {
	GetCryptoList(ctx *gin.Context)
	GetAdsList(ctx *gin.Context)
}

type IJWTMiddlewareP2P interface {
	JWTMiddlewareP2P() gin.HandlerFunc
	SetUpCors() gin.HandlerFunc
}
