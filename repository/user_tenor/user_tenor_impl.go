package user_tenor_repository

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

func (m *userTenorImpl) GetUserTenorProfile(payload *models.UserTenor) (*models.UserTenor, error) {
	model := new(models.UserTenor)
	err := m.DB.Where("id_user = ?", payload.ID_User).First(model).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, gorm.ErrRecordNotFound
		}
		return nil, err
	}

	return model, nil

}

func (m *userTenorImpl) UpdateUserTenorProfile(payload *models.UserTenor) error {
	model := new(models.UserTenor)
	err := m.DB.Model(model).Where("id_user = ?", payload.ID_User).Updates(payload).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return gorm.ErrRecordNotFound
		}
		return err
	}

	return nil
}
