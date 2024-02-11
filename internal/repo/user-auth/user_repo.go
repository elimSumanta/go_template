package userauth

import (
	"context"

	"github.com/bitwyre/bitwyre/shared/go-entity/entity"
	"gorm.io/gorm"
)

type UserAuthenticationRepo struct {
	db *gorm.DB
}

func NewUserAuthenticationRepo(db *gorm.DB) *UserAuthenticationRepo {
	return &UserAuthenticationRepo{db}
}

func (r *UserAuthenticationRepo) FindByUUID(user_uuid string, c context.Context) (entity.Authentication, error) {
	var data entity.Authentication

	err := r.db.WithContext(c).Where("user_uuid = ?", user_uuid).First(&data).Error
	if err != nil {
		return data, err
	}

	return data, nil
}
