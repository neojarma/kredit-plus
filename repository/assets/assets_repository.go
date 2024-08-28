package assets_repository

import "kredit_plus/models"

type AssetsRepository interface {
	GetAssetRepository(payload *models.Assets) (*models.Assets, error)
}
