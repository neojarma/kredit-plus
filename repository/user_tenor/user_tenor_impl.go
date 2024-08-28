package user_tenor

import (
	"kredit_plus/models"

	"gorm.io/gorm"
)

type userTenorImpl struct {
	DB *gorm.DB
}

func NewUserTenorRepository(db *gorm.DB) UserTenor {
	return &userTenorImpl{
		DB: db,
	}
}

func (m *userTenorImpl) InsertUserTenorProfile(payload *models.UserTenor) error {
	return m.DB.Create(payload).Error
}
