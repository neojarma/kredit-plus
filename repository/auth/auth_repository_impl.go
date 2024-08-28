package auth_repository

import (
	"kredit_plus/models"

	"gorm.io/gorm"
)

type authRepositoryImpl struct {
	DB *gorm.DB
}

func NewAuthRepository(db *gorm.DB) AuthRepository {
	return &authRepositoryImpl{
		DB: db,
	}
}

func (m *authRepositoryImpl) InsertToAuth(payload *models.Auths) error {
	return m.DB.Create(payload).Error
}

func (m *authRepositoryImpl) Login(payload *models.Auths) (*models.Auths, error) {
	model := new(models.Auths)
	return model, m.DB.Where("email = ?", payload.Email).First(model).Error
}
