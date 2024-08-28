package assets_repository

import (
	"kredit_plus/models"

	"gorm.io/gorm"
)

type assetsRepositoryImpl struct {
	DB *gorm.DB
}

func NewAssetRepository(db *gorm.DB) AssetsRepository {
	return &assetsRepositoryImpl{
		DB: db,
	}
}

func (r *assetsRepositoryImpl) GetAssetRepository(payload *models.Assets) (*models.Assets, error) {
	model := new(models.Assets)
	return model, r.DB.Where("id_assets = ?", payload.ID).First(model).Error
}
