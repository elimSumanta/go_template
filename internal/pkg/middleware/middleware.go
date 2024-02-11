package middleware

import (
	"context"

	jwtModel "github.com/bitwyre/bitwyre/gateway/rest/p2p_api/internal/model"
	jwtService "github.com/bitwyre/bitwyre/gateway/rest/p2p_api/internal/repo/jwt-service"
	userAuth "github.com/bitwyre/bitwyre/gateway/rest/p2p_api/internal/repo/user-auth"
	"github.com/bitwyre/bitwyre/shared/go-entity/entity"
)

type IJWTService interface {
	ParseValidate(tokenStr string, pubKey string) (*jwtModel.BitwyreJWTClaim, error)
	ParseNoValidate(tokenStr string) (*jwtModel.BitwyreJWTClaim, error)
	GetUserJWTKey(userUuid string, c context.Context) (*entity.JWTKeys, error)
}

type IUserAuthentication interface {
	FindByUUID(user_uuid string, c context.Context) (entity.Authentication, error)
}

type JwtService struct {
	IUserAuthentication
	IJWTService
}

func NewService(userAuth *userAuth.UserAuthenticationRepo, jwtService *jwtService.JwtService) *JwtService {
	return &JwtService{
		IUserAuthentication: userAuth,
		IJWTService:         jwtService,
	}
}
