package jwtservice

import (
	"context"
	"fmt"
	"time"

	jwtModel "github.com/bitwyre/bitwyre/gateway/rest/p2p_api/internal/model"
	"github.com/bitwyre/bitwyre/shared/go-entity/entity"
	"github.com/golang-jwt/jwt/v4"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type JwtService struct {
	db *gorm.DB
}

func NewJWTService(db *gorm.DB) *JwtService {
	return &JwtService{
		db: db,
	}
}

func (s *JwtService) ParseValidate(tokenStr string, pubKey string) (*jwtModel.BitwyreJWTClaim, error) {
	publicKey, err := jwt.ParseRSAPublicKeyFromPEM([]byte(pubKey))
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	token, err := jwt.ParseWithClaims(tokenStr, &jwtModel.BitwyreJWTClaim{}, func(parsed *jwt.Token) (interface{}, error) {
		if _, ok := parsed.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", parsed.Header["alg"])
		}
		return publicKey, nil
	})

	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	claim, ok := token.Claims.(*jwtModel.BitwyreJWTClaim)
	if !ok || !token.Valid {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	tokenExpired := time.Unix(claim.ExpiresIn, 0).Before(time.Now())

	if tokenExpired {
		return nil, status.Error(codes.PermissionDenied, "token expired")
	}

	return claim, nil
}

func (s *JwtService) ParseNoValidate(tokenStr string) (*jwtModel.BitwyreJWTClaim, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &jwtModel.BitwyreJWTClaim{}, nil, jwt.WithoutClaimsValidation())

	if err != nil {
		if err.Error() != "no Keyfunc was provided." {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
	}

	claim, _ := token.Claims.(*jwtModel.BitwyreJWTClaim)

	return claim, nil
}

func (s *JwtService) GetUserJWTKey(userUuid string, c context.Context) (*entity.JWTKeys, error) {

	var data entity.JWTKeys

	err := s.db.WithContext(c).Where("user_uuid = ?", userUuid).First(&data).Error
	if err != nil {
		return nil, status.Error(codes.PermissionDenied, "user jwt key not found")
	}

	return &data, nil
}
