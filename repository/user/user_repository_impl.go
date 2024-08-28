package user_repository

import (
	"kredit_plus/models"

	"gorm.io/gorm"
)

type userRepositoryImpl struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepositoryImpl{
		DB: db,
	}
}

func (m *userRepositoryImpl) InsertUserProfile(payload *models.User) error {
	return m.DB.Create(payload).Error
}
