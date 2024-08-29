package assets_service

import (
	"kredit_plus/models"
	assets_repository "kredit_plus/repository/assets"
)

type AssetsServiceImpl struct {
	AssetRepo assets_repository.AssetsRepository
}

func NewAssetsService(repo assets_repository.AssetsRepository) AssetsService {
	return &AssetsServiceImpl{
		AssetRepo: repo,
	}
}

func (s *AssetsServiceImpl) GetAllAsset() ([]*models.Assets, error) {
	return s.AssetRepo.GetAllAsset()
}
