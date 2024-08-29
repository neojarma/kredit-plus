package assets_service

import "kredit_plus/models"

type AssetsService interface {
	GetAllAsset() ([]*models.Assets, error)
}
