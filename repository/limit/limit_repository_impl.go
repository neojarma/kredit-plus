package limit_repository

import (
	"kredit_plus/models"

	"gorm.io/gorm"
)

type tenorLimitRepositoryImpl struct {
	DB *gorm.DB
}

func NewTenorLimitRepository(db *gorm.DB) TenorLimitRepository {
	return &tenorLimitRepositoryImpl{
		DB: db,
	}
}

func (m *tenorLimitRepositoryImpl) GetTenorLimit(salary float64) (*models.LimitTenor, error) {
	model := new(models.LimitTenor)
	err := m.DB.Model(model).
		Where("? >= salary_min AND ? <= salary_max", salary, salary).
		Order("salary_max DESC").Find(model).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}

	return model, nil
}
