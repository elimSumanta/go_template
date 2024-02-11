package model

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type JWTSigningPayload struct {
	Email            string    `json:"email"`
	HasMasterAccess  int       `json:"has_master_access"`
	IsAllowedTrading int       `json:"is_allowed_trading"`
	IsReadOnly       int       `json:"is_read_only"`
	ExpiresIn        int64     `json:"expires_in"`
	ExpiresTime      time.Time `json:"expires_time"`
	Subject          string    `json:"sub"` // Subject is User UUID
	IssuedAt         int64     `json:"iat"`
	jwt.RegisteredClaims
}

type BitwyreJWTClaim struct {
	Email            string    `json:"email"`
	HasMasterAccess  int       `json:"has_master_access"`
	IsAllowedTrading int       `json:"is_allowed_trading"`
	IsReadOnly       int       `json:"is_read_only"`
	ExpiresIn        int64     `json:"expires_in"`
	ExpiresTime      time.Time `json:"expires_time"`
	jwt.RegisteredClaims
}
