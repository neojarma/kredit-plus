package assets_repository

import "kredit_plus/models"

type AssetsRepository interface {
	GetAsset(payload *models.Assets) (*models.Assets, error)
	GetAllAsset() ([]*models.Assets, error)
}
